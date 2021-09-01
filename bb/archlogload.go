package bb

import (
	"bbswot/common"
	"bbswot/db"
	"github.com/influxdata/influxdb-client-go/api"
	"log"
	"strconv"
	"strings"
)

// ArchiveLogLoad
// Load Archive log from Bybit website
func ArchiveLogLoad(file string) {
	client := db.OpenClient()
	defer client.Close()
	writer := db.NewWriteAPI(client)
	defer writer.Flush()

	stream := common.OpenFileReader(file)

	stream.Scan() // skip header line

	var recordNumber int64
	for stream.Scan() {
		rAction, rTimeMs, rPrice, rVolume, rOption := ParseArchivedLogRec(stream.Text())

		db.WriteTradePointDb(writer, rAction, rTimeMs, rPrice, rVolume, rOption)

		recordNumber += 1
	}

	log.Printf("Log loaded %s (%d)\n", file, recordNumber)
}

func flushBoardBuffer(writer api.WriteAPI, timeE6 int64) {
	db.WriteBoardPointDb(writer, common.PARTIAL, timeE6, 0, 0)

	writeBoardBuffer(writer, common.UPDATE_BUY, timeE6, buyBoardBuffer)
	writeBoardBuffer(writer, common.UPDATE_SELL, timeE6, sellBoardBuffer)

	buyBoardBuffer.Reset()
	sellBoardBuffer.Reset()
}

func writeBoardBuffer(writer api.WriteAPI, action int, timeE6 int64, board Board) {
	for price, volume := range board {
		db.WriteBoardPointDb(writer, action, timeE6, price, volume)
	}
}

// WsLogLoad
// Load Web service log to influxdb
func WsLogLoad(file string) {
	client := db.OpenClient()
	defer client.Close()
	writer := db.NewWriteAPI(client)
	defer writer.Flush()

	stream := common.OpenFileReader(file)

	var recordNumber int64
	var execNumber int64
	var boardNumber int64
	var oiNumber int64
	var frNumber int64

	log.Println("---start--")
	buyBoardBuffer.Reset()
	sellBoardBuffer.Reset()
	var lastTimeE6 int64
	const timeIntervalE6 = 1_000_000 * 60 * 3

	for stream.Scan() {
		rec := stream.Text()
		rAction, rTimeE6, rPrice, rVolume, rOption := ParseWsLogRec(rec)

		if rAction == common.TRADE_BUY || rAction == common.TRADE_SELL {
			db.WriteTradePointDb(writer, rAction, rTimeE6, rPrice, rVolume, rOption)
			execNumber += 1
		} else if rAction == common.PARTIAL || rAction == common.UPDATE_BUY || rAction == common.UPDATE_SELL {
			if rAction == common.PARTIAL {
				sellBoardBuffer.Reset()
				buyBoardBuffer.Reset()
			}
			if rAction == common.UPDATE_BUY {
				buyBoardBuffer.Add(rPrice, rVolume)
			}
			if rAction == common.UPDATE_SELL {
				sellBoardBuffer.Add(rPrice, rVolume)
			}

			// TODO: write snapshot every 3-5 min
			db.WriteBoardPointDb(writer, rAction, rTimeE6, rPrice, rVolume)
			boardNumber += 1
		} else if rAction == common.OPEN_INTEREST {
			db.WriteOpenInterests(writer, rTimeE6, int64(rVolume))
			oiNumber += 1
		} else if rAction == common.FUNDING_RATE {
			db.WriteFundingRate(writer, rTimeE6, rVolume)
			frNumber += 1
		} else if rAction == common.PREDICTED_FUNDING_RATE {
			timeE6, _ := strconv.ParseInt(rOption, 10, 64)
			db.WritePredictedFundingRate(writer, rTimeE6, rVolume, timeE6)
			log.Println("[NEXT FR", rTimeE6, rVolume, rOption)
		}

		if timeIntervalE6 < rTimeE6-lastTimeE6 && boardNumber != 0 {
			flushBoardBuffer(writer, rTimeE6)

			lastTimeE6 = rTimeE6
		}

		recordNumber += 1
	}

	log.Printf("Log loaded %s exec=%d  board=%d oi=%d fr=%d (total=%d)\n", file, execNumber, boardNumber, oiNumber, frNumber, recordNumber)
}

// CheckArchiveLog
// read first line and check whether the log file is archive log(true) or WS log(false)
func CheckArchiveLog(file string) (result bool) {
	const archiveHeader = "timestamp,symbol,side,size,price,tickDirection,trdMatchID,grossValue,homeNotional,foreignNotional"
	result = false
	stream := common.OpenFileReader(file)

	// read header
	if stream.Scan() {
		header := stream.Text()

		if strings.HasPrefix(header, archiveHeader) {
			result = true
		}
	}

	return result
}
