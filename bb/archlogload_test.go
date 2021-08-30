package bb

import "testing"

func TestArchiveLogLoad(t *testing.T) {
	const file = "../TEST_DATA/BTCUSD2021-08-29.csv.gz"

	ArchiveLogLoad(file)
}

func TestWsLogLoad(t *testing.T) {
	const file = "../TEST_DATA/2021-08-30T02-31-26.log.gz"

	WsLogLoad(file)
}
