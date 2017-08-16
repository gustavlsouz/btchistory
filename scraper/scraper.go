package scraper

import (
	// import standard libraries

	"strconv"
	"strings"

	// import third party libraries
	"../../btchistory/utils"
	"github.com/PuerkitoBio/goquery"
)

type Cotacao struct {
	ValorReais float64
}

func NewCotacao() *Cotacao {
	return &Cotacao{}
}

func PostScrape(linkAlvo string, conteudoParaProcurar string, attribute string) *Cotacao {
	conteudoHTML, erro := goquery.NewDocument(linkAlvo)
	utils.CheckErr(erro)

	resultadoDaBusca := conteudoHTML.Find(conteudoParaProcurar).Map(func(index int, item *goquery.Selection) string {
		bitValue, _ := item.Attr(attribute)
		return bitValue
	})
	//fmt.Println(len(resultadoDaBusca))
	cotacao := NewCotacao()
	cotacao.ValorReais, _ = strconv.ParseFloat(strings.Replace(resultadoDaBusca[0], ",", ".", -1), 64)
	return cotacao
}
