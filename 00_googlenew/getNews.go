package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/mmcdole/gofeed"
)

type Item struct {
	Title     string `json:"title"`
	Link      string `json:"link"`
	Published string `json:"date"`
}
type RSS struct {
	Items []Item
}

func main() {
	var new_url string = "https://news.google.com/rss/search?hl=ja&gl=JP&ceid=JP:ja&q=%E6%97%A5%E7%B5%8C225"
	resp, err := http.Get(new_url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	// fmt.Println(string(body))

	if err != nil {
		log.Fatal("ReadAll error")
	}
	fp := gofeed.NewParser()
	feed, _ := fp.ParseString(string(body))
	fmt.Println(feed.Title)
	var items []*Item
	// var items []Item

	for _, item := range feed.Items {
		it := new(Item)
		// it := Item{}
		it.Title = item.Title
		it.Link = item.Link
		it.Published = item.Published
		items = append(items, it)
	}

	for _, t := range items {
		tm, _ := time.Parse(time.RFC1123, t.Published)
		tm1 := tm.Format("2006/01/02 03:04:05")
		fmt.Printf("%s	%s	%s\n", tm1, t.Title, t.Link)
	}

	// fmt.Println("--------------------------------")

	// data,err := json.Marshal(items)
	// if err != nil {
	// 	log.Fatal("json Marshal error")
	// }
	// fmt.Printf("%#v",string(data))
}
