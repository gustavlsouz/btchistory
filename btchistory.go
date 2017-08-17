package main

import (
	"log"
	"runtime"
	"sync"
	"time"

	"../../gustavlsouz/btchistory/conf"
	"../../gustavlsouz/btchistory/scraper"
)

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

	controle.Add(1)
	go taskBTCCheck(configuration.Freq, &controle)

	controle.Wait()
}

func taskBTCCheck(value string, controle *sync.WaitGroup) {
	defer controle.Done()

	dolarHojeLink := "http://dolarhoje.com/bitcoin-hoje/"
	var timeToWait time.Duration
	timeToWait, _ = time.ParseDuration(value)

	// scraping
	for true {
		cotacao := scraper.PostScrape(dolarHojeLink, "#nacional", "value")
		log.Printf("\n\t1BTC : R$%5.2f reais\n\t[%s]", cotacao.ValorReais, dolarHojeLink)
		time.Sleep(timeToWait)
	}
}
