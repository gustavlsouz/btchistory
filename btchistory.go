package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"runtime"
	"sync"
	"time"

	"../../gustavlsouz/btchistory/conf"
	"../../gustavlsouz/btchistory/scraper"
	"../../gustavlsouz/btchistory/utils"
)

type JSONBlockchain struct {
	USD map[string]interface{} `json:"USD"`
	BRL map[string]interface{} `json:"BRL"`
	EUR map[string]interface{} `json:"EUR"`
}

func main() {
	// load confs
	configuration, _ := conf.LoadConf()

	log.Printf("\nDB conf:\nuser:%s\npassword:%s\nFrequency:%s\nCores:%d\n\n",
		configuration.DB["user"], configuration.DB["passwd"],
		configuration.Freq, configuration.Core)

	// number core
	runtime.GOMAXPROCS(configuration.Core)

	// scrap
	var controle sync.WaitGroup

	var timeToWait time.Duration
	timeToWait, _ = time.ParseDuration(configuration.Freq)

	controle.Add(2)
	go taskDolarHoje(timeToWait, &controle)
	go taskBlockchain(timeToWait, &controle)
	controle.Wait()
}

func taskDolarHoje(timeToWait time.Duration, controle *sync.WaitGroup) {
	defer controle.Done()

	dolarHojeLink := "http://dolarhoje.com/bitcoin-hoje/"

	// scraping
	for {
		cotacao := scraper.PostScrape(dolarHojeLink, "#nacional", "value")
		log.Printf("\n\t1BTC : R$%5.2f reais\n\t[%s]", cotacao.ValorReais, dolarHojeLink)
		time.Sleep(timeToWait)
	}
}

func taskBlockchain(timeToWait time.Duration, controle *sync.WaitGroup) {
	defer controle.Done()

	var link = "https://blockchain.info/pt/ticker"

	for {
		resp, err := http.Get(link)
		utils.ShowErr(err)
		body, err := ioutil.ReadAll(resp.Body)
		utils.ShowErr(err)
		json, err := getJSONBlockchain([]byte(body))
		log.Printf("\n\t1BTC : %s %5.2f\n\t[%s]\n", json.BRL["symbol"], json.BRL["last"], link)
		time.Sleep(timeToWait)
	}
}

func getJSONBlockchain(body []byte) (*JSONBlockchain, error) {
	var JSONModel = new(JSONBlockchain)
	err := json.Unmarshal(body, &JSONModel)
	if err != nil {
		log.Println(err, "Erro na leitura de JSON")
	}
	return JSONModel, err
}
