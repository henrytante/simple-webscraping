package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)


func main()  {
	url := "https://google.com"	
	response, err := http.Get(url)
	if err != nil{
		log.Fatal(err)
	}
	defer response.Body.Close()
	links, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println("Links encontrados : ")
	links.Find("a").Each(func(i int, s *goquery.Selection) {
		link, exists := s.Attr("href")
		if exists {
			fmt.Println(link)
		}
	})
}