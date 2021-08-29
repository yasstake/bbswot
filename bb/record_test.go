package bb

import (
	"fmt"
	"testing"
)

func TestMakeLogRec(t *testing.T) {
	rec := MakeLogRec(1, 2_001, 1.0, 1.0, "dummy")
	fmt.Println(rec)

	rec2 := MakeLogRec(1, 2_001, 1.0, 1.2, "dummy")
	fmt.Println(rec2)
}
