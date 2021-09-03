package bb

import (
	"bbswot/db"
	"context"
	"github.com/labstack/gommon/log"
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
  |> range(start: -5d)
  |> filter(fn: (r) => r["_measurement"] == "board")
  |> filter(fn: (r) => r["_field"] == "price" or r["_field"] == "size" or r["_field"] == "side")
  |> pivot(rowKey: ["_time"], columnKey: ["_field"], valueColumn: "_value")
    `

	client := db.OpenClient()
	reader := db.NewQueryAPI(client)

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
