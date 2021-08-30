package bb

import (
	"fmt"
	"log"
	"testing"
)

// for archived log

func TestParseArchivedLogRec(t *testing.T) {
	const bblogTestData = "1630281599.872,BTCUSD,Buy,5000,48815.5,ZeroMinusTick,fc7b42fa-0b1a-50d1-9a9c-cfc70a5a9b66,10242753.0,5000,0.10242753"

	action, orgTimeMs, orgPrice, volume, option := ParseArchivedLogRec(bblogTestData)

	log.Println(action, orgTimeMs, orgPrice, volume, option)
}

// for WS log records

func TestMakeLogRec(t *testing.T) {
	rec := MakeWsLogRec(1, 2_001, 1.0, 1.0, "dummy")
	fmt.Println(rec)

	rec2 := MakeWsLogRec(1, 2_001, 1.0, 1.2, "dummy")
	fmt.Println(rec2)
}

func TestParseLogRec(t *testing.T) {
	s1 := "5,1630205952783,48338,28,10342242"
	s2 := "5,0,-11,27210,10342243"
	s3 := "5,7,27.5,40000,10342250"
	s4 := "5,-5,-9.5,989,10342268"

	rAction, rTimeMs, rPrice, rVolume, rOption := ParseWsLogRec(s1)
	log.Println("[DECODE]", rAction, rTimeMs, rPrice, rVolume, rOption)

	rAction, rTimeMs, rPrice, rVolume, rOption = ParseWsLogRec(s2)
	log.Println("[DECODE]", rAction, rTimeMs, rPrice, rVolume, rOption)

	rAction, rTimeMs, rPrice, rVolume, rOption = ParseWsLogRec(s3)
	log.Println("[DECODE]", rAction, rTimeMs, rPrice, rVolume, rOption)

	rAction, rTimeMs, rPrice, rVolume, rOption = ParseWsLogRec(s4)
	log.Println("[DECODE]", rAction, rTimeMs, rPrice, rVolume, rOption)
}

// For bulk download transaction log records
// wget https://public.bybit.com/trading/BTCUSD/BTCUSD2021-08-25.csv.gz
// Format:
// timestamp(e6)    ,symbol,side,size,price,  tickDirection,trdMatchID                           ,grossValue,homeNotional,foreignNotional
// 1621036794.507791,BTCUSD,Sell,10  ,49910.0,ZeroMinusTick,c00e7c56-feb6-51ea-a24e-c827583fb5c5 ,20036.0   ,10          ,0.00020036
