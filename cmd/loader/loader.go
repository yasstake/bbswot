package main

import (
	"bbswot/bb"
	"flag"
	"log"
)

func main() {
	flag.Parse()
	files := flag.Args()

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
