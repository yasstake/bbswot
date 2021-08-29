package db

import (
	"fmt"
	influxdb2 "github.com/influxdata/influxdb-client-go"
	"github.com/influxdata/influxdb-client-go/api"
	"os"
	"strconv"
	"time"
)

var INFLUXDB_KEY string

func init() {
	INFLUXDB_KEY = os.Getenv("INFLUXDB_KEY")
}

func OpenClient() influxdb2.Client {
	// Store the URL of your InfluxDB instance
	url := "http://localhost:8086"

	//client := influxdb2.NewClient(url, token)
	options := influxdb2.DefaultOptions()
	options.SetBatchSize(5000)
	options.SetPrecision(1)
	client := influxdb2.NewClientWithOptions(url, INFLUXDB_KEY, options)

	return client
}

func NewWriteApi(client influxdb2.Client) api.WriteApi {
	bucket := "btc"
	org := "bb"

	//writeAPI := client.WriteAPIBlocking(org, bucket)
	writeAPI := client.WriteAPI(org, bucket)

	return writeAPI
}

func NewQueryApi(client influxdb2.Client) api.QueryApi {
	api := client.QueryAPI("bb")

	return api
}

func WriteTradeDb(w api.WriteAPI, time_stamp time.Time, id int, side string, price float64, size float64) {

	id_string := strconv.Itoa(id)

	p := influxdb2.NewPoint("order",
		map[string]string{"id": id_string},
		map[string]interface{}{"tran": side, "price": price, "size": size},
		time_stamp)

	w.WritePoint(p)
}

func IdToInt(id string) int64 {
	var id1, id2, id3, id4, id5 int64

	fmt.Sscanf(id, "%x-%x-%x-%x-%x", &id1, &id2, &id3, &id4, &id5)

	return (id5) % 1_000_000
	//return (id1 + id2 + id3 + id4 + id5) % 1_000_000
	//return (id1*0x10000000 + id2*0x1000000 + id3*0x100000 + id4*0x10000 + id5) % 1_000_000
}
