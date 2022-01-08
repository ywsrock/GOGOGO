package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/geziyor/geziyor"
	"github.com/geziyor/geziyor/client"
	"golang.org/x/net/html"
	"os"
	"os/signal"
	"strings"

	"github.com/gocolly/colly/v2"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	//NG Forbidden by robots.
	//geziyorQuery()
	//OK http.get
	//get()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)

	var key string
	if len(os.Args) > 1 {
		key = os.Args[1]
	} else {
		fmt.Scanln(&key)
	}

	collyQuery(key)

	// 受信まで、待機
	<-quit

}

//NG Forbidden by robots.
func geziyorQuery() {
	geziyor.NewGeziyor(&geziyor.Options{
		//StartURLs: []string{"https://eow.alc.co.jp/search?q=marshal"},
		StartRequestsFunc: func(g *geziyor.Geziyor) {
			req, _ := client.NewRequest("GET", "https://eow.alc.co.jp/search?q=marshal", nil)
			//req.Meta["key"] = "value"
			g.Do(req, g.Opt.ParseFunc)
		},
		ParseFunc: func(g *geziyor.Geziyor, r *client.Response) {
			fmt.Println(string(r.Body))
		},
	}).Start()
}

func get() {
	req, _ := http.NewRequest("GET", "https://eow.alc.co.jp/search?q=marshal", nil)
	client := new(http.Client)
	res, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}
	defer res.Body.Close()
	b, _ := ioutil.ReadAll(res.Body)
	fmt.Println(string(b))
}

func collyQuery(key string) {
	c := colly.NewCollector()

	//Find and visit all links
	c.OnHTML("#resultsList > ul", func(e *colly.HTMLElement) {
		e.DOM.Children().Each(func(i int, selection *goquery.Selection) {
			if i == 0 {
				selection.Children().Each(func(i int, selection *goquery.Selection) {
					for _, node := range selection.Nodes {
						printChild(node, selection)
					}
				})
			}
		})
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit("https://eow.alc.co.jp/search?q=" + key)
}

func printChild(node *html.Node, selection *goquery.Selection) {
	//fmt.Println(node.Type)
	if node.Type == html.TextNode {
		fmt.Print(node.Data)
	}

	for n := node.FirstChild; n != nil; n = n.NextSibling {
		if strings.ToLower(n.Data) == "br" || strings.ToLower(n.Data) == "li" {
			fmt.Println()
		}
		printChild(n, selection)
	}
}
