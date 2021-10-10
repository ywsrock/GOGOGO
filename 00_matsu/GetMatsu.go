package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/sclevine/agouti"
)

func main() {
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)

	// url := "https://www.matsukiyo.co.jp/store/online"
	url := "https://www.matsukiyo.co.jp/store/online/search?category=003"

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

	err = page.Navigate(url)
	if err != nil {
		log.Printf("Failed to navigate: %v", err)
	}

	// // 商品をさがすボタンクリック
	// err = page.FindByLink("商品をさがす").Click()
	// if err != nil {
	// 	log.Printf("商品探すリンクエラー　%v", err)
	// }
	// 化粧品
	// page.FindByLink("化粧品").Click()
	// time.Sleep(time.Second)
	// time.Sleep(5 * time.Second)
	// content, err := page.HTML()
	// if err != nil {
	// 	log.Printf("Failed to get html: %v", err)
	// }
	// fmt.Println(content)
	// turi, terr := page.URL()
	// if terr != nil {
	// 	fmt.Printf("%+v\n", terr)
	//

	// fmt.Println(turi)
	// 化粧品itemList検索
	// itemList := make([]string, 15)
	itemList := page.FindByID("itemList").AllByClass("ttl").Find("a")
	elements, _ := itemList.Elements()

	for _, elment := range elements {
		vl, _ := elment.GetText()
		fmt.Printf("%+v\n", vl)
		elment.Click()
		url, _ := page.HTML()
		fmt.Printf("%+v\n", url)
		subItem := page.FindByXPath("/html/body/div[4]/div/div[2]/div[1]/div[1]/div[1]/div[1]/ul/li/a")
		i, _ := subItem.Count()
		fmt.Printf("-----%+v\n", i)
		s, _ := subItem.Attribute("style")
		fmt.Println(s)
		// err := subItem.Click()
		// if err != nil {
		// 	log.Println(err)
		// }
		// break

	}
	// count, _ := itemLise.Count()

	// fmt.Println(count)

	// // contentにHTMLが入る
	// content, err := page.HTML()
	// if err != nil {
	// 	log.Printf("Failed to get html: %v", err)
	// }
	// fmt.Println(content)
	// 受信まで、待機
	<-quit

}
