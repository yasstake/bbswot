package main

import (
	"bbswot/bb"
	"bbswot/common"
	"flag"
	"fmt"
)

func splitTime(TimeE6 int64) (sec int64, msec int64) {
	sec = int64(TimeE6 / 1_000_000)
	msec = int64((TimeE6 % 1_000_000) / 10000)

	return sec, msec
}

func extractSingleFile(f string) {
	archiveLogMode := bb.CheckArchiveLog(f)

	if archiveLogMode {
		extractArchiveLog(f)
		return
	}

	stream := common.OpenFileReader(f)

	var recordNumber int64

	for stream.Scan() {
		rec := stream.Text()

		r1, r2, r3, r4, r5 := bb.ParseWsLogRec(rec)

		sec, msec := splitTime(r2)

		action := ""

		if r1 == common.TRADE_BUY {
			action = common.TRADE_BUY_STR
		} else if r1 == common.TRADE_SELL {
			action = common.TRADE_SELL_STR
		}

		size := int64(r4)
		if action != "" {
			fmt.Printf("%d.%02d,BTCUSD,%s,%d,%.0f,%s\n", sec, msec, action, size, r3, r5)
		}

		// TODO: for debug perpose [after debugging below lines shold be removed.]
		if 60_000 < r3 {
			fmt.Println("HIT")
		}

		recordNumber += 1
	}
}

func extractArchiveLog(file string) {
	stream := common.OpenFileReader(file)

	stream.Scan() // skip header line

	var recordNumber int64
	for stream.Scan() {
		rAction, rTimeMs, rPrice, rVolume, rOption := bb.ParseArchivedLogRec(stream.Text())

		action := ""

		sec, msec := splitTime(rTimeMs)

		if rAction == common.TRADE_BUY {
			action = common.TRADE_BUY_STR
		} else if rAction == common.TRADE_SELL {
			action = common.TRADE_SELL_STR
		}

		fmt.Printf("%d.%02d,BTCUSD,%s,%d,%.0f,%s\n", sec, msec, action, int(rVolume), rPrice, rOption)

		recordNumber += 1
	}
}

func main() {
	var enable_compress = flag.Bool("compress", true, "Enable log differential compress mode")

	flag.Parse()

	if *enable_compress {
		bb.EnableLogCompress()
	}

	files := flag.Args()

	for _, file := range files {
		extractSingleFile(file)
	}
}
