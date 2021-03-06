package bb

// to set up gorilla/websocket type command as;
// > go get github.com/gorilla/websocket

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"io"
	"log"
	"net/url"
	"os"
	"os/signal"
	"time"
)

const CHANNEL_ORDER_BOOK_200 = "orderBook_200.100ms.BTCUSD"
const CHANNEL_TRADE = "trade.BTCUSD"
const CHANNEL_INFO = "instrument_info.100ms.BTCUSD"
const CHANNEL_LIQUIDATION = "liquidation.BTCUSD"

const writeWait = 30 * time.Second
const pongWait = 60 * time.Second
const pingPeriod = 60 * time.Second

func Connect(flagFileName string, w io.WriteCloser, closeWaitMin int) {

	var flagFile FlagFile
	flagFile.Init(flagFileName)
	flagFile.Create()
	peerReset := make(chan struct{})
	record := make(chan string)

	// wait 300 sec to terminate
	go flagFile.Check_other_process_loop(300, peerReset)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	// wss://stream.bybit.com/realtime
	// u := url.URL{Scheme: "ws", Host: "stream.bybit.com", Path: "/realtime"}
	u := url.URL{Scheme: "wss", Host: "stream.bybit.com", Path: "/realtime"}
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}

	// c.SetReadDeadline(time.Now().Set(pongWait))
	c.SetPongHandler(func(string) error { c.SetReadDeadline(time.Now().Add(pongWait)); return nil })

	//websocket.PingMessage

	defer func(c *websocket.Conn) {
		err := c.Close()
		if err != nil {

		}
	}(c)

	write := func(s string) {
		record <- s
	}

	// message receive loop (go routine)
	go func() {
		var messageCount int
		var boardUpdateCount int
		var tradeCount int
		var infoCount int
		var liquidCount int

		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("[ERROR] ReadMessage:", err)
				close(interrupt)
				return
			}
			decoded := ParseMessage(message)

			if messageCount%1000 == 0 {
				log.Printf("%d total / %d board update/ %d execute/ %d info / %d liquid", messageCount, boardUpdateCount, tradeCount, infoCount, liquidCount)
			}
			messageCount += 1

			switch decoded.Topic {
			case CHANNEL_ORDER_BOOK_200:
				boardUpdateCount += 1
				s := ParseOrderBookMessage(decoded)
				write(s)
			case CHANNEL_TRADE:
				tradeCount += 1
				s := ParseTradeMessage(decoded)
				write(s)
			case CHANNEL_INFO:
				infoCount += 1
				s := ""

				if decoded.Type == "snapshot" {
					s = ParseInstrumentSnapshot(decoded.Data, decoded.TimeStampE6)
				} else if decoded.Type == "delta" {
					s = ParseInstrumentDelta(decoded.Data, decoded.TimeStampE6)
				} else {
					log.Println("unknown instrument info type", string(message))
				}
				write(s)

			case CHANNEL_LIQUIDATION:
				liquidCount += 1
				s := ParseLiquidationMessage(decoded.Data)
				write(s)

			default:
				log.Println("[OTHER CHANNEL]", string(message))
			}
		}
	}()

	subscribe := func(ch string) {
		param := make(map[string]interface{})
		param["op"] = "subscribe"
		args := []string{ch}
		param["args"] = args
		req, _ := json.Marshal(param)
		c.WriteMessage(websocket.TextMessage, []byte(req))
	}

	subscribe(CHANNEL_ORDER_BOOK_200)
	subscribe(CHANNEL_TRADE)
	subscribe(CHANNEL_INFO)
	subscribe(CHANNEL_LIQUIDATION)

	// main message loop
	for {
		select {
		case <-peerReset:
			log.Println("Peer reset")
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
			}
			w.Close()
			goto closeWait

		case <-interrupt:
			log.Println("interrupt")

			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
			if err != nil {
				log.Println("write close:", err)
			}
			w.Close()
			goto exit

		case s := <-record:
			if s != "" {
				w.Write([]byte(s))
			}
		}
	}

closeWait:
	{
		log.Println("Peer reset close")

		s := 0
		for s < closeWaitMin {
			s += 1
			time.Sleep(time.Minute) // sleep min
			log.Printf("[wait min] %4d/%d", s, closeWaitMin)
		}
	}
exit:
	log.Println("Logger End")
}
