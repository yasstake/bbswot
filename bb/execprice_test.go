package bb

import (
	"bbswot/common"
	"fmt"
	"log"
	"testing"
)

var execDATA = []ExecPrice{
	{1630439895807000, 47357, 832},
	{1630439895808000, 47357, 832},
	{1630439895808000, 47359.5, 2},
	{1630439895808000, 47354.5, 2},
	{1630439895808000, 47354.5, 954},
	{1630439895808000, 47354.5, 1},
	{1630439894032000, 47354, 2},
	{1630439894032000, 47354, 1},
	{1630439893874000, 47353.5, 1},
	{1630439893838000, 47353.5, 1},
	{1630439893831000, 47353.5, 10},
	{1630439893794000, 47353.5, 1},
	{1630439893787000, 47353.5, 1},
	{1630439893773000, 47353.5, 1},
	{1630439893772000, 47353.5, 11},
	{1630439893771000, 47351.5, 4},
	{1630439893769000, 47349.5, 1},
	{1630439893764000, 47351, 1875},
	{1630439893764000, 47351, 2},
	{1630439893764000, 47351, 250},
	{1630439893764000, 47350.5, 2711},
	{1630439893764000, 47350.5, 2},
	{1630439893764000, 47350.5, 100},
	{1630439893764000, 47350, 2},
	{1630439893764000, 47350, 3000},
	{1630439893764000, 47350, 1},
	{1630439896764000, 47350, 1},
}

func TestEnqueueAction(t *testing.T) {
	var queue []ExecPrice

	queue = EnqueueAction(queue, 10, 10.0, 10.0)
	queue = EnqueueAction(queue, 11, 10.0, 10.0)
	queue = EnqueueAction(queue, 12, 10.0, 10.0)

	log.Println(queue)
}

func TestDequeAction(t *testing.T) {
	var queue []ExecPrice

	queue = EnqueueAction(queue, 10, 10.0, 10.0)
	queue = EnqueueAction(queue, 11, 10.0, 10.0)
	queue = EnqueueAction(queue, 12, 10.0, 10.0)

	deque, queue := DequeAction(queue, 11)

	fmt.Println("Deque", deque)
	fmt.Println("Queue", queue)
}

func TestIsExistOutside(t *testing.T) {
	var queue []ExecPrice

	queue = EnqueueAction(queue, 10, 10.0, 10.0)
	queue = EnqueueAction(queue, 11, 11.0, 10.0)
	queue = EnqueueAction(queue, 12, 12.0, 10.0)

	if !CompareExecPrice(9, queue, true) {
		t.Error("")
	}
	if !CompareExecPrice(10, queue, true) {
		t.Error("")
	}
	if !CompareExecPrice(11, queue, true) {
		t.Error("")
	}
	if CompareExecPrice(12, queue, true) {
		t.Error("")
	}
	if CompareExecPrice(13, queue, true) {
		t.Error("")
	}

	if CompareExecPrice(9, queue, false) {
		t.Error("")
	}
	if CompareExecPrice(10, queue, false) {
		t.Error("")
	}
	if !CompareExecPrice(11, queue, false) {
		t.Error("")
	}
	if !CompareExecPrice(12, queue, false) {
		t.Error("")
	}
	if !CompareExecPrice(13, queue, false) {
		t.Error("")
	}
}

func TestExecQueue_Action_AddBuy(t *testing.T) {
	var q ExecQueue

	q.Init()
	q.durationE6 = 2

	edge, buy, sell := q.Action(common.TRADE_BUY, 0, 0.5, 1)
	log.Print(edge, buy, sell)

	edge, buy, sell = q.Action(common.TRADE_BUY, 1, 1.5, 1)
	log.Print(edge, buy, sell)

	edge, buy, sell = q.Action(common.TRADE_BUY, 2, 2.5, 1)
	log.Print(edge, buy, sell)

	edge, buy, sell = q.Action(common.TRADE_BUY, 3, 1, 1)
	log.Print(edge, buy, sell)

	edge, buy, sell = q.Action(common.TRADE_BUY, 4, 10, 1)
	log.Print(edge, buy, sell)

	edge, buy, sell = q.Action(common.TRADE_BUY, 5, 2, 1)
	log.Print(edge, buy, sell)

	edge, buy, sell = q.Action(common.TRADE_BUY, 6, 3, 1)
	log.Print(edge, buy, sell)

	edge, buy, sell = q.Action(common.TRADE_BUY, 7, 1, 1)
	log.Print(edge, buy, sell)
}

