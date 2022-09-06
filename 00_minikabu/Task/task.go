package task

import (
	"fmt"
	"github.com/gocolly/colly"
	"log"
	"strings"
)

type Message struct {
	Url       string
	Selectors []Selector
}

type Selector struct {
	Name  string
	Value string
}

type MyColly struct {
	c *colly.Collector
}

func NewMyColly() *MyColly {
	m := &MyColly{
		c: colly.NewCollector(),
	}
	return m
}

func (m *MyColly) VistUrl(url string) {
	m.c.Visit(url)
}

func (m *MyColly) Request() {
	m.c.OnRequest(func(request *colly.Request) {
		fmt.Println(request.URL)
	})
}

func (m *MyColly) Html(mes *Message) {
	for _, s := range mes.Selectors {
		s := s
		f := func(e *colly.HTMLElement) {
			fmt.Printf("%s:%s\n", s.Name, e.Text)
			if strings.Contains(s.Name, "(") {
				fmt.Println("------------")
			}
		}
		m.c.OnHTML(s.Value, f)
	}
}

func StartJob(Job interface{}) {
	mes, ok := Job.(*Message)
	if !ok {
		log.Fatal(ok)
		return
	}
	m := NewMyColly()
	m.Request()

	m.Html(mes)

	m.c.Visit(mes.Url)
}
