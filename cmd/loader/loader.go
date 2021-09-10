package main

import (
	"bbswot/bb"
	"bbswot/db"
	"flag"
	"log"
)

func main() {
	log.Println("[Influxdb URL= ", db.INFLUXDB_URL)
	log.Println("[Influxdb KEY= ", db.INFLUXDB_KEY)
	log.Println("[Influxdb org= ", db.INFLUXDB_ORG)
	log.Println("[Influxdb bucket=", db.INFLUXDB_BUCKET)

	var enable_compress = flag.Bool("compress", true, "Enable log differential compress mode")
	var deleteMode = flag.Bool("delete", false, "delete db for log time frame")
	flag.Parse()

	nFlags := flag.NFlag()
	files := flag.Args()[nFlags:]
	log.Println("FLAG", nFlags, files)

	if *deleteMode {
		log.Printf("[Delete mode]")

		for _, file := range files {
			bb.DeleteExecForArchiveLog(file)
		}
		return
	}

	if *enable_compress {
		log.Printf("[Enable Compress]")
		bb.EnableLogCompress()
	}

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
