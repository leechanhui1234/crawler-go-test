package main

import (
	"fmt"
	"net/http"
	// "strconv"
	iconv "github.com/djimenez/iconv-go"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	url := "https://www.g2b.go.kr:8101/ep/tbid/tbidList.do?searchType=1&bidSearchType=1&taskClCds=5&bidNm=%BA%ED%B7%CF%C3%BC%C0%CE&searchDtType=1&fromBidDt=2018%2F4%2F01&toBidDt=2018%2F5%2F01&setMonth1=1&fromOpenBidDt=&toOpenBidDt=&radOrgan=1&instNm=&instSearchRangeType=&refNo=&area=&areaNm=&strArea=&orgArea=&industry=&industryCd=&upBudget=&downBudget=&budgetCompare=&detailPrdnmNo=&detailPrdnm=&procmntReqNo=&intbidYn=&regYn=Y&recordCountPerPage=100" //나라장터 url
	titles := []string{} //declare array
	res, _ := http.Get(url) //Get Http
		doc, _ := goquery.NewDocumentFromResponse(res) //get html origin

		links := doc.Find(".table_list_tbidTbl tbody tr") //get html tag valuables
		
		links.Each(func(i int, s *goquery.Selection) {
			// For each item found, get the title
			title := s.Find("a").Text() //get a tag variable
			out, _ := iconv.ConvertString(string(title), "euc-kr", "utf-8") //decode
			titles = append(titles, out[14:]) //add variable
		})

		for i := 0; i < len(titles); i++ {
			fmt.Println(titles[i]) //print array valuable
		}
}