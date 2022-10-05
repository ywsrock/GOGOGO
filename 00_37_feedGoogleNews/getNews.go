package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/mmcdole/gofeed"
)

type Item struct {
	Title     string `json:"title"`
	Link      string `json:"link"`
	Published string `json:"date"`
}

// 出力構造体
var items []*Item

func main() {
	var newUrl = "https://news.google.com/rss/search?hl=ja&gl=JP&ceid=JP:ja&q=%E6%97%A5%E7%B5%8C225"

	println(url.PathUnescape(newUrl))
	resp, err := http.Get(newUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("ReadAll error")
	}

	items = feedNews(body)

	for _, t := range items {
		tm, _ := time.Parse(time.RFC1123, t.Published)
		tm1 := tm.Format("2006/01/02 03:04:05")
		fmt.Printf("%s \t %s \t %s\n", tm1, t.Title, t.Link)
	}
}

func feedNews(body []byte) []*Item {
	fp := gofeed.NewParser()
	feed, _ := fp.ParseString(string(body))
	fmt.Println(feed.Title)

	for _, item := range feed.Items {
		it := new(Item)
		it.Title = item.Title

		if link, err := url.PathUnescape(item.Link); err != nil {
			it.Link = item.Link
		} else {
			it.Link = link
		}

		it.Published = item.Published
		items = append(items, it)
	}
	return items
}
