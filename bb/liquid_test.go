package bb

import (
	"log"
	"testing"
)

func TestLiquidRec_ToLog(t *testing.T) {
	TEST_MESSAGE := `[{"id":10337551,"qty":10,"side":"Buy","time":1630161605869,"symbol":"BTCUSD","price":48270}, {"id":10337552,"qty":20,"side":"Sell","time":1630161605870,"symbol":"BTCUSD","price":48271}]`

	recs, _ := LiquidMessage(TEST_MESSAGE)

	log.Println(recs.ToLog())
}

func TestLiquidRequest(t *testing.T) {
	var currentId int64

	rec, timeStampMs, err := LiquidRequest(&currentId)

	log.Println(rec, timeStampMs, err)
}

func TestLiquidRequestAndMessage(t *testing.T) {
	var currentId int64
	recs, timeStampMs, err := LiquidRequest(&currentId)

	log.Println(recs.ToLog(), timeStampMs, err)
}
