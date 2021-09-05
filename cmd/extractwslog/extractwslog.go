package main

import (
	"bbswot/bb"
	"bbswot/common"
	"fmt"
	"log"
)

func extractSingleFile(f string) {
	log.Println("[extract]", f)
	archiveLogMode := bb.CheckArchiveLog(f)

	if archiveLogMode {
		log.Println("Unsupported archive log")
		return
	}

	stream := common.OpenFileReader(f)

	var recordNumber int64

	for stream.Scan() {
		rec := stream.Text()

		r1, r2, r3, r4, r5 := bb.ParseWsLogRec(rec)
		fmt.Printf("%d,%d,%F,%F,%s\n", r1, r2, r3, r4, r5)

		recordNumber += 1
	}
}

func main() {
	extractSingleFile("./TEST_DATA/2021-09-02T10-12-16.log.gz")

	/*
		flag.Parse()

		files := flag.Args()

		for _, file := range files {
			extractSingleFile(file)
		}
	*/
}
