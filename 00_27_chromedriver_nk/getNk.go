package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/sclevine/agouti"
)

func getURL(strURL string) {

	// driver := agouti.ChromeDriver(options)
	driver := agouti.ChromeDriver()
	err := driver.Start()

	if err != nil {
		log.Printf("Failed to start driver: %v", err)
	}
	defer driver.Stop()

	page, err := driver.NewPage(agouti.Browser("chrome"))
	if err != nil {
		log.Printf("Failed to open page: %v", err)
	}

	err = page.Navigate(strURL)
	if err != nil {
		log.Printf("Failed to navigate: %v", err)
	}

	content, err := page.HTML()
	if err != nil {
		log.Printf("Failed to get html: %v", err)
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(content))
	if err != nil {
		log.Fatal("goquery error")
	}
	doc.Find("#buyTools").Each(func(index int, s *goquery.Selection) {
		s.Find("input").Each(func(_ int, d *goquery.Selection) {
			_, err := d.Attr("disabled")
			if !err {
				mtText1 := d.Next().Text()
				fmt.Printf("%s \n %s\n", "", mtText1)
			}
		})
	})
}

func main() {
	strURL := `https://www.nike.com/jp/t/%E3%83%8A%E3%82%A4%E3%82%AD-%E3%82%A8%E3%82%A2-%E3%83%95%E3%82%A9%E3%83%BC%E3%82%B9-1-07-lv8-%E3%83%A1%E3%83%B3%E3%82%BA%E3%82%B7%E3%83%A5%E3%83%BC%E3%82%BA-BRxxJ2/DA1343-170`
	getURL(strURL)
	fmt.Println("123")

}
