package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type XML struct {
	Book []Bookmarks `xml:"item"`
}

type Bookmarks struct {
	Title string `xml:"title"`
	Link  string `xml:"link"`
	Date  string `xml:"date"`
	Count int    `xml:"bookmarkcount"`
}

func main() {
	data := httpGet("http://b.hatena.ne.jp/hotentry/it.rss")

	result := XML{}
	err := xml.Unmarshal([]byte(data), &result)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}


	for _, bookmark := range result.Book {
		datetime, _ := time.Parse(time.RFC3339, bookmark.Date)

		fmt.Printf("%v\n", datetime.Format("2006/01/02 15:04:05"))
		fmt.Printf("%s - %duser\n", bookmark.Title, bookmark.Count)
		fmt.Printf("%v\n", bookmark.Link)
		fmt.Println()
	}
}

func httpGet(url string) string {
	response, _ := http.Get(url)
	body, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()
	return string(body)
}
