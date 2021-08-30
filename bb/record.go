package bb

import (
	"bbswot/common"
	"bytes"
	"encoding/csv"
	"fmt"
	"github.com/labstack/gommon/log"
	"math"
	"strconv"
)

/*
var (
	last_time  time.TimeStampE3
	last_price float64
)
*/

var (
	cacheLastTime  int64
	cacheLastPrice float64
	doCompress     bool
)

func EnableLogCompress() {
	doCompress = true
}

//////////// Archived Log records ///////////

func ParseArchivedLogRec(rec string) (rAction int, rTimeE6 int64, rPrice float64, rVolume float64, rTransactionId string) {
	buffer := bytes.NewBufferString(rec)

	reader := csv.NewReader(buffer)

	r, err := reader.Read()
	if err != nil {
		log.Error(err)
	}

	if len(r) < 7 {
		log.Error("too shot format", r)
	}
	var sec, msec int64
	fmt.Sscanf(r[0], "%d.%d", &sec, &msec)
	rTimeE6 = common.TimeSecToE6(sec, msec)

	actionName := r[2]
	if actionName == common.TRADE_BUY_STR {
		rAction = common.TRADE_BUY
	} else if actionName == common.TRADE_SELL_STR {
		rAction = common.TRADE_SELL
	} else {
		log.Error("unknown action", actionName)
	}

	rVolume, err = strconv.ParseFloat(r[3], 64)
	if err != nil {
		log.Error("Format error", r[3])
	}
	rPrice, err = strconv.ParseFloat(r[4], 64)

	if 7 <= len(r) {
		rTransactionId = r[6]
	} else {
		rTransactionId = ""
	}

	return rAction, rTimeE6, rPrice, rVolume, rTransactionId
}

//////////// WS log records ///////////////

func MakeWsLogRec(action int, orgTimeMs int64, orgPrice float64, volume float64, option string) (result string) {
	var timeMs int64
	var price float64

	if doCompress {
		timeMs = orgTimeMs - cacheLastTime
		cacheLastTime = orgTimeMs

		price = orgPrice - cacheLastPrice
		cacheLastPrice = orgPrice
	} else {
		timeMs = orgTimeMs
		price = orgPrice
	}
	result = fmt.Sprintf("%d,%d", action, timeMs)

	priceString := strconv.FormatFloat(price, 'f', -1, 64)
	result += ","
	result += priceString

	volumeString := strconv.FormatFloat(volume, 'f', -1, 64)
	result += ","
	result += volumeString

	if option != "" {
		result += ","
		result += option
	}

	result += "\n"

	return result
}

//
// Sample record
// 5,1630205955793000,48253,209,10342343
func ParseWsLogRec(rec string) (rAction int, rTimeE6 int64, rPrice float64, rVolume float64, rOption string) {
	buffer := bytes.NewBufferString(rec)

	reader := csv.NewReader(buffer)

	r, err := reader.Read()
	if err != nil {
		log.Error(err)
	}

	if len(r) < 4 {
		log.Error("too shot format", r)
	}

	rAction, err = strconv.Atoi(r[0])
	if err != nil {
		log.Error("Id Parse Error", err, r[0])
	}

	timeE6, err := strconv.ParseInt(r[1], 10, 64)
	if err != nil {
		log.Error("TimeE6 Parse Error", err, r[1])
	}

	price, err := strconv.ParseFloat(r[2], 64)
	if err != nil {
		log.Error("Price  error", r[2])
	}

	if 1_000_000_000_000 < math.Abs(float64(timeE6)) {
		// uncompressed mode
		cacheLastTime = timeE6
		rTimeE6 = timeE6

		cacheLastPrice = price
		rPrice = price
	} else {
		// Diff mode (compressed)
		rTimeE6 = cacheLastTime + timeE6
		cacheLastTime = rTimeE6

		rPrice = cacheLastPrice + price
		cacheLastPrice = rPrice
	}

	rVolume, err = strconv.ParseFloat(r[3], 64)
	if err != nil {
		log.Error("Price  error", r[3])
	}

	if 5 <= len(r) {
		rOption = r[4]
	} else {
		rOption = ""
	}

	return rAction, rTimeE6, rPrice, rVolume, rOption
}
