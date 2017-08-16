// file: list_posts.go
package main

import (
	"fmt"

	"../../gustavlsouz/btchistory/scraper"
)

func main() {
	cotacao := scraper.PostScrape("http://dolarhoje.com/bitcoin-hoje/", "#nacional", "value")
	fmt.Printf("1BTC : R$%5.2f reais", cotacao.ValorReais)
}
