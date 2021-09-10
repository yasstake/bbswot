package bb

import (
	"bbswot/common"
	"github.com/labstack/gommon/log"
	"sort"
)

type ExecPrice struct {
	timeE6 int64
	price  float64
	size   float64
}

type ExecQueue struct {
	durationE6 int64
	buyPrice   float64
	sellPrice  float64
	buyQ       []ExecPrice
	sellQ      []ExecPrice
}

func (c *ExecQueue) Init() {
	c.buyQ = make([]ExecPrice, 0)
	c.sellQ = make([]ExecPrice, 0)
}

func (c *ExecQueue) Stat() (buyList []ExecPrice, sellList []ExecPrice) {

	sortList := func(execList []ExecPrice) (result []ExecPrice) {
		sorted := make([]ExecPrice, len(execList))
		copy(sorted, execList)

		sort.Slice(sorted, func(i, j int) bool {
			return sorted[i].price < sorted[j].price
		})

		var lastIndex int
		var lastPrice float64

		for _, item := range sorted {
			if lastPrice != item.price {
				result = append(result, item)
				lastIndex = len(result) - 1
				lastPrice = item.price
			} else {
				result[lastIndex].size += item.size
			}
		}
		return result
	}

	buyList = sortList(c.buyQ)
	sellList = sortList(c.sellQ)

	return buyList, sellList
}

func (c *ExecQueue) Action(action int, timeE6 int64, price float64, size float64) (edgeTimeE6 int64, buyEdge float64, sellEdge float64) {

	append := func(q []ExecPrice) (deque []ExecPrice, appendQ []ExecPrice) {

		exec := ExecPrice{timeE6: timeE6, price: price, size: size}
		q = append(q, exec)
		if len(q) == 1 {
			return nil, q
		}

		deque = make([]ExecPrice, 0)

		for {
			timeDiff := timeE6 - q[0].timeE6
			if c.durationE6 < timeDiff {
				deque = append(deque, q[0])
				q = q[1:]
			} else {
				break
			}
		}

		return deque, q
	}

	sortQ := func(q []ExecPrice) (sorted []ExecPrice) {
		sorted = make([]ExecPrice, len(q))
		copy(sorted, q)

		sort.Slice(sorted, func(i, j int) bool {
			return sorted[i].price < sorted[j].price
		})

		return sorted
	}

	var deq []ExecPrice
	if action == common.TRADE_BUY {
		// Select low edge price in the buy Queue
		deq, c.buyQ = append(c.buyQ)
		l := len(deq)

		if price < c.buyPrice || c.buyPrice == 0 {
			c.buyPrice = price
		} else if l != 0 {
			for _, item := range deq {
				if item.price <= c.buyPrice {
					q := sortQ(c.buyQ)
					c.buyPrice = q[0].price
					break
				}
			}
		}
		edgeTimeE6 = c.buyQ[0].timeE6
	} else if action == common.TRADE_SELL {
		// Select high edge price in the sell Queue
		deq, c.sellQ = append(c.sellQ)
		l := len(deq)
		qlen := len(c.sellQ) - 1

		if c.sellPrice < price {
			c.sellPrice = price
		} else if l != 0 {
			for _, item := range deq {
				if c.sellPrice <= item.price {
					q := sortQ(c.sellQ)
					c.sellPrice = q[qlen].price
					break
				}
			}
		}
		edgeTimeE6 = c.sellQ[qlen].timeE6
	} else {
		log.Error("Unknown action ", action)
	}

	return edgeTimeE6, c.buyPrice, c.sellPrice
}
