package bb

import (
	"log"
	"testing"
)

func TestParseTradeMessage(t *testing.T) {
	m := ParseMessage([]byte(TRADE_RECORD))
	result := ParseTradeMessage(m)

	log.Println(result)
}
