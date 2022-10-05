package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// GOScrape Test
func GOScrape(key string) {
	// Request the HTML page.
	res, err := http.Get(fmt.Sprintf(`https://m.finance.yahoo.co.jp/stock?code=%s.T`, key))
	if err != nil {
		log.Fatal(err)
	}

	// 最後クローズ
	defer res.Body.Close()
	// ファイル読み込み
	// bs, err := ioutil.ReadAll(res.Body)
	bs, err := io.ReadAll(res.Body)

	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(bs))
	if err != nil {
		log.Fatal(err)
	}

	// 文字抽出
	doc.Find("#detail").Each(func(i int, s *goquery.Selection) {
		s.Find("ul  li").Each(func(n int, s1 *goquery.Selection) {
			// 表示出力に書き出す
			str := outPutFmt(s1.Text())
			fmt.Fprintln(os.Stdout, str)
		})
	})
}

func outPutFmt(str string) string {
	if len(str) == 0 {
		return str
	}
	reg := regexp.MustCompile(`(\d.*)`)
	b := reg.ReplaceAllString(str, `:$1`)
	return string(b)
}

func main() {
	// 入力監視
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("終了は「exit」入力")
	fmt.Println(`株コード入力:`)
	for scanner.Scan() {
		key := scanner.Text()
		if key == "exit" {
			os.Exit(0)
		}
		GOScrape(strings.TrimSpace(key))
	}
}
