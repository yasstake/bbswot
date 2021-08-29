package common

import (
	"fmt"
	"testing"
	"time"
)


func TestDivideSecAndMs(t *testing.T) {
	sec, msec := DivideSecAndMs(2_001)

	if sec != 2 {
		t.Errorf("Does not match sec %d %d", sec, 2)
	}

	if msec != 1 {
		t.Errorf("Does not match msec %d %d", msec, 1)
	}
}

func TestMsToTime(t *testing.T) {
	timeObj := MsToTime(2_001)   // 1970/1/1 9:0:2 (JST)

	fmt.Println(timeObj)
}

func TestParseIsoTimeToMs(t *testing.T) {
	const timeString = "1970-01-01T00:00:02.001Z"

	timeMs := ParseIsoTimeToMs(timeString)

	if timeMs != 2_001 {
		t.Errorf("Time Parse mismatch %d, %d", 2001, timeMs)
	}
}

func TestDateMs(t *testing.T) {
	tradeTimeMs := int64(1619398389868)
	//,"timestamp":"2021-04-26T00:53:09.000Z"

	tm1 := time.Unix(tradeTimeMs/1_000, (tradeTimeMs%1_000)*1_000_000)
	fmt.Println(tm1)

	tm2 := MsToTime(tradeTimeMs)
	fmt.Println(tm2)

	if tm1 != tm1 {
		t.Errorf("does not match")
	}
}

