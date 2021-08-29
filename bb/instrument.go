package bb

import (
	"bbswot/common"
	"encoding/json"
	"fmt"
	"log"
)

type InstrumentDelta struct {
	Delete []Instrument `json:"delete"`
	Update []Instrument `json:"update"`
	Insert []Instrument `json:"insert"`
}

type Instrument struct {
	Id                int    `json:"id"`
	Symbol            string `json:"symbol"`
	LastPrice         int    `json:"last_price_e4"`       //:536925000,
	BitPrice          int    `json:"bid1_price_e4"`       //:536925000,
	AskPrice          int    `json:"ask1_price_e4"`       //:536930000,
	LastTickDirection string `json:"last_tick_direction"` //:"ZeroMinusTick",
	PrevPrice         int    `json:"prev_price_24h_e4"`   //:503145000,
	// `json:"price_24h_pcnt_e6"`//:67137,
	HighPrice int `json:"high_price_24h_e4"` //:539840000,
	LowPrice  int `json:"low_price_24h_e4"`  //:470000000,
	// `json:"prev_price_1h_e4"`//:537670000,
	//`json:"price_1h_pcnt_e6"`//:-1385,
	MarkPrice            int `json:"mark_price_e4"`             //:536850500,
	IndexPrice           int `json:"index_price_e4"`            //:536796200,
	OpenInterest         int `json:"open_interest"`             //:1905193274,
	OpenValue            int `json:"open_value_e8"`             //:1461680310351,
	TotalTurnOver        int `json:"total_turnover_e8"`         //:7597234355112527,
	TurnOver24h          int `json:"turnover_24h_e8"`           //:17278982841461,
	TotalVolume          int `json:"total_volume"`              //:1465809840880,
	Volume24h            int `json:"volume_24h"`                //:8787821958,
	FundingRate          int `json:"funding_rate_e6"`           //:-51,
	PredictedFundingRate int `json:"predicted_funding_rate_e6"` //:9,
	//`json:"cross_seq"`//:6183139859,
	//`json:"created_at"`//:"2018-11-14T16:33:26Z",
	//`json:"updated_at"`//:"2021-04-26T15:05:12Z",
	NextFundingTime string `json:"next_funding_time"` //:"2021-04-26T16:00:00Z",
	//`json:"countdown_hour"`//:1},
	TimeStampE6 int64 `json:"timestamp_e6"`
}

func (c *Instrument) update(d Instrument) {
	if d.Id != 0 {
		c.Id = d.Id
	}
	if d.Symbol != "" {
		c.Symbol = d.Symbol
	}
	if d.LastPrice != 0 {
		c.LastPrice = d.LastPrice
	}
	if d.BitPrice != 0 {
		c.BitPrice = d.BitPrice
	}
	if d.AskPrice != 0 {
		c.AskPrice = d.AskPrice
	}
	if d.LastTickDirection != "" {
		c.LastTickDirection = d.LastTickDirection
	}
	if d.PrevPrice != 0 {
		c.PrevPrice = d.PrevPrice
	}
	if d.HighPrice != 0 {
		c.HighPrice = d.HighPrice
	}
	if d.LowPrice != 0 {
		c.LowPrice = d.LowPrice
	}
	if d.MarkPrice != 0 {
		c.MarkPrice = d.MarkPrice
	}
	if d.IndexPrice != 0 {
		c.IndexPrice = d.IndexPrice
	}
	if d.OpenInterest != 0 {
		c.OpenInterest = d.OpenInterest
	}
	if d.OpenValue != 0 {
		c.OpenValue = d.OpenValue
	}
	if d.TotalTurnOver != 0 {
		c.TotalTurnOver = d.TotalTurnOver
	}
	if d.TurnOver24h != 0 {
		c.TurnOver24h = d.TurnOver24h
	}
	if d.TotalVolume != 0 {
		c.TotalVolume = d.TotalVolume
	}
	if d.Volume24h != 0 {
		c.Volume24h = d.Volume24h
	}
	if d.FundingRate != 0 {
		c.FundingRate = d.FundingRate
	}
	if d.PredictedFundingRate != 0 {
		c.PredictedFundingRate = d.PredictedFundingRate
	}
	if d.NextFundingTime != "" {
		c.NextFundingTime = d.NextFundingTime
	}

	if d.TimeStampE6 != 0 {
		c.TimeStampE6 = d.TimeStampE6
	}
}

func (c *Instrument) ToLog() (result string) {
	result = ""

	t := c.TimeStampE6

	// Open Interest
	if c.OpenInterest != 0 {
		result += MakeWsLogRec(common.OPEN_INTEREST, t, 0, float64(c.OpenInterest), "")
	}

	// Open Value
	if c.OpenValue != 0 {
		result += MakeWsLogRec(common.OPEN_VALUE, t, 0, float64(c.OpenValue), "")
	}

	// TurnOver
	if c.TotalTurnOver != 0 {
		result += MakeWsLogRec(common.TURN_OVER, t, 0, float64(c.TotalTurnOver), "")
	}

	if c.FundingRate != 0 {
		result += MakeWsLogRec(common.FUNDING_RATE, t, 0, float64(c.FundingRate), "")
	}

	if c.PredictedFundingRate != 0 {
		if c.NextFundingTime != "" {
			timeMs := common.ParseIsoTimeToE6(c.NextFundingTime)
			sec, msec := common.DivideSecAndMs(timeMs)
			nextFundingTime := fmt.Sprintf("%d.%d", sec, msec)
			result += MakeWsLogRec(common.PREDICTED_FUNDING_RATE, t, 0, float64(c.PredictedFundingRate), nextFundingTime)
		}
	}

	return result
}

func ParseInstrumentSnapshot(message json.RawMessage, timeE6 int64) (result string) {
	var InstrumentData Instrument

	err := json.Unmarshal(message, &InstrumentData)
	if err != nil {
		log.Fatalln("Fail to pase message", err, message)
	}

	InstrumentData.TimeStampE6 = timeE6
	result += InstrumentData.ToLog()

	return result
}

func ParseInstrumentDelta(message json.RawMessage, timeE6 int64) (result string) {
	var data InstrumentDelta

	err := json.Unmarshal(message, &data)
	if err != nil {
		log.Fatalln("Fail to pase message", err, message)
	}

	result = ""

	for i := range data.Update {
		data.Update[i].TimeStampE6 = timeE6
		result += data.Update[i].ToLog()
	}

	// Assume Delete and Insert message is not implemented
	for i := range data.Delete {
		data.Delete[i].TimeStampE6 = timeE6
		log.Println("INFO delete ", data.Delete[i])

		result += data.Delete[i].ToLog()
	}

	for i := range data.Insert {
		log.Println("INFO Insert", data.Insert[i])
		data.Insert[i].TimeStampE6 = timeE6
		result += data.Insert[i].ToLog()
	}

	return result
}
