package common

import (
	"log"
	"strconv"
	"time"
)

// DivideSecAndMs Divide  Millsecond to Sec part and msec part
func DivideSecAndMs(timeMs int64) (sec int64, msec int64) {
	sec = timeMs / 1_000
	msec = timeMs % 1_000

	return sec, msec
}

// MsToTime convert unix time(milliseconds) to TimeStampE3 object
func MsToTime(msec int64) time.Time {
	tm := time.Unix(msec/1_000, (msec%1_000)*1_000_000)

	return tm
}

func TimeSecToE6(sec int64, msec int64) int64 {
	return sec*1_000_000 + msec*1_000
}

// MsToPrintDate Convert msec to printable format [for debugging]
func MsToPrintDate(msec int64) string {
	t := MsToTime(msec)

	s := t.UTC().String() + "(" + strconv.Itoa(int(msec)) + ")"

	return s
}

func ParseIsoTime(t string) (result time.Time) {
	const layout = "2006-01-02T15:04:05Z"
	result, err := time.Parse(layout, t)

	if err != nil {
		log.Println("Dateformat error in log ", err, t)
	}

	return result
}

func ParseIsoTimeToE6(isoTime string) (timeMs int64) {
	timeObj := ParseIsoTime(isoTime)

	timeMs = timeObj.UnixNano() / 1_000

	return timeMs
}
