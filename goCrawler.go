package main

import (
	"fmt"
	"net/http"
	// "strconv"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	url := "https://yepp0517.tistory.com/2"

	res, _ := http.Get(url)
		doc, _ := goquery.NewDocumentFromResponse(res)

		links := doc.Find(".tt_article_useless_p_margin")
		
		links.Each(func(i int, s *goquery.Selection) {
			// For each item found, get the title
			title := s.Find("#og_1627608827011").Text()
			fmt.Printf("Review %d: %s\n", i, title)
		})
}