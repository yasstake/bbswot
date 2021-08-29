package bb

import (
	"fmt"
	"testing"
)

func TestParseOrderBookMessage(t *testing.T) {
	m := ParseMessage([]byte(ORDER_BOOK_SNAP_RECORD))
	message := ParseOrderBookMessage(m)
	fmt.Println(message)

	fmt.Println("------------")

	m = ParseMessage([]byte(ORDER_BOOK_DELTA_RECORD))
	message = ParseOrderBookMessage(m)
	fmt.Println(message)
}
