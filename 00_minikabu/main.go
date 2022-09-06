package main

import (
	"fmt"
	"minikabu/Chromdb"
	task "minikabu/Task"
)

func main() {
	// d := worker.NewWorker(10, 20)
	// d.Start()

	// message := MakeMessage()
	mes1 := MakeMessageFutureIndex()
	Chromdb.GetHtmlValue(mes1)
	Chromdb.GetMini("https://finance.yahoo.co.jp/indices/list")

	// m := [...]*task.Message{message}
	//
	// for _, m := range m {
	// 	d.Message <- m
	// 	d.Wg.Add(1)
	// }
	//
	// d.Wg.Wait()
}

func MakeMessage() *task.Message {

	message := &task.Message{
		Url: "https://kabutan.jp/news/marketnews/",
	}
	s := make([]task.Selector, 0)
	for i := 1; i <= 4; i++ {
		n1, s1 := `指数名：`, `#header_shisuu_big > thead > tr > th:nth-child(%d) > div > a`
		n2, s2 := `価格： `, `#header_shisuu_big > tbody > tr:nth-child(%d) > td:nth-child(%d)`
		n3, s3 := `前日比：`, `#header_shisuu_big > tbody > tr:nth-child(%d) > td:nth-child(%d) > span`
		n4, s4 := `前日比(%)：`, `#header_shisuu_big > tbody > tr.percentage > td:nth-child(%d)`

		s = append(s, task.Selector{n1, fmt.Sprintf(s1, i)})

		if i == 1 {
			s = append(s, task.Selector{n2, fmt.Sprintf(s2, 1, i)})
			s = append(s, task.Selector{n3, fmt.Sprintf(s3, 1, i+1)})
		} else {
			s = append(s, task.Selector{n2, fmt.Sprintf(s2, 1, i+1)})
			s = append(s, task.Selector{n3, fmt.Sprintf(s3, 1, i+2)})
		}
		s = append(s, task.Selector{n4, fmt.Sprintf(s4, i)})
	}

	message.Selectors = s
	return message
}

func MakeMessageFutureIndex() *task.Message {
	message := &task.Message{
		Url: "https://fu.minkabu.jp/index-price",
	}
	selector := make([]task.Selector, 0)
	s := task.Selector{
		Name:  "futures　Index",
		Value: `body > div > div.content > div > div.ly_row > div.main.ly_mr_30 > div.ng-scope > div:nth-child(1) > div > div > table > tbody`,
	}

	selector = append(selector, s)
	message.Selectors = selector

	return message
}
