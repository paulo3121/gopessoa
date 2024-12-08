package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

func make_url() string {
	var randint string = strconv.Itoa(rand.Intn(4530) + 6)
	url := "http://arquivopessoa.net/textos/" + randint
	return url
}

func get_texto(url string) {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)

	autor := doc.Find(".autor").Text()
	title := doc.Find(".titulo-texto").Text()

	var texto_poesia []string
	doc.Find(".texto-poesia").Each(func(i int, s *goquery.Selection) {
		texto_poesia = append(texto_poesia, s.Text())
	})

	var texto_prosa []string
	doc.Find(".texto-prosa").Each(func(i int, s *goquery.Selection) {
		texto_prosa = append(texto_prosa, s.Text())
	})

	fmt.Println()
	fmt.Println("title: ", title)
	if len(texto_poesia) == 1 {
		fmt.Println("poesia:", texto_poesia[0])
	}
	if len(texto_prosa) == 1 {
		fmt.Println("prosa:", texto_prosa[0])
	}

	fmt.Println("autor: ", autor)
}

func main() {
	get_texto(make_url())
}

/*
   0. calcular duração da requisição
   1. adicionar interação para gerar novo texto e para sair
   2. criar arquivo com todos os textos
*/
