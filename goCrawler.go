package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	url := "https://yepp0517.tistory.com/"

	for i := 2; i < 5; i++ {
		res, _ := http.Get(url + strconv.Itoa(i))
		doc, _ := goquery.NewDocumentFromResponse(res)
		fmt.Println(doc.Find("h4").Text())
	}
}