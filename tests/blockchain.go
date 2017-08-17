package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

/*
{
	"USD" : {"15m" : 4467.64, "last" : 4467.64, "buy" : 4468.12, "sell" : 4467.15, "symbol" : "$"},
	"BRL" : {"15m" : 14114.51, "last" : 14114.51, "buy" : 14116.04, "sell" : 14112.97, "symbol" : "R$"},
	"EUR" : {"15m" : 3844.74, "last" : 3844.74, "buy" : 3851.86, "sell" : 3837.61, "symbol" : "€"},
}
*/
type JSONBlockchain struct {
	USD map[string]interface{} `json:"USD"`
	BRL map[string]interface{} `json:"BRL"`
	EUR map[string]interface{} `json:"EUR"`
	/*AUD map[string]interface{} `json:"AUD"`
	CAD map[string]interface{} `json:"CAD"`
	CHF map[string]interface{} `json:"CHF"`
	CLP map[string]interface{} `json:"CLP"`
	CNY map[string]interface{} `json:"CNY"`
	DKK map[string]interface{} `json:"DKK"`
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
	TWD map[string]interface{} `json:"TWD"`*/
}

var link = "https://blockchain.info/pt/ticker"

func getJSONBlockchain(body []byte) (*JSONBlockchain, error) {
	var JSONModel = new(JSONBlockchain)
	err := json.Unmarshal(body, &JSONModel)
	if err != nil {
		log.Println(err, "Erro na leitura de JSON")
	}
	return JSONModel, err
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

	json, err := getJSONBlockchain([]byte(body))
	fmt.Println(json.BRL["last"])
}
