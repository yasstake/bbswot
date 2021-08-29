package bb

import (
	common "bbswot/common"
	"fmt"
	"strconv"
)

/*
var (
	last_time  time.TimeStampMs
	last_price float64
)
*/



func MakeLogRec(action int, timeMs int64, price float64, volume float64, option string) (result string) {
	sec, msec :=  common.DivideSecAndMs(timeMs)

	result = fmt.Sprintf("%d,%d%03d", action, sec, msec)

	priceString :=  strconv.FormatFloat(price, 'f', -1, 64)
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

type LogPointer struct {
	LastMilSec int64
	LastPrice float64
}

func CompressLogRec(lastPointer LogPointer, action int, timeMs int64, price float64, volume float64, option string) (c_action int, c_timeMs int64, c_price float64, c_volume float64, c_option string) {

	return c_action, c_timeMs, c_price, c_volume,c_option
}
