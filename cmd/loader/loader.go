package main

import (
	"bbswot/bb"
	"bbswot/db"
	"flag"
	"log"
)

func main() {
	flag.Parse()
	files := flag.Args()

	log.Println("[Influxdb URL= ", db.INFLUXDB_URL)
	log.Println("[Influxdb KEY= ", db.INFLUXDB_KEY)
	log.Println("[Influxdb org= ", db.INFLUXDB_ORG)
	log.Println("[Influxdb bucket=", db.INFLUXDB_BUCKET)

	for _, file := range files {
		log.Println("[loading..]", file)
		archiveLogMode := bb.CheckArchiveLog(file)
		if archiveLogMode {
			log.Println("Loading Archive Log file")
			bb.ArchiveLogLoad(file)
		} else {
			log.Println("Loading WS Log file")
			bb.WsLogLoad(file)
		}
	}
}
