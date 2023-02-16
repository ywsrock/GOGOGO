package Chromdb

import (
	"context"
	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/cdproto/dom"
	"github.com/chromedp/chromedp"
	"github.com/olekukonko/tablewriter"
	"log"
	task "minikabu/Task"
	"os"
	"strings"
	"time"
)

func GetEventList(mes *task.Message) {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	chromedp.Run(ctx,
		chromedp.Navigate(mes.Url),
		chromedp.WaitVisible(mes.Selectors[0].Value),

		chromedp.ActionFunc(func(ctx context.Context) error {
			chromedp.Click("#str-main > div.blink > div:nth-child(2) > div > a", chromedp.NodeVisible).Do(ctx)
			time.Sleep(5 * time.Second)
			// chromedp.WaitReady(mes.Selectors[0].Value)
			// chromedp.RunResponse(ctx, chromedp.Click("#str-main > div.blink > div:nth-child(2) > div > a", chromedp.NodeVisible))
			root, err := dom.GetDocument().Do(ctx)
			if err != nil {
				return err
			}
			html, err := dom.GetOuterHTML().WithNodeID(root.NodeID).Do(ctx)
			if err != nil {
				return err
			}
			findValue(html, mes)
			return nil

		}),
	)
}

func findValue(html string, mes *task.Message) {
	r := strings.NewReader(html)

	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		log.Fatal("findValue--.", err)
	}
	sec := doc.Find(mes.Selectors[0].Value)
	header := make([]string, 0)
	data := make([][]string, sec.Children().Length())

	sec.Children().Each(func(i int, s *goquery.Selection) {
		if i == 0 {
			s.Children().Each(func(j int, ts *goquery.Selection) {
				if j == 3 {
					header = append(header, "国")
				} else {
					header = append(header, ts.Text())
				}
			})
			sec.Next()
		} else {
			cv := make([]string, s.Children().Length())
			s.Children().Each(func(j int, ts *goquery.Selection) {
				cv[j] = strings.ReplaceAll(strings.TrimSpace(ts.Text()), "★", "＠")
			})
			data[i] = cv
		}
	})
	printEvent(header, data)
}

func printEvent(header []string, data [][]string) {
	table := tablewriter.NewWriter(os.Stdout)
	// table.SetRowLine(true)
	table.SetHeader(header)

	for _, v := range data {
		table.Append(v)
	}
	table.Render()
}
