package main

import (
	"fmt"
	"regexp"

	"github.com/PuerkitoBio/goquery"
)

//GetPage Image 画像取得
// func GetPage(url string) {
// 	doc, _ := goquery.NewDocument(url)
// 	doc.Find("img").Each(func(_ int, s *goquery.Selection) {
// 		url, _ := s.Attr("src")
// 		fmt.Println(url)
// 	})
// }

//GetPage Image 画像取得
func GetTxt(url string) {
	doc, _ := goquery.NewDocument(url)
	doc.Find(".goodsBox").Each(func(_ int, s *goquery.Selection) {
		txt := s.Text()
		re := regexp.MustCompile(`\t`)
		txt = re.ReplaceAllString(txt, "")
		fmt.Println(txt)
	})
}

//GetDescribe 商品説明
func GetDescribe(url string) {
	doc, _ := goquery.NewDocument(url)
	doc.Find(".ctBox02").Each(func(_ int, s *goquery.Selection) {
		s.Find("")
		txt := s.Text()
		re := regexp.MustCompile(`●`)
		txt = re.ReplaceAllString(txt, "\n●")
		fmt.Println(txt)
	})
}
func main() {
	url := "https://www.matsukiyo.co.jp/store/online/p/4901301360007"
	// imgURL := "https://www.matsukiyo.co.jp/store/online/search?category=00332215"
	// GetPage(imgURL)
	GetTxt(url)
	GetDescribe(url)

}
