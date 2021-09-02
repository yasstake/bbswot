package bb

import (
	"sort"
)

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

func (c *Board) Set(price float64, volume float64) {
	if volume == 0 {
		delete(*c, price)
	} else {
		(*c)[price] = volume
	}
}

func (c *Board) GetHighPrice() float64 {
	return c.GetEdgePrice(true)
}

func (c *Board) GetLowPrice() float64 {
	return c.GetEdgePrice(false)
}

func (c *Board) GetEdgePrice(high bool) (result float64) {
	l := len(*c)
	prices := make([]float64, 0, l)

	for key, _ := range *c {
		prices = append(prices, key)
	}

	sort.Float64s(prices)

	if high {
		result = prices[l-1]
	} else {
		result = prices[0]
	}

	return result
}
