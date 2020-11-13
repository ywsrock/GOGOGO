package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/PuerkitoBio/goquery"
)

// GOScrape Test
func GOScrape() {
	// Request the HTML page.
	res, err := http.Get("https://m.finance.yahoo.co.jp/stock?code=1357.T")
	if err != nil {
		log.Fatal(err)
	}

	//最後クローズ
	defer res.Body.Close()
	//ファイル読み込み
	bs, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	//ファイル新規作成
	// f, err := os.Create("aa.txt")
	//ファイルに書き出す
	// fmt.Fprint(f, string(bs))

	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(bs))
	if err != nil {
		log.Fatal(err)
	}

	// 文字抽出
	doc.Find("#detail-module").Each(func(i int, s *goquery.Selection) {
		s.Find("div.stockDetail  li").Each(func(n int, s1 *goquery.Selection) {
			//表示出力に書き出す
			fmt.Fprintln(os.Stdout, s1.Text())
		})
	})
}

func main() {
	GOScrape()
	//入力監視
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("終了「exit」入力")
	for scanner.Scan() {
		key := scanner.Text()
		if key == "exit" {
			os.Exit(0)
		}
	}
}
