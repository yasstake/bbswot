package db

import (
	"bbswot/common"
	"log"
	"testing"
)

func TestOpenClient(t *testing.T) {
	client := OpenClient()

	log.Println(client)
}

func TestNewWriteApi(t *testing.T) {
	client := OpenClient()
	api := NewWriteAPI(client)

	log.Println(api)
}

func TestNewQueryApi(t *testing.T) {
	client := OpenClient()
	api := NewQueryAPI(client)

	log.Println(api)
}

func TestWriteTradePointDb(t *testing.T) {
	client := OpenClient()
	api := NewWriteAPI(client)

	// 5,1630209803585000,,10345016

	WriteTradePointDb(api, common.TRADE_SELL, 1630209803585123, 47862.5, 100.5, "0550ec6c-1222-5bd3-8993-61f05e4f87e9")
	api.Flush()
	client.Close()
}

func TestWriteBoardPointDb(t *testing.T) {
	client := OpenClient()
	api := NewWriteAPI(client)

	WriteBoardPointDb(api, common.UPDATE_BUY, 1630209803100000, 47862.5, 100.5)
	WriteBoardPointDb(api, common.UPDATE_SELL, 1630209803200000, 47863, 100)
	WriteBoardPointDb(api, common.PARTIAL, 1630209803300000, 0, 0)
	api.Flush()
	client.Close()
}

func TestUniqExecTimeStampE9(t *testing.T) {
	timeE9 := UniqExecTimeStampE9(1_000_123_456_789, "9fd6a16a-bfe5-580d-9c7c-000000001")

	log.Println(timeE9)
}

//
// b18cf816-ba56-5258-8c35-2a4e1066048b
func TestExecIdToInt(t *testing.T) {
	id := "b18cf816-ba56-5258-8c35-2a4e1066048b"

	idNumber := ExecIdToInt(id)

	log.Println(idNumber)
}

func TestLiquidIdToInt(t *testing.T) {
	id := "12345678"

	idNumber := ExecIdToInt(id)

	if idNumber != 345678 {
		t.Errorf("id mismatch %s %d", id, idNumber)
	}
	log.Println(idNumber)
}
