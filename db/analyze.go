package db

import (
	"context"
	"github.com/labstack/gommon/log"
)

func TraverseDb(query string) {
	client := OpenClient()
	reader := NewQueryAPI(client)

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
