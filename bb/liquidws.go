package bb

import (
	"bbswot/common"
	"encoding/json"
	"log"
)

type LiquidWs struct {
	Symbol      string      `json:"symbol"` // "symbol":"BTCUSD",
	Side        string      `json:"side"`   // "side":"Sell",
	Price       json.Number `json:"price"`  // "price":49490.5}
	Volume      json.Number `json:"qty"`    // "qty":1600,
	TimeStampMs int64       `json:"time"`   // "time":1630110808068,
}

func (c *LiquidWs) ToLog() string {
	var action int
	if c.Side == "Buy" {
		action = common.TRADE_BUY_LIQUID
	} else if c.Side == "Sell" {
		action = common.TRADE_SELL_LIQUID
	} else {
		log.Fatalln("Unknown side", c.Side)
	}

	price, _ := c.Price.Float64()
	volume, _ := c.Volume.Float64()

	// return MakeWsLogRec(action, c.TimeStampE3, price, volume, strconv.Itoa(int(c.Id)))
	return MakeWsLogRec(action, c.TimeStampMs*1_000, price, volume, "")
}

func ParseLiquidationMessage(message json.RawMessage) (result string) {
	var data LiquidWs

	err := json.Unmarshal(message, &data)
	if err != nil {
		log.Fatalln("Fail to Parse Liquid message", err, message)
	}

	result = data.ToLog()

	return result
}
