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

	sellBoardBuffer.Add(10.0, 10.0)
	fmt.Println(len(sellBoardBuffer))
	sellBoardBuffer.Add(10.0, 10.1)
	fmt.Println(len(sellBoardBuffer))
	sellBoardBuffer.Add(10.1, 10.1)
	fmt.Println(len(sellBoardBuffer))
}
