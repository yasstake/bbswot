package bb

import (
	"bbswot/common"
	"encoding/json"
	"log"
)


type Trade struct {
	Topic string          `json:"topic"`   // "topic":"ParseTradeMessage.BTCUSD"
	// Type  string          `json:"type"`
	Data  json.RawMessage `json:"data"`    //
}

type TradeRec struct {
	TimeStampMs      int64       `json:"trade_time_ms"`  // "trade_time_ms":1619398389868
	Timestamp string      `json:"timestamp"`             // "timestamp":"2021-04-26T00:53:09.000Z"
	Symbol    string      `json:"symbol"`                // "symbol":"BTCUSD"
	Side      string      `json:"side"`                  // "side":"Sell"
	Size      json.Number `json:"size"`                  // "size":2000
	Price     json.Number `json:"price"`                 // "price":50703.5
	TickDirection string  `json:"tick_direction""`       // "tick_direction":"ZeroMinusTick"
	TradeId   string      `json:"trade_id"`              // "trade_id":"8241a632-9f07-5fa0-a63d-06cefd570d75"
}


func (c *TradeRec) ToLog() (result string) {
	var action int
	if c.Side == "Buy" {
		action = common.TRADE_BUY
	} else if c.Side == "Sell" {
		action = common.TRADE_SELL
	} else {
		log.Fatalln("Unknown side", c.Side)
	}

	price, err := c.Price.Float64()
	if err != nil {
		log.Println(err)
	}
	volume, err := c.Size.Float64()
	if err != nil {
		log.Println(err)
	}

	return MakeLogRec(action, c.TimeStampMs, price, volume, c.TradeId)
}




type TradeRecs []TradeRec

func ParseTradeMessage(message Message) (result string) {
	var data TradeRecs

	result = ""

	err := json.Unmarshal(message.Data, &data)
	if err != nil {
		log.Fatalln("Fail to pase message", err, message)
	}

	l := len(data)

	for i := 0; i < l; i++ {
		result += data[i].ToLog()
	}

	return result
}

