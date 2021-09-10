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

	if action == common.TRADE_BUY {
		_, c.buyQ = append(c.buyQ)
		q := sortQ(c.buyQ)
		edgeTimeE6 = c.buyQ[0].timeE6
		buyEdge = q[0].price
	} else if action == common.TRADE_SELL {
		_, c.sellQ = append(c.sellQ)
		q := sortQ(c.sellQ)
		pos := len(q) - 1
		edgeTimeE6 = c.sellQ[pos].timeE6
		sellEdge = q[pos].price
	} else {
		log.Error("Unknown action ", action)
	}

	return edgeTimeE6, buyEdge, sellEdge
}
