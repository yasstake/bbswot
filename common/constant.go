package common

const PARTIAL = 1

const UPDATE_BUY = 2

const UPDATE_SELL = 3
const TRADE_SELL_STR = "Sell"

const TRADE_BUY = 4
const TRADE_BUY_STR = "Buy"

const TRADE_BUY_LIQUID = 5
const TRADE_BUY_LIQUID_STR = "SELL"

const TRADE_SELL = 6
const TRADE_SELL_LIQUID = 7
const TRADE_SELL_LIQUID_STR = "BUY"

// TRADE_BUY_PRICE
// action, time, BUY_PRICE, 0, 0
const TRADE_BUY_PRICE = 8

// TRADE_SELL_PRICE
// action, time, SELL_PRICE, 0, 0
const TRADE_SELL_PRICE = 9

// OPEN_INTEREST
// action, time, 0,, volume,
const OPEN_INTEREST = 10

// OPEN_VALUE
// action, time, 0, volume
const OPEN_VALUE = 11

// TRUN_OVER
// action, time, 0, volume
const TURN_OVER = 12

// FUNDING_RATE
// action, time, 0, volume, next time
const FUNDING_RATE = 20

// PREDICTED_FUNDING_RATE
// action, time, 0, volume, next time
const PREDICTED_FUNDING_RATE = 21

var ACTION_STRING map[int]string

func init() {
	ACTION_STRING = map[int]string{
		PARTIAL:                "PARTIAL",
		UPDATE_BUY:             "UPD_BUY",
		UPDATE_SELL:            "UPD_SEL",
		TRADE_BUY:              "TR__BUY",
		TRADE_BUY_LIQUID:       "TR_BUYL",
		TRADE_SELL:             "TR__SEL",
		TRADE_SELL_LIQUID:      "TR_SELL",
		TRADE_BUY_PRICE:        "TR_BUYP",
		TRADE_SELL_PRICE:       "TR_SELP",
		OPEN_INTEREST:          "OP_INTT",
		OPEN_VALUE:             "OP_VALU",
		TURN_OVER:              "TU_OVER",
		FUNDING_RATE:           "FU_RATE",
		PREDICTED_FUNDING_RATE: "PR_FD_R",
	}
}
