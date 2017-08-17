package scraper

import (
	// import standard libraries

	"log"
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
	log.Printf("loading %s", linkAlvo)
	conteudoHTML, erro := goquery.NewDocument(linkAlvo)
	utils.CheckErr(erro, "erro ao carregar pagina "+linkAlvo)
	log.Printf("Page " + linkAlvo + " loaded successfully")
	resultadoDaBusca := conteudoHTML.Find(conteudoParaProcurar).Map(func(index int, item *goquery.Selection) string {
		bitValue, _ := item.Attr(attribute)

		return bitValue
	})
	if len(resultadoDaBusca) == 0 {
		log.Println("Nenhum resultado encontrado para "+conteudoParaProcurar, " atributo "+attribute)
		return NewCotacao()
	}
	cotacao := NewCotacao()
	cotacao.ValorReais, _ = strconv.ParseFloat(strings.Replace(resultadoDaBusca[0], ",", ".", -1), 64)
	return cotacao

}
