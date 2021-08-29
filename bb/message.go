package bb

import (
	"encoding/json"
	"log"
)

type Message struct {
	Topic       string          `json:"topic"`
	Type        string          `json:"type"`
	Data        json.RawMessage `json:"data"`
	Sequence    int64           `json:"cross_seq"`
	TimeStampE6 int64           `json:"timestamp_e6"`
}

func ParseMessage(m []byte) (message Message) {
	err := json.Unmarshal(m, &message)
	if err != nil {
		log.Fatalln("Fail to pase message", err, message)
	}

	return message
}

func GetMessageJson(m []byte) json.RawMessage {
	message := ParseMessage(m)
	return message.Data
}
