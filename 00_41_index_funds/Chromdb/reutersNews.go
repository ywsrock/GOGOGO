package Chromdb

import (
	"fmt"
	"github.com/gocolly/colly"
	"html/template"
	"log"
	"os"
	"strings"
)

type Target struct {
	Url string
}

var temp = `
{{ range $i, $m := . }}
No: {{$i}}  {{$m.Title}}    {{$m.Time}} {{$m.Url}}
{{$m.Content}}
{{end}}
`

type Ret struct {
	Title   string
	Content string
	Time    string
	Url     string
}

var Rets = make(map[int]Ret)

// var selectors = "#moreSectionNews > section > div > article:nth-child(1) > div.story-content"
var selectors = "#moreSectionNews > section > div > article > div.story-content"
var baseUrl = `https://jp.reuters.com`

func GetNews(url string) {
	c := colly.NewCollector()
	c.OnRequest(func(request *colly.Request) {
		fmt.Printf("request URL:%s", request.URL)
	})

	Count := 1

	c.OnHTML(selectors, func(e *colly.HTMLElement) {
		// Title
		t := e.ChildText("a > h3")
		// content
		content := e.ChildText("p")
		// Time
		articleTime := e.ChildText("time > span")

		// Url https://jp.reuters.com/article/usa-economy-inflation-idJPKBN2R818P
		s := strings.Builder{}
		s.WriteString(baseUrl)
		s.WriteString(e.ChildAttr("a", "href"))
		url := s.String()
		s.Reset()

		Rets[Count] = Ret{Title: t, Content: CoutContent(content), Time: articleTime, Url: url}
		Count += 1
	})

	c.Visit(url)

	f := template.FuncMap{
		"Index": func(v int) int {
			return v + 1
		},
		"CoutContent": CoutContent,
	}

	t, err := template.New("Out").Funcs(f).Parse(temp)
	if err != nil {
		// log.Fatal(err)
	}

	err = t.Execute(os.Stdout, Rets)
	if err != nil {
		log.Fatal(err)
	}
}

func CoutContent(content string) string {
	str := ""
	c := 0
	for _, v := range content {
		if c == 80 {
			str += string(v) + "\n"
			c = 0
		} else {
			str += string(v)
			c++
		}
	}
	return str
}
