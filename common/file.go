package common

import (
	"bufio"
	"compress/gzip"
	"log"
	"os"
	"strings"
)

func OpenFileReader(file string) (reader *bufio.Scanner) {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}

	isCompress := strings.HasSuffix(file, ".gz")

	if isCompress {
		zipFile, _ := gzip.NewReader(f)
		reader = bufio.NewScanner(zipFile)
	} else {
		reader = bufio.NewScanner(f)
	}

	return reader
}
