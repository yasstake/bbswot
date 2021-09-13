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
	BuyPrice   float64
	BuyEdge    float64
	SellPrice  float64
	SellEdge   float64
	EdgeTime   int64

	buyQ  []ExecPrice
	sellQ []ExecPrice
}

func (c *ExecQueue) Init(durationE6 int64) {
	c.buyQ = make([]ExecPrice, 0)
	c.sellQ = make([]ExecPrice, 0)
	c.durationE6 = durationE6
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

func MaxPrice(queue []ExecPrice) (price float64) {
	for _, item := range queue {
		if price < item.price || price == 0 {
			price = item.price
		}
	}

	return price
}

func MinPrice(queue []ExecPrice) (price float64) {
	for _, item := range queue {
		if item.price < price || price == 0 {
			price = item.price
		}
	}
	return price
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
	// Enqueue action
	if action == common.TRADE_BUY {
		c.buyQ = EnqueueAction(c.buyQ, timeE6, price, size)
	} else if action == common.TRADE_SELL {
		c.sellQ = EnqueueAction(c.sellQ, timeE6, price, size)
	} else {
		log.Error("Unknown action", action)
	}

	// Dequeue old queue
	c.EdgeTime = timeE6 - c.durationE6

	var buyDequeue []ExecPrice

	// delete buy q
	buyDequeue, c.buyQ = DequeAction(c.buyQ, c.EdgeTime)
	l := len(buyDequeue)
	if 0 < l {
		//fmt.Println("Dequeue Buy", buyDequeue)
		c.BuyEdge = buyDequeue[l-1].price
	}

	if action == common.TRADE_BUY {
		l = len(c.buyQ)
		if 0 < l {
			if c.BuyPrice < price || c.BuyPrice == 0 {
				c.BuyPrice = price
			} else {
				c.BuyPrice = MaxPrice(c.buyQ)
			}

		} else {
			c.BuyPrice = 0
		}
	}

	// delete sell q
	var sellDequeue []ExecPrice

	sellDequeue, c.sellQ = DequeAction(c.sellQ, c.EdgeTime)
	l = len(sellDequeue)
	if 0 < l {
		// fmt.Println("Dequeue Sell", buyDequeue)
		c.SellEdge = sellDequeue[l-1].price
	}

	if action == common.TRADE_SELL {
		l = len(c.sellQ)
		if 0 < l {
			if price < c.SellPrice || c.SellPrice == 0 {
				c.SellPrice = price
			} else {
				c.SellPrice = MinPrice(c.sellQ)
			}
		} else {
			c.SellPrice = 0
		}
	}

	if len(c.buyQ) == 0 {
		c.BuyPrice = 0
	}
	if len(c.sellQ) == 0 {
		c.SellPrice = 0
	}

	return c.EdgeTime, c.BuyPrice, c.SellPrice
}
