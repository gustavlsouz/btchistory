// file: list_posts.go
package main

import (
	// import standard libraries

	"fmt"
	"log"
	"strconv"
	"strings"

	// import third party libraries
	"github.com/PuerkitoBio/goquery"
)

type Cotacao struct {
	valorReais float64
}

func NewCotacao() *Cotacao {
	return &Cotacao{}
}

func postScrape(linkAlvo string, conteudoParaProcurar string, attribute string) *Cotacao {
	conteudoHTML, erro := goquery.NewDocument(linkAlvo)
	CheckErr(erro)

	resultadoDaBusca := conteudoHTML.Find(conteudoParaProcurar).Map(func(index int, item *goquery.Selection) string {
		bitValue, _ := item.Attr(attribute)
		return bitValue
	})
	fmt.Println(len(resultadoDaBusca))
	cotacao := NewCotacao()
	cotacao.valorReais, _ = strconv.ParseFloat(strings.Replace(resultadoDaBusca[0], ",", ".", -1), 64)
	return cotacao
}

func main() {
	cotacao := postScrape("http://dolarhoje.com/bitcoin-hoje/", "#nacional", "value")
	fmt.Printf("1BTC : %5.2fReais", cotacao.valorReais)
}

func CheckErr(erro error) {
	if erro != nil {
		log.Fatal(erro)
	}
}
