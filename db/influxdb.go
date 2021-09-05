package db

import (
	"bbswot/common"
	"fmt"
	influxdb2 "github.com/influxdata/influxdb-client-go"
	"github.com/influxdata/influxdb-client-go/api"
	"github.com/labstack/gommon/log"
	"os"
	"strconv"
	"time"
)

var INFLUXDB_KEY string
var INFLUXDB_BUCKET string
var INFLUXDB_ORG string
var INFLUXDB_URL string
var INFLUXDB_BATCHSIZE uint

func init() {
	INFLUXDB_KEY = os.Getenv("INFLUX_TOKEN")
	if INFLUXDB_KEY == "" {
		log.Printf("[warning] no $INFLUX_TOKEN")
	}

	INFLUXDB_BUCKET = "btc"
	INFLUXDB_ORG = "bb"
	url := os.Getenv("INFLUX_HOST")
	if url != "" {
		INFLUXDB_URL = url
	} else {
		INFLUXDB_URL = "http://localhost:8086"
	}

	INFLUXDB_BATCHSIZE = 5000
}

func SetInfluxDbBatchSize(size uint) {
	INFLUXDB_BATCHSIZE = size
}

func OpenClient() influxdb2.Client {
	// Store the URL of your InfluxDB instance
	options := influxdb2.DefaultOptions()
	options.SetBatchSize(INFLUXDB_BATCHSIZE)
	options.SetPrecision(1)
	client := influxdb2.NewClientWithOptions(INFLUXDB_URL, INFLUXDB_KEY, options)

	return client
}

func NewWriteAPI(client influxdb2.Client) api.WriteAPI {
	writeAPI := client.WriteAPI(INFLUXDB_ORG, INFLUXDB_BUCKET)

	return writeAPI
}

func NewQueryAPI(client influxdb2.Client) api.QueryAPI {
	api := client.QueryAPI("bb")

	return api
}

func WriteBoardPointDb(w api.WriteAPI, action int, timestampE6 int64, price float64, size float64) {
	t := time.Unix(0, timestampE6*1_000)

	var side string
	priceStr := strconv.FormatInt(int64(price*10), 10)

	if action == common.UPDATE_BUY {
		side = "Buy"
	} else if action == common.UPDATE_SELL {
		side = "Sell"
	} else if action == common.PARTIAL {
		side = "Partial"
		priceStr = "PARTIAL"
		log.Print("[PARTIAL]", common.TimeE6ToString(timestampE6))
	}

	p := influxdb2.NewPoint("board",
		map[string]string{"s": side, "p": priceStr},
		map[string]interface{}{"side": side, "price": price, "size": size},
		t)

	w.WritePoint(p)
}

func WriteTradePointDb(w api.WriteAPI, action int, timestampE6 int64, price float64, size float64, execId string) {
	uniqTime := UniqExecTimeStampE9(timestampE6, execId)
	t := time.Unix(0, uniqTime)

	// TODO: remove after debugg
	if 60_000 < price {
		log.Warn("Too much price", price)
	}

	var side string
	if action == common.TRADE_BUY {
		side = common.TRADE_BUY_STR
	} else if action == common.TRADE_SELL {
		side = common.TRADE_SELL_STR
	} else if action == common.TRADE_BUY_LIQUID {
		side = common.TRADE_BUY_LIQUID_STR
	} else if action == common.TRADE_SELL_LIQUID {
		side = common.TRADE_SELL_LIQUID_STR
	} else {
		log.Error("unknown action no", action)
	}

	p := influxdb2.NewPoint("exec",
		//map[string]string{"s": side},
		map[string]string{},
		map[string]interface{}{"side": side, "price": price, "size": size},
		t)

	w.WritePoint(p)
}

// WriteOpenInterests
// Write Open interests at the point
func WriteOpenInterests(w api.WriteAPI, timestampE6 int64, openInterest int64) {
	t := time.Unix(0, FloorTimeStampE6ToE9(timestampE6))

	p := influxdb2.NewPoint("oi",
		map[string]string{},
		map[string]interface{}{"size": openInterest},
		t)

	w.WritePoint(p)
}

func WriteFundingRate(w api.WriteAPI, timestampE6 int64, fundingRate float64) {
	t := time.Unix(0, FloorTimeStampE6ToE9(timestampE6))

	p := influxdb2.NewPoint("funding",
		map[string]string{"current": ""},
		map[string]interface{}{"rate": fundingRate},
		t)

	w.WritePoint(p)
}

func WritePredictedFundingRate(w api.WriteAPI, timestampE6 int64, fundingRate float64, nextTimeE6 int64) {
	t := time.Unix(0, FloorTimeStampE6ToE9(timestampE6))
	next_time := time.Unix(0, FloorTimeStampE6ToE9(nextTimeE6))

	p := influxdb2.NewPoint("funding",
		map[string]string{"predict": ""},
		map[string]interface{}{"rate": fundingRate, "next_time": next_time},
		t)

	w.WritePoint(p)
}

const FloorUnit = 1_000
const ConvertE6toE9 = 1_000

func FloorTimeStampE6ToE9(timeStampE6 int64) (timeE9 int64) {
	t := int64(timeStampE6 / FloorUnit)
	t = t * FloorUnit * ConvertE6toE9

	return t
}

func UniqExecTimeStampE9(timestampE6 int64, id string) (timeE9 int64) {
	timeE9 = FloorTimeStampE6ToE9(timestampE6)
	timeE9 = timeE9 + ExecIdToInt(id)

	return timeE9
}

const IdWidth = 1_000_000

// ExecIdToInt
// 9fd6a16a-bfe5-580d-9c7c-0168aeb4c93e      // buy or sell
// 468925                                    // liquid
// Parse execute id and convert it to 0-1_000_000 numbers
func ExecIdToInt(id string) (idInt int64) {
	l := len(id)

	if 7 <= l {
		var id1, id2, id3, id4, id5 int64
		fmt.Sscanf(id, "%x-%x-%x-%x-%x", &id1, &id2, &id3, &id4, &id5)

		idInt = id5 % IdWidth
	} else {
		var id6 int64
		fmt.Sscanf(id, "%d", id6)

		idInt = id6 % IdWidth
	}
	return idInt
}
