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
	"sync"
	"time"
)

const CHANNEL_ORDER_BOOK_200 = "orderBook_200.100ms.BTCUSD"
const CHANNEL_TRADE = "trade.BTCUSD"
const CHANNEL_INFO = "instrument_info.100ms.BTCUSD"

const writeWait = 30 * time.Second
const pongWait = 60 * time.Second
const pingPeriod = 60 * time.Second

func Connect(flagFileName string, w io.WriteCloser, closeWaitMin int) {

	var flagFile FlagFile
	flagFile.Init(flagFileName)
	flagFile.Create()
	peerReset := make(chan struct{})

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

	// c.SetReadDeadline(time.Now().Add(pongWait))
	c.SetPongHandler(func(string) error { c.SetReadDeadline(time.Now().Add(pongWait)); return nil })

	//websocket.PingMessage

	defer func(c *websocket.Conn) {
		err := c.Close()
		if err != nil {

		}
	}(c)

	var mutex sync.Mutex
	inLoop := true

	write := func(s string) {
		mutex.Lock()
		defer mutex.Unlock()

		w.Write([]byte(s))
	}

	go func() {
		var lastLiquidTime int64
		var sleepTime int

		for {
			liqs, _, err := LiquidRequest(&lastLiquidTime)
			if err != nil {
				log.Println(err)
			}

			if len(liqs) != 0 {
				write(liqs.ToLog())
				log.Println("liquid ", len(liqs), " records")
				sleepTime = 1
			} else {
				log.Println("liquid sleep", sleepTime)
				sleepTime = sleepTime + 5
				if 30 <= sleepTime {
					sleepTime = 30
				}
			}

			if !inLoop {
				log.Println("exit liquid listen loop")
				break
			}
			time.Sleep(time.Duration(sleepTime * int(time.Second)))
		}
	}()

	go func() {
		var messageCount int
		var boardUpdateCount int
		var tradeCount int
		var infoCount int

		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				log.Println("[ERROR] ReadMessage:", err)
				close(interrupt)
				return
			}
			decoded := ParseMessage(message)

			if messageCount%1000 == 0 {
				log.Printf("%d total / %d board update/ %d execute/ %d info",
					messageCount, boardUpdateCount, tradeCount, infoCount)
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
					s = ParseInstrumentSnapshot(decoded.Data, decoded.TimeStampMs)
				} else if decoded.Type == "delta" {
					s = ParseInstrumentDelta(decoded.Data, decoded.TimeStampMs)
				} else {
					log.Println("unknown instrument info type", string(message))
				}
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

	// ticker := time.NewTicker(pingPeriod)

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
		/*
		case <-ticker.C:
			log.Println("ping")
			c.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		 */
		}
	}

closeWait:
	{
		inLoop = false
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

