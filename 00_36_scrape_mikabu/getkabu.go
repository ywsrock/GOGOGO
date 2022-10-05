package main

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"golang.org/x/net/html"
	"os"
	"os/signal"
	"strings"
)

var depth int

func main() {
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	cKey := make(chan string)
	go func() {
		var key string
		for {
			if len(os.Args) > 1 {
				key = os.Args[1]
				cKey <- key
			} else {
				fmt.Scanln(&key)
				if key == "exit" {
					quit <- os.Interrupt
					break
				}
				cKey <- key
			}
		}
	}()

	go func() {
		for {
			select {
			case code := <-cKey:
				getKabulColly(code)
			}
		}
	}()

	<-quit
}

func getKabulColly(code string) {
	if strings.TrimSpace(code) == "" {
		code = "1570"
	}

	c := colly.NewCollector()
	// Request
	c.OnRequest(func(request *colly.Request) {
		fmt.Println(request.URL)
	})

	// stock-detail
	// c.OnHTML(".stock-detail", func(element *colly.HTMLElement) {
	c.OnHTML(".ly_col.ly_colsize_5.md_box.ly_row.md_target_box_group", func(element *colly.HTMLElement) {
		childNode(element.DOM.Get(0), nil, nil)
	})
	c.Visit("https://minkabu.jp/stock/" + code)
}

func childNode(node *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(node)
	}
	if node.Type == html.TextNode {
		if strings.TrimSpace(node.Data) != "" {
			fmt.Println(node.Data)
		}
	}
	for n := node.FirstChild; n != nil; n = n.NextSibling {
		if strings.ToLower(n.Data) == "br" || strings.ToLower(n.Data) == "li" {
			fmt.Println()
		}
		childNode(n, pre, post)
	}

	if post != nil {
		post(node)
	}
}

/*
*
Node開始
*/
func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		fmt.Printf("%*s<%s class=%v>\n", depth*2, "", n.Data, n.Attr)
		depth++
	}
}

/*
*
Node終了
*/
func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
	}

}
