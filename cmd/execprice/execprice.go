package main

import (
	"bbswot/bb"
	"bbswot/common"
	"bbswot/stat"
	"flag"
	"fmt"
	"log"
	"strings"
)

func main() {
	/*
		log.Println("[Influxdb URL= ", db.INFLUXDB_URL)
		log.Println("[Influxdb KEY= ", db.INFLUXDB_KEY)
		log.Println("[Influxdb org= ", db.INFLUXDB_ORG)
		log.Println("[Influxdb bucket=", db.INFLUXDB_BUCKET)

	*/

	parseFile := func(file string) {
		fmt.Println("timestamp,buyprice,buyedge,sellprice,selledge")

		stream := common.OpenFileReader(file)

		var q stat.ExecQueue

		q.Init(1_000_000 * 300) // 300 sec = 6 min

		var lastBuyPrice float64
		var lastEdgeBuyPrice float64
		var lastSellPrice float64
		var lastEdgeSellPrice float64

		var recNumber int64

		for stream.Scan() {
			if strings.HasPrefix(stream.Text(), "timestamp") {
				continue
			}

			rAction, rTimeE6, rPrice, rVolume, _ := bb.ParseArchivedLogRec(stream.Text())

			q.Action(rAction, rTimeE6, rPrice, rVolume)

			if lastBuyPrice != q.BuyPrice || lastEdgeBuyPrice != q.BuyEdge || lastSellPrice != q.SellPrice || lastEdgeSellPrice != q.SellEdge {
				lastBuyPrice = q.BuyPrice
				lastEdgeBuyPrice = q.BuyEdge
				lastSellPrice = q.SellPrice
				lastEdgeSellPrice = q.SellEdge

				if q.BuyEdge != 0 && q.SellEdge != 0 {
					fmt.Println(q.EdgeTime, ",", q.BuyPrice, ",", q.BuyEdge, ",", q.SellPrice, ",", q.SellEdge)
				}
			}

			recNumber += 1
		}
	}

	flag.Parse()

	files := flag.Args()

	for _, file := range files {
		log.Println("[loading..]", file)
		parseFile(file)
	}
}
