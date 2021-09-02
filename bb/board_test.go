package bb

import (
	"fmt"
	"github.com/labstack/gommon/log"
	"testing"
)

func TestRest(t *testing.T) {
	sellBoardBuffer.Reset()

	if len(sellBoardBuffer) != 0 {
		log.Error(sellBoardBuffer)
	}
}

func TestAdd(t *testing.T) {
	sellBoardBuffer.Reset()

	sellBoardBuffer.Set(10.0, 10.0)
	fmt.Println(len(sellBoardBuffer))
	sellBoardBuffer.Set(10.0, 10.1)
	fmt.Println(len(sellBoardBuffer))
	sellBoardBuffer.Set(10.1, 10.1)
	fmt.Println(len(sellBoardBuffer))

	sellBoardBuffer.Set(10.1, 0)
	fmt.Println(len(sellBoardBuffer))
}

func TestBoard_GetHighPrice(t *testing.T) {
	var board Board
	board.Reset()

	board.Set(10.0, 10.0)
	board.Set(12.0, 10.0)
	board.Set(9.0, 10.0)

	high := board.GetHighPrice()
	log.Print("Price=", high)
}

func TestBoard_GetLowPrice(t *testing.T) {
	var board Board
	board.Reset()

	board.Set(10.0, 10.0)
	board.Set(12.0, 10.0)
	board.Set(9.0, 10.0)

	low := board.GetLowPrice()
	log.Print("Price=", low)

	board.Set(9.0, 0)
	low = board.GetLowPrice()
	log.Print("Price=", low)

}
