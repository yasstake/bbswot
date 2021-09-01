package bb

type Board map[float64]float64
type BoardItem struct {
	price  float64
	volume float64
}

var buyBoardBuffer Board
var sellBoardBuffer Board

func (c *Board) Reset() {
	*c = make(map[float64]float64)
}

func (c *Board) Add(price float64, volume float64) {
	(*c)[price] = volume
}
