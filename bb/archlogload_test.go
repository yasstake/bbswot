package bb

import "testing"

func TestArchiveLogLoad(t *testing.T) {
	const file = "../TEST_DATA/BTCUSD2021-08-29.csv.gz"

	ArchiveLogLoad(file)
}
