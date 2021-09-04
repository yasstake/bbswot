package bb

import (
	"bbswot/db"
	"context"
	"github.com/labstack/gommon/log"
	"time"
)

func TraverseDb(query string) {
	client := db.OpenClient()
	reader := db.NewQueryAPI(client)

	result, err := reader.Query(context.Background(), query)

	if err != nil {
		log.Error(err)
	}

	var count int
	for result.Next() {

		// values := record.Values()
		/*
			if result.TableChanged() {
				log.Printf("Table changed%s\n", result.TableMetadata().String() )
			}
		*/
		count += 1
		log.Printf("Value: %v", result.Record().Values())
	}

	log.Printf("Recods=%d\n", count)
	client.Close()
}

func CountDb(query string) {
	client := db.OpenClient()
	reader := db.NewQueryAPI(client)

	query += "|> count()"
	result, err := reader.Query(context.Background(), query)

	if err != nil {
		log.Error(err)
	}

	var count int
	for result.Next() {
		if result.TableChanged() {
			log.Printf("Table changed%s\n", result.TableMetadata().String())
		}
		count += 1
		log.Printf("Value: %v", result.Record().Values())
	}

	log.Printf("Recods=%d\n", count)
	client.Close()
}

func FindEdgePrice() {
	query := `
from(bucket: "btc")
  |> range(start: 1970-01-01T00:00:00Z)
  |> filter(fn: (r) => r["_measurement"] == "board")
  |> filter(fn: (r) => r["_field"] == "price" or r["_field"] == "size" or r["_field"]=="side")
  |> drop(columns: ["_start", "_stop", "_measurement"])
  |> pivot(rowKey: ["_time", "s", "p"], columnKey:["_field"], valueColumn: "_value")
  |> drop(columns: ["p", "s"])
  |> sort(columns: ["_time", "price"])
    `

	client := db.OpenClient()
	reader := db.NewQueryAPI(client)

	result, err := reader.Query(context.Background(), query)

	if err != nil {
		log.Error(err)
	}

	var count int
	var sellBoard Board
	var buyBoard Board

	sellBoard.Reset()
	buyBoard.Reset()

	// TODO: analyze last tick price
	//	var lastLow float64
	//var lastHigh float64
	//	lastTick float64

	for result.Next() {
		if result.TableChanged() {
			log.Printf("Table changed%s\n", result.TableMetadata().String())
		}
		count += 1
		values := result.Record().Values()
		tick := values["_time"].(time.Time)
		side := values["side"].(string)
		price := values["price"].(float64)
		size := values["size"].(float64)

		log.Print(tick)

		if side == "Partial" {
			log.Printf("---PARTIAL----")
			buyBoard.Reset()
			sellBoard.Reset()
		} else if side == "Buy" {
			buyBoard.Set(price, size)
		} else if side == "Sell" {
			sellBoard.Set(price, size)
		} else {
			log.Error("Unknown side", side)
		}

	}

	log.Printf("Recods=%d\n", count)
	client.Close()
}
