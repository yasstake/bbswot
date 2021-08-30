package common

import (
	"log"
	"testing"
)

func TestOpenFileReader(t *testing.T) {
	const file = "../TEST_DATA/BTCUSD2021-08-29.csv.gz"

	r := OpenFileReader(file)

	var count int
	for r.Scan() {
		log.Print(r.Text())

		count += 1
		if 10 < count {
			break
		}
	}
}
