package main

import (
	"bbswot/bb"
	"testing"
)

func TestLoadWsFile(t *testing.T) {
	file := "../../TEST_DATA/2021-09-02T18-22-38.log.gz"

	bb.EnableLogCompress()
	extractSingleFile(file)
}
