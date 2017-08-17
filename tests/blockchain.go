package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type JSONBlockchain struct {
	USD map[string]interface{} `json:"USD"`
	AUD map[string]interface{} `json:"AUD"`
	BRL map[string]interface{} `json:"BRL"`
	CAD map[string]interface{} `json:"CAD"`
	CHF map[string]interface{} `json:"CHF"`
	CLP map[string]interface{} `json:"CLP"`
	CNY map[string]interface{} `json:"CNY"`
	DKK map[string]interface{} `json:"DKK"`
	EUR map[string]interface{} `json:"EUR"`
	GBP map[string]interface{} `json:"GBP"`
	HKD map[string]interface{} `json:"HKD"`
	INR map[string]interface{} `json:"INR"`
	ISK map[string]interface{} `json:"ISK"`
	JPY map[string]interface{} `json:"JPY"`
	KRW map[string]interface{} `json:"KRW"`
	NZD map[string]interface{} `json:"NZD"`
	PLN map[string]interface{} `json:"PLN"`
	RUB map[string]interface{} `json:"RUB"`
	SEK map[string]interface{} `json:"SEK"`
	SGD map[string]interface{} `json:"SGD"`
	THB map[string]interface{} `json:"THB"`
	TWD map[string]interface{} `json:"TWD"`
}

var link = "https://blockchain.info/pt/ticker"

func getStations(body []byte) (*JSONBlockchain, error) {
	var s = new(JSONBlockchain)
	err := json.Unmarshal(body, &s)
	if err != nil {
		fmt.Println("whoops:", err)
	}
	return s, err
}

func main() {

	res, err := http.Get(link)
	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}

	s, err := getStations([]byte(body))
	fmt.Println(s.BRL["last"])
}
