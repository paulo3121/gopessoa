package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
)

type Texto struct {
	Title   string
	Autor   string
	Content string
}

func make_url() string {
	var randint string = strconv.Itoa(rand.Intn(4530) + 6)
	return "http://arquivopessoa.net/textos/" + randint
}

func get_texto(url string) Texto {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)

	author := doc.Find(".autor").Text()
	title := doc.Find(".titulo-texto").Text()

	var texto_poesia []string
	doc.Find(".texto-poesia").Each(func(i int, s *goquery.Selection) {
		texto_poesia = append(texto_poesia, s.Text())
	})

	var texto_prosa []string
	doc.Find(".texto-prosa").Each(func(i int, s *goquery.Selection) {
		texto_prosa = append(texto_prosa, s.Text())
	})

	texto := append(texto_poesia, texto_prosa...)

	return Texto{title, author, texto[0]}
}

func server() {
	fmt.Println("requesting...")

	r := gin.Default()
	r.GET("/api/", func(c *gin.Context) {
		c.JSON(http.StatusOK, get_texto(make_url()))
	})

	r.Run()

}

func main() {
	server()
}
