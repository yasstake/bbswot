package bb

import (
	"bbswot/common"
	"bbswot/db"
	"log"
)

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
