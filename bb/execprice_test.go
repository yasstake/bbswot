package bb

import (
	"bbswot/common"
	"log"
	"testing"
)

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

	edge, buy, sell := q.Action(common.TRADE_SELL, 0, 0.5, 1)
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

	edge, buy, sell = q.Action(common.TRADE_SELL, 7, 1, 1)
	log.Print(edge, buy, sell)
}
