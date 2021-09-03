package bb

import (
	"fmt"
	"github.com/labstack/gommon/log"
	"testing"
)

func TestRest(t *testing.T) {
	sellBoardBuffer.Reset()

	if sellBoardBuffer.Len() != 0 {
		log.Error(sellBoardBuffer)
	}
}

func TestAdd(t *testing.T) {
	sellBoardBuffer.Reset()

	sellBoardBuffer.Set(10.0, 10.0)
	fmt.Println(sellBoardBuffer.Len())
	sellBoardBuffer.Set(10.0, 10.1)
	fmt.Println(sellBoardBuffer.Len())
	sellBoardBuffer.Set(10.1, 10.1)

	fmt.Println(sellBoardBuffer.Len())
	fmt.Println(sellBoardBuffer)

	sellBoardBuffer.Set(10.1, 0)
	fmt.Println(sellBoardBuffer.Len())

	fmt.Println(sellBoardBuffer)

}

func TestBoard_GetHighPrice(t *testing.T) {
	var board Board
	board.Reset()

	board.Set(10.0, 10.0)
	board.Set(12.0, 10.0)
	board.Set(9.0, 10.0)

	high := board.CalcHighPrice()
	log.Print("Price=", high)
	if high != 12 {
		t.Error("mismatch")
	}
}

func TestBoard_GetLowPrice(t *testing.T) {
	var board Board
	board.Reset()

	board.Set(10.0, 10.0)
	board.Set(12.0, 10.0)
	board.Set(9.0, 10.0)

	low := board.CalcLowPrice()
	log.Print("Price=", low)
	if low != 9 {
		t.Error("mismatch")
	}

	board.Set(9.0, 0)
	low = board.CalcLowPrice()
	log.Print("Price=", low)

	if low != 10 {
		t.Error("mismatch")
	}
}
