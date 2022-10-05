package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"

	"github.com/sclevine/agouti"
)

func main() {
	options := agouti.ChromeOptions(
		"args", []string{
			"--headless",
			"--disable-gpu", // 暫定的に必要らしいです。
		})

	driver := agouti.ChromeDriver(options)
	defer driver.Stop()
	driver.Start()

	page, err := driver.NewPage()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		return
	}

	page.Navigate("https://www.tokovalue.jp/API_INDEX_G.htm")
	fmt.Println(page.Title())
	page.FindByLink("GdiAlphaBlend").Click()
	fmt.Println(page.Title())
	str, _ := page.HTML()
	res := strings.NewReader(str)
	dom, _ := goquery.NewDocumentFromReader(res)
	fmt.Println(dom.Text())

}
