package bb

import (
	"bbswot/common"
	"encoding/json"
	"log"
	"strconv"
)

type LiquidRec struct {
    Id     int64       `json:"id"`           // "id":10334039,
	Price  json.Number `json:"price"`        // "price":49490.5}
	Volume json.Number `json:"qty"`          // "qty":1600,
	Symbol string      `json:"symbol"`       // "symbol":"BTCUSD",
	TimeStampMs   json.Number `json:"time"`  // "time":1630110808068,
	Side   string      `json:"side"`         // "side":"Sell",
}

func (c *LiquidRec) ToString() (r string) {
	msec, _ := c.TimeStampMs.Int64()

	r += common.MsToPrintDate(msec)
	r += c.Price.String() + " "
	r += c.Volume.String() + " "
	r += c.Side

	return r
}

func (c *LiquidRec) ToLog() (r string) {
	var action int
	if c.Side == "Sell" {
		action = common.TRADE_SELL_LIQUID
	} else if c.Side == "Buy" {
		action = common.TRADE_BUY_LIQUID
	} else {
		log.Println("unknown action side", c.Side)
	}

	price, _ := c.Price.Float64()
	volume, _ := c.Volume.Float64()
	time, _ := c.TimeStampMs.Int64()

	return MakeLogRec(action, time, price, volume, strconv.Itoa(int(c.Id)))       // enable ID version(for liquid)
}

type LiquidRecs []LiquidRec

func (c *LiquidRecs) ToLog() (result string) {
	l := len(*c)

	for i := 0; i < l; i++ {
		r := (*c)[i].ToLog()
		result += r
	}
	return result
}

// LiquidRequest Request Liquid rest API
func LiquidRequest(from_id *int64) (liq LiquidRecs, timeStampMs int64, err error) {
	url := "https://api.bybit.com/v2/public/liq-records?symbol=BTCUSD&limit=1000"

	if *from_id != 0 {
		url += "&from=" + strconv.Itoa(int(*from_id))
	}

	body, timeStampMs, err := RestRequest(url)
	if err != nil {
		log.Println(err)
	}

	liq, err = LiquidMessage(body)
	if err != nil {
		log.Println(err)
	}

	l := len(liq)
	if l != 0 {
		*from_id = liq[l-1].Id + 1
	}

	return liq, timeStampMs, err
}

// LiquidMessage Parse Liquid JSON message and returns Header and body
func LiquidMessage(message string) (liquid LiquidRecs, err error) {
	// var liquid []LiquidRec
	err = json.Unmarshal([]byte(message), &liquid)

	if err != nil {
		log.Println(err)
		return liquid, err
	}
	return liquid, nil
}

