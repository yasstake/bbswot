package main

import (
	"bbswot/bb"
)

func main() {
	const file = "./TEST_DATA/2021-08-30T02-31-26.log.gz"

	bb.WsLogLoad(file)
}
