package Chromdb

import (
	"context"
	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/cdproto/dom"
	"github.com/chromedp/chromedp"
	"github.com/olekukonko/tablewriter"
	"log"
	"os"
	"strings"
)

func GetMini(url string) {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()
	chromedp.Run(ctx,
		// open url
		chromedp.Navigate(url),
		chromedp.WaitVisible("#continent3"),
		chromedp.ActionFunc(func(ctx context.Context) error {
			// Htmlのroot dom情報取得
			document, err := dom.GetDocument().Do(ctx)
			if err != nil {
				return err
			}
			// dom.GetOuterHTML().WithNodeID(document.NodeID).Do(ctx)
			// NodeIDからHtml 取得
			html, err := dom.GetOuterHTML().WithNodeID(document.NodeID).Do(ctx)
			if err != nil {
				return err
			}
			// fmt.Println(html)
			toGoquery(html)
			return nil
		}),
	)
}

func toGoquery(html string) {
	r := strings.NewReader(html)
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		log.Fatal("----->", err)
	}
	s := doc.Find("#idx_bond table tbody > tr._2yc_QILv")
	data := make([][]string, s.Length())
	s.Each(func(i int, se *goquery.Selection) {
		childLen := se.Children().Length()
		if childLen == 0 {
			s.Next()
		}
		childNodeValue := make([]string, childLen)
		se.Children().Each(func(i int, selection *goquery.Selection) {
			s := selection.Text()
			childNodeValue[i] = s
		})
		data[i] = childNodeValue
	})
	printTable1(data)
}

func printTable1(data [][]string) {
	table := tablewriter.NewWriter(os.Stdout)
	// table.SetRowLine(true)
	table.SetHeader([]string{"指数名", "時間", "価格", "前日比"})

	for _, v := range data {
		table.Append(v)
	}
	table.Render()
}
