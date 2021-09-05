package main

import (
	"bbswot/bb"
	"bbswot/common"
	"flag"
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

		// TODO: for debug perpose [after debugging below lines shold be removed.]
		if 60_000 < r3 {
			fmt.Println("HIT")
		}

		recordNumber += 1
	}
}

func main() {
	var enable_compress = flag.Bool("compress", true, "Enable log differential compress mode")

	flag.Parse()

	fmt.Println("flag", *enable_compress)

	if *enable_compress {
		fmt.Printf("[Enable Compress]")
		bb.EnableLogCompress()
	}

	files := flag.Args()

	for _, file := range files {
		extractSingleFile(file)
	}
}
