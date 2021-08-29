package bb

import (
	"fmt"
	"testing"
)

func TestInstrument(t *testing.T) {
	result := ParseInstrumentSnapshot(GetMessageJson([]byte(INSTRUMENT_SNAPSHOT)), 1000)

	fmt.Println(result)
}


