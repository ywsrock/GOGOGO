package main

import (
	"flag"
	"fmt"
	"minikabu/Chromdb"
	. "minikabu/Task"
	"strings"
)

func main() {

	var flg string

	flag.StringVar(&flg, "d", "2", "表示内容 1:News 2:Index 3:Calendar 4:ALL")
	flag.Parse()

	f := strings.TrimSpace(flg)
	if f == "1" || f == "4" {
		Url := `https://jp.reuters.com/news/archive/special20`
		Chromdb.GetNews(Url)
	}
	if f == "2" || f == "4" {
		mes1 := MakeMessageFutureIndex()
		Chromdb.GetHtmlValue(mes1)
		Chromdb.GetMini("https://finance.yahoo.co.jp/indices/list")
	}

	if f == "3" || f == "4" {
		mes1 := EventCalendar()
		Chromdb.GetEventList(mes1)
	}
}

func MakeMessage() *Message {

	message := &Message{
		Url: "https://kabutan.jp/news/marketnews/",
	}
	s := make([]Selector, 0)
	for i := 1; i <= 4; i++ {
		n1, s1 := `指数名：`, `#header_shisuu_big > thead > tr > th:nth-child(%d) > div > a`
		n2, s2 := `価格： `, `#header_shisuu_big > tbody > tr:nth-child(%d) > td:nth-child(%d)`
		n3, s3 := `前日比：`, `#header_shisuu_big > tbody > tr:nth-child(%d) > td:nth-child(%d) > span`
		n4, s4 := `前日比(%)：`, `#header_shisuu_big > tbody > tr.percentage > td:nth-child(%d)`

		s = append(s, Selector{Name: n1, Value: fmt.Sprintf(s1, i)})

		if i == 1 {
			s = append(s, Selector{Name: n2, Value: fmt.Sprintf(s2, 1, i)})
			s = append(s, Selector{Name: n3, Value: fmt.Sprintf(s3, 1, i+1)})
		} else {
			s = append(s, Selector{Name: n2, Value: fmt.Sprintf(s2, 1, i+1)})
			s = append(s, Selector{Name: n3, Value: fmt.Sprintf(s3, 1, i+2)})
		}
		s = append(s, Selector{Name: n4, Value: fmt.Sprintf(s4, i)})
	}

	message.Selectors = s
	return message
}

func MakeMessageFutureIndex() *Message {
	message := &Message{
		Url: "https://fu.minkabu.jp/index-price",
	}
	selector := make([]Selector, 0)
	s := Selector{
		Name: "futures　Index",
		// Value: `body > div > div.content > div > div.ly_row > div.main.ly_mr_30 > div.ng-scope > div:nth-child(1) > div > div > table > tbody`,
		// Value: `table.ui-table > tbody`,
		Value: `/html/body/div/div/div/div[3]/div[1]/div[3]/div[1]/div/div/table/tbody`,
	}

	selector = append(selector, s)
	message.Selectors = selector

	return message
}

func EventCalendar() *Message {
	message := &Message{Url: `https://www.rakuten-sec.co.jp/web/market/calendar/`}
	selector := make([]Selector, 0)
	s := Selector{
		Name:  "Event Calendar",
		Value: `#CALTBL > table > tbody`,
	}

	selector = append(selector, s)
	message.Selectors = selector

	return message
}
