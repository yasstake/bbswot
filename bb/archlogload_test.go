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

func TestCheckArchiveLog(t *testing.T) {
	const archiveLog = "../TEST_DATA/BTCUSD2021-08-29.csv.gz"
	const wsLog = "../TEST_DATA/2021-08-30T02-31-26.log.gz"

	r := CheckArchiveLog(archiveLog)
	if !r {
		t.Error("Must be archive log", archiveLog)
	}

	r = CheckArchiveLog(wsLog)
	if r {
		t.Error("Must be ws log", wsLog)
	}
}
