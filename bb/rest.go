package bb

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type RestResponse struct {
	Code    json.Number     `json:"ret_code"`   // "ret_code":0
	Message string          `json:"ret_msg"`    // "ret_msg":"OK",
	ExtCode string          `json:"ext_code"`  // "ext_code":"",
    ExtInfo string          `json:"ext_info"`  // "ext_info":""
	Result  json.RawMessage `json:"result"`    // "result": {JSON}
	TimeStampSec    json.Number     `json:"time_now"`  // "time_now":"1630186820.789949"
}


func RestRequest(url string) (result string, timeStampMs int64, err error) {
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		log.Println(err)
		return result, 0, err
	}

	client := new(http.Client)
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return result, 0, err
	}

	defer resp.Body.Close()
	byteArray, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Println(err)
		return result, 0, err
	}

	var response RestResponse
	err = json.Unmarshal(byteArray, &response)
	if err != nil {
		log.Println(err)
		return result, 0, err
	}

	code, err := response.Code.Int64()

	if err != nil || code != 0 {
		return result, 0, err
	}

	t, err := response.TimeStampSec.Float64()
	if err != nil {
		return result, 0, err
	}
	timeStampMs = int64(t * 1_000)

	return string(response.Result), timeStampMs, err
}
