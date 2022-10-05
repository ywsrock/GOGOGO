package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"golang.org/x/text/width"
	"log"
	"net/http"
	"os"
	"os/signal"
	"regexp"
	"strings"
)

type resultInfo struct {
	info map[int]string
}

func getURL(url string) *resultInfo {
	result := new(resultInfo)
	res, err := http.Get(url)
	if err != nil {
		log.Fatal("http Get err")
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	doc.Find("#summary").Each(func(_ int, s *goquery.Selection) {
		s1 := s.Find(".summaryM").Text()
		s2 := s.Find(".intrst").Text()
		kv := make(map[int]string)
		kv[1] = s1
		kv[2] = s2
		result.info = kv
	})
	return result
}

func main() {
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)

	var key string
	if len(os.Args) > 1 {
		key = os.Args[1]
	} else {
		fmt.Scanln(&key)
	}

	fmt.Println("----")
	targetURL := "https://ejje.weblio.jp/content/" + key

	ret := getURL(targetURL)
	for _, v := range ret.info {
		v = width.Fold.String(v)
		fmt.Printf("%s\n", repFunc(v))
	}
	// 受信まで、待機
	<-quit
}

func repFunc(str string) string {
	str = strings.Replace(str, " ", "", -1)
	reg := regexp.MustCompile(`(?mi)^\s$`)
	st := reg.ReplaceAllString(str, "")
	return st
}
