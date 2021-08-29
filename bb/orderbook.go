package bb

import (
	"bbswot/common"
	"encoding/json"
	"log"
)

type Order struct {
	Price  json.Number `json:"price"`     // "price":"48608.50"
	Symbol string      `json:"symbol"`    // "symbol":"BTCUSD"
	Id     int64       `json:"id"`        // "id":486085000
	Side   string      `json:"side"`      // "side":"Sell"
	Size   json.Number `json:"size"`      // "size":409566
	TimeStampMs   int64
}

func (c *Order) ToLog() string {
	var action int
	if c.Side == "Buy" {
		action = common.UPDATE_BUY
	} else if c.Side == "Sell" {
		action = common.UPDATE_SELL
	} else {
		log.Fatalln("Unknown side", c.Side)
	}

	price, _ := c.Price.Float64()
	volume, _ := c.Size.Float64()

	// return MakeLogRec(action, c.TimeStampMs, price, volume, strconv.Itoa(int(c.Id)))
	return MakeLogRec(action, c.TimeStampMs, price, volume, "")
}

type SnapShot []Order

type Delta struct {
	Delete []Order `json:"delete"`
	Update []Order `json:"update"`
	Insert []Order `json:"insert"`
}

func ParseOrderBookMessage(message Message) (result string) {
	switch message.Type {
	case "snapshot":
		return ParseOrderBookSnapshot(message.Data, message.TimeStampMs)
	case "delta":
		return ParseOrderBookDelta(message.Data, message.TimeStampMs)
	}

	log.Fatalln("Unknown Message type", message.Type)

	return ""
}

func ParseOrderBookSnapshot(message json.RawMessage, timeMs int64) (result string) {
	var data SnapShot

	err := json.Unmarshal(message, &data)
	if err != nil {
		log.Fatalln("Fail to pase message", err, message)
	}

	l := len(data)
	if l == 0 {
		return ""
	}

	result = MakeLogRec(common.PARTIAL, timeMs, 0, 0, "")

	for i := 0; i < l; i++ {
		data[i].TimeStampMs = timeMs
		result += data[i].ToLog()
	}

	return result
}

func ParseOrderBookDelta(message json.RawMessage, timeMs int64) (result string) {
	var data Delta
	result = ""

	err := json.Unmarshal(message, &data)
	if err != nil {
		log.Fatalln("Fail to pase message", err, message)
	}

	l := len(data.Insert)
	for i := 0; i < l; i++ {
		data.Insert[i].TimeStampMs = timeMs
		result += data.Insert[i].ToLog()
	}

	l = len(data.Update)
	for i := 0; i < l; i++ {
		data.Update[i].TimeStampMs = timeMs
		result += data.Update[i].ToLog()
	}

	l = len(data.Delete)
	for i := 0; i < l; i++ {
		data.Delete[i].TimeStampMs = timeMs
		result += data.Delete[i].ToLog()
	}

	return result
}