func TestExecQueue_Action_AddSell(t *testing.T) {
	var q ExecQueue

	q.Init()

	q.Init()
	q.durationE6 = 2

	edge, buy, sell := q.Action(common.TRADE_BUY, 0, 0.5, 1)
	log.Print(edge, buy, sell)

	edge, buy, sell = q.Action(common.TRADE_SELL, 0, 0.5, 1)
	log.Print(edge, buy, sell)

	edge, buy, sell = q.Action(common.TRADE_SELL, 1, 1.5, 1)
	log.Print(edge, buy, sell)
	edge, buy, sell = q.Action(common.TRADE_SELL, 1, 1.5, 1)
	log.Print(edge, buy, sell)
	edge, buy, sell = q.Action(common.TRADE_SELL, 1, 1.5, 1)
	log.Print(edge, buy, sell)
	edge, buy, sell = q.Action(common.TRADE_SELL, 1, 1.5, 1)
	log.Print(edge, buy, sell)

	edge, buy, sell = q.Action(common.TRADE_SELL, 2, 2.5, 1)
	log.Print(edge, buy, sell)

	edge, buy, sell = q.Action(common.TRADE_SELL, 3, 1, 1)
	log.Print(edge, buy, sell)

	edge, buy, sell = q.Action(common.TRADE_SELL, 4, 10, 1)
	log.Print(edge, buy, sell)

	edge, buy, sell = q.Action(common.TRADE_SELL, 5, 2, 1)
	log.Print(edge, buy, sell)

	edge, buy, sell = q.Action(common.TRADE_SELL, 6, 3, 1)
	log.Print(edge, buy, sell)

	edge, buy, sell = q.Action(common.TRADE_SELL, 7, 5.5, 1)
	log.Print(edge, buy, sell)

	edge, buy, sell = q.Action(common.TRADE_SELL, 8, 5.5, 1)
	log.Print(edge, buy, sell)

	edge, buy, sell = q.Action(common.TRADE_SELL, 20, 3, 1)
	log.Print(edge, buy, sell)
}

func TestLoadExecData2(t *testing.T) {
	var q ExecQueue

	q.Init()
	q.durationE6 = 10

	for _, item := range execDATA {
		edge, buy, sell := q.Action(common.TRADE_SELL, item.timeE6, item.price, item.size)
		log.Print(edge, buy, sell)
	}
}

func TestMakeUnique(t *testing.T) {
	file := "../TEST_DATA/BTCUSD2021-08-31.sort.csv.gz"
	stream := common.OpenFileReader(file)

	stream.Scan() // skip header line

	var lastBuyPrice float64
	var lastSellPrice float64

	for stream.Scan() {
		rAction, rTimeE6, rPrice, rVolume, _ := ParseArchivedLogRec(stream.Text())

		if rAction == common.TRADE_BUY {
			if lastBuyPrice != rPrice {
				fmt.Println(common.TimeE6ToString(rTimeE6), rAction, rTimeE6, rPrice, rVolume)

				lastBuyPrice = rPrice
			}

		} else if rAction == common.TRADE_SELL {
			if lastSellPrice != rPrice {
				fmt.Println(common.TimeE6ToString(rTimeE6), rAction, rTimeE6, rPrice, rVolume)

				lastSellPrice = rPrice
			}
		}

	}
}

func TestLoadExec(t *testing.T) {
	var q ExecQueue

	q.Init()
	//q.durationE6 = 1_000_000 * 60
	q.durationE6 = 2

	file := "../TEST_DATA/BTCUSD2021-08-31.sort.csv.gz"
	stream := common.OpenFileReader(file)

	stream.Scan() // skip header line

	var recordNumber int64
	var lastBuyPrice float64
	var lastSellPrice float64

	fmt.Println("load start")

	for stream.Scan() {
		rAction, rTimeE6, rPrice, rVolume, _ := ParseArchivedLogRec(stream.Text())
		// fmt.Println(common.TimeE6ToString(rTimeE6), rAction, rTimeE6, rPrice, rVolume)
		timeE6, buyPrice, sellPrice := q.Action(rAction, rTimeE6, rPrice, rVolume)

		if lastBuyPrice != buyPrice {
			lastBuyPrice = buyPrice
			fmt.Println("BUY", common.TimeE6ToString(timeE6), timeE6, buyPrice, q.buyEdge)
		}

		if lastSellPrice != sellPrice {
			lastSellPrice = sellPrice
			fmt.Println("SELL", common.TimeE6ToString(timeE6), timeE6, sellPrice, q.sellEdge)
		}

		recordNumber += 1
	}
}
