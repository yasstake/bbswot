package bb

import (
	"bbswot/common"
	"fmt"
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
	buyEdge    float64
	sellPrice  float64
	sellEdge   float64
	edgeTime   int64

	buyQ  []ExecPrice
	sellQ []ExecPrice
}

func (c *ExecQueue) Init() {
	c.buyQ = make([]ExecPrice, 0)
	c.sellQ = make([]ExecPrice, 0)
}

func (c *ExecQueue) Stat() (buyList []ExecPrice, buyVolume float64, sellList []ExecPrice, sellVolume float64) {

	// sort and group by price with sum price
	groupByPrice := func(execList []ExecPrice) (result []ExecPrice, volume float64) {
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
			volume += item.size
		}
		return result, volume
	}

	buyList, buyVolume = groupByPrice(c.buyQ)
	sellList, sellVolume = groupByPrice(c.sellQ)

	return buyList, buyVolume, sellList, sellVolume
}

// SplitQueue
//Split queue in two(before, after) by time
func SplitQueue(queue []ExecPrice, timeE6 int64) (before []ExecPrice, after []ExecPrice) {
	l := len(queue)
	if l == 0 {
		return before, after
	}

	i := l - 1

	for {
		if queue[i].timeE6 <= timeE6 {
			break
		}

		i -= 1

		if i < 0 {
			break
		}
	}

	if i == -1 {
		before = []ExecPrice{}
		after = queue
	} else if i == l-1 {
		before = queue
		after = []ExecPrice{}
	} else {
		before = queue[:i+1]
		after = queue[i+1:]
	}

	return before, after
}

func EnqueueAction(queue []ExecPrice, timeE6 int64, price float64, size float64) (result []ExecPrice) {
	exec := ExecPrice{timeE6: timeE6, price: price, size: size}
	//result = append(queue, exec)
	before, after := SplitQueue(queue, timeE6)

	result = append(before, exec)
	result = append(result, after...)

	return result
}

func DequeAction(q []ExecPrice, timeE6 int64) (deque []ExecPrice, remainQ []ExecPrice) {

	if len(q) == 0 {
		return deque, q
	}

	deque = make([]ExecPrice, 0)

	// deque until the queue length within duration
	for {
		if len(q) == 0 || timeE6 < q[0].timeE6 {
			break
		}
		deque = append(deque, q[0])
		q = q[1:]
	}

	return deque, q
}

func CompareExecPrice(price float64, q []ExecPrice, higher bool) bool {
	for _, item := range q {
		if higher {
			if price <= item.price {
				return true
			}
		} else {
			if item.price <= price {
				return true
			}
		}
	}
	return false
}

func (c *ExecQueue) Action(action int, timeE6 int64, price float64, size float64) (edgeTimeE6 int64, buyEdge float64, sellEdge float64) {
	sortPrice := func(q []ExecPrice) (sorted []ExecPrice) {
		sorted = make([]ExecPrice, len(q))
		copy(sorted, q)

		sort.Slice(sorted, func(i, j int) bool {
			return sorted[i].price < sorted[j].price
		})

		return sorted
	}

	// Enqueue action
	if action == common.TRADE_BUY {
		c.buyQ = EnqueueAction(c.buyQ, timeE6, price, size)
	} else if action == common.TRADE_SELL {
		c.sellQ = EnqueueAction(c.sellQ, timeE6, price, size)
	} else {
		log.Error("Unknown action", action)
	}

	// Dequeue old queue
	c.edgeTime = timeE6 - c.durationE6

	var dequeue []ExecPrice

	// delete buy q
	dequeue, c.buyQ = DequeAction(c.buyQ, c.edgeTime)
	l := len(dequeue)
	if 0 < l {
		//fmt.Println("Dequeue Buy", dequeue)
		c.buyEdge = dequeue[l-1].price
	}

	if action == common.TRADE_BUY {
		l = len(c.buyQ)
		if 0 < l {
			if price < c.buyPrice || c.buyPrice == 0 {
				c.buyPrice = price
			} else if CompareExecPrice(c.buyPrice, dequeue, false) {
				fmt.Println("Update")
				c.buyPrice = sortPrice(c.buyQ)[0].price
			}
		} else {
			c.buyPrice = 0
		}
	}

	// delete sell q
	dequeue, c.sellQ = DequeAction(c.sellQ, c.edgeTime)
	l = len(dequeue)
	if 0 < l {
		// fmt.Println("Dequeue Sell", dequeue)
		c.sellEdge = dequeue[l-1].price
	}

	if action == common.TRADE_SELL {
		l = len(c.sellQ)
		if 0 < l {
			if c.sellPrice < price {
				c.sellPrice = price
			} else if CompareExecPrice(c.sellPrice, dequeue, true) {
				c.sellPrice = sortPrice(c.sellQ)[l-1].price
			}
		} else {
			c.sellPrice = 0
		}
	}

	if len(c.buyQ) == 0 {
		c.buyPrice = 0
	}
	if len(c.sellQ) == 0 {
		c.sellPrice = 0
	}

	return c.edgeTime, c.buyPrice, c.sellPrice
}
