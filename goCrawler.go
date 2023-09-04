package main

import (
	"fmt"
	"net/http"
	// "strconv"
	iconv "github.com/djimenez/iconv-go"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	url := "https://www.g2b.go.kr:8101/ep/tbid/tbidList.do?searchType=1&bidSearchType=1&taskClCds=5&bidNm=%BA%ED%B7%CF%C3%BC%C0%CE&searchDtType=1&fromBidDt=2018%2F2%2F01&toBidDt=2018%2F3%2F01&setMonth1=1&fromOpenBidDt=&toOpenBidDt=&radOrgan=1&instNm=&instSearchRangeType=&refNo=&area=&areaNm=&strArea=&orgArea=&industry=&industryCd=&upBudget=&downBudget=&budgetCompare=&detailPrdnmNo=&detailPrdnm=&procmntReqNo=&intbidYn=&regYn=Y&recordCountPerPage=100"
	titles := []string{}
	res, _ := http.Get(url)
		doc, _ := goquery.NewDocumentFromResponse(res)

		links := doc.Find(".table_list_tbidTbl tbody tr")
		
		links.Each(func(i int, s *goquery.Selection) {
			// For each item found, get the title
			title := s.Find("a").Text()
			out, _ := iconv.ConvertString(string(title), "euc-kr", "utf-8")
			titles = append(titles, out)
		})

		fmt.Println(titles)
}