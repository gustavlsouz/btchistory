package main

import (
	"encoding/json"
	"log"
	"os"
	"runtime"
	"sync"
	"time"

	"../../gustavlsouz/btchistory/scraper"
	"../../gustavlsouz/btchistory/utils"
)

type Configuration struct {
	DB   map[string]string
	Freq map[string]string
	Core int
}

var (
	confFile = "conf/conf.json"
)

func main() {
	// load confs
	file, _ := os.Open(confFile)
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)
	utils.CheckErr(err, "erro na realizacao de decode das configurações...")
	log.Printf("\nDB conf:\nuser:%s\npassword:%s\nFrequency:\n%s : %s\nCores:%d\n\n",
		configuration.DB["user"], configuration.DB["passwd"], configuration.Freq["time"],
		configuration.Freq["value"], configuration.Core)

	// number core
	runtime.GOMAXPROCS(configuration.Core)

	// scrap
	var controle sync.WaitGroup

	controle.Add(1)
	go taskBTCCheck(configuration.Freq["time"], configuration.Freq["value"], &controle)

	controle.Wait()
}

func taskBTCCheck(timeType string, value string, controle *sync.WaitGroup) {
	defer controle.Done()

	dolarHojeLink := "http://dolarhoje.com/bitcoin-hoje/"
	var timeToWait time.Duration

	if timeType == "sec" {
		timeToWait, _ = time.ParseDuration(value + "s")
	} else if timeType == "min" {
		timeToWait, _ = time.ParseDuration(value + "m")
	} else if timeType == "hour" {
		timeToWait, _ = time.ParseDuration(value + "h")
	}

	// scraping
	for true {
		cotacao := scraper.PostScrape(dolarHojeLink, "#nacional", "value")
		log.Printf("\n\t1BTC : R$%5.2f reais\n\t[%s]", cotacao.ValorReais, dolarHojeLink)
		time.Sleep(timeToWait)
	}
}
