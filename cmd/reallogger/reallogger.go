package main

import (
	"bbswot/bb"
	"bbswot/db"
	"github.com/influxdata/influxdb-client-go/api"
)

type DbWriter struct {
	writer api.WriteAPI
}

func (w *DbWriter) Open() {
	db.SetInfluxDbBatchSize(10)
	client := db.OpenClient()
	(*w).writer = db.NewWriteAPI(client)
}

func (w *DbWriter) Write(b []byte) (n int, err error) {
	bb.LoadWsRecord(w.writer, string(b))

	return len(b), nil
}

func (w *DbWriter) Close() error {
	w.writer.Flush()
	w.writer.Close()

	return nil
}

func main() {
	bb.DisableLogCompress()
	exitWait := 0
	flagFile := ""

	var dbWriter DbWriter
	dbWriter.Open()

	bb.Connect(flagFile, &dbWriter, exitWait)

	dbWriter.Close()
}
