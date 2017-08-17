package main

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"runtime"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"../../gustavlsouz/btchistory/conf"
	"../../gustavlsouz/btchistory/scraper"
	"../../gustavlsouz/btchistory/structs"
	"../../gustavlsouz/btchistory/utils"
)

func main() {
	// load confs
	configuration, _ := conf.LoadConf()

	log.Printf("\nDB conf:\ndbname:%s\nuser:%s\npassword:%s\nFrequency:%s\nCores:%d\n\n",
		configuration.DB["dbname"],
		configuration.DB["user"], configuration.DB["passwd"],
		configuration.Freq, configuration.Core)

	// number core
	numCPU := runtime.NumCPU()
	if configuration.Core <= numCPU {
		runtime.GOMAXPROCS(configuration.Core)
	} else {
		runtime.GOMAXPROCS(numCPU)
	}

	var stringConnection = configuration.DB["user"] + ":" + configuration.DB["passwd"] + "@/" + configuration.DB["dbname"] + "?charset=utf8"
	db, err := sql.Open("mysql", stringConnection)
	utils.CheckErr(err)
	defer db.Close()

	// testing connection
	rows, err := db.Query("select nm_currency from btchist.currency")
	utils.CheckErr(err)

	for rows.Next() {
		var currency structs.Currency
		err := rows.Scan(&currency.Nmmode)
		utils.ShowErr(err)
		log.Println(currency.Nmmode)
	}

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

func getJSONBlockchain(body []byte) (*structs.JSONBlockchain, error) {
	var JSONModel = new(structs.JSONBlockchain)
	err := json.Unmarshal(body, &JSONModel)
	if err != nil {
		log.Println(err, "Erro na leitura de JSON")
	}
	return JSONModel, err
}
