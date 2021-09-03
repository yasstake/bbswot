package bb

import (
	"sort"
)

type Board struct {
	data map[float64]float64
	high float64
	low  float64
}

var buyBoardBuffer Board
var sellBoardBuffer Board

func (c *Board) Reset() {
	(*c).data = make(map[float64]float64)
}

func (c *Board) Data() map[float64]float64 {
	return (*c).data
}

// Set
// if volume is 0, the board price will be erased.
func (c *Board) Set(price float64, volume float64) {
	if c.high == 0 {
		c.high = price
	}
	if c.low == 0 {
		c.low = price
	}

	if volume == 0 {
		delete(c.data, price)
		c.low, c.high = c.CalcEdgePrice()
	} else {
		c.data[price] = volume

		if c.high < price {
			c.high = price
		}
		if price < c.low {
			c.low = price
		}
	}
}

func (c *Board) Get(price float64) float64 {
	if (*c).data == nil {
		return 0
	}

	return (*c).data[price]
}

func (c *Board) Len() int {
	if (*c).data == nil {
		return 0
	}

	return len((*c).data)
}

func (c *Board) GetHigh() float64 {
	return c.high
}

func (c *Board) GetLow() float64 {
	return c.low
}

func (c *Board) CalcHighPrice() float64 {
	_, high := c.CalcEdgePrice()

	return high
}

func (c *Board) CalcLowPrice() float64 {
	low, _ := c.CalcEdgePrice()
	return low
}

// CalcEdgePrice
// calculate Low and High edge price of board
// if board have no value, returns 0
// GetLow and GetHigh is available for retrieve cached result.
func (c *Board) CalcEdgePrice() (low, high float64) {
	l := len((*c).data)

	if l == 0 {
		return 0, 0
	}

	prices := make([]float64, 0, l)

	for key, _ := range (*c).data {
		prices = append(prices, key)
	}

	sort.Float64s(prices)

	low = prices[0]
	high = prices[l-1]

	return low, high
}
