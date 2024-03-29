package Chromdb

import (
	"context"
	"fmt"
	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/dom"
	"github.com/chromedp/chromedp"
	"github.com/olekukonko/tablewriter"
	"log"
	task "minikabu/Task"
	"os"
	"strings"
)

func GetHtmlValue(mes *task.Message) {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// Open Url
	if err := chromedp.Run(ctx, chromedp.Navigate(mes.Url)); err != nil {
		log.Fatal("Open URL-->", err)
	}

	sel := mes.Selectors[0].Value
	// wait page loaded
	if err := chromedp.Run(ctx, chromedp.WaitVisible(sel)); err != nil {
		log.Fatal("wait selector")
	}

	// Get Node
	var nodes []*cdp.Node
	var ids []cdp.NodeID
	// Get child Nodes
	err1 := chromedp.Run(ctx,
		// 指定セレクターのコードID取得
		// chromedp.NodeIDs(sel, &ids, chromedp.ByQuery),
		chromedp.NodeIDs(sel, &ids, chromedp.BySearch),
		// 指定NodeIdの子のNode取得
		chromedp.ActionFunc(func(ctx context.Context) error {
			// Get the childNodes
			return dom.RequestChildNodes(ids[0]).WithDepth(-1).Do(ctx)
		}),
		// Node取得
		chromedp.Nodes("tr", &nodes, chromedp.ByQueryAll),
	)
	if err1 != nil {
		log.Fatal("sub Node-->", err1)
	}

	// print table　format
	data := make([][]string, 0)
	bs := &strings.Builder{}
	// 　Dispose of child nodes
	GetChild(nodes, bs)
	sl := strings.Split(bs.String(), "\n")
	for _, s := range sl {
		if strings.Trim(s, "@") == "" {
			continue
		}
		s := strings.TrimRight(s, "@")
		// Display　Cell　Array
		td := strings.Split(s, "@")
		es := strings.TrimSpace(td[0])
		// unify the array
		if es == "日経225オプション" || strings.Contains(es, "権利行使価格 ") {
			if len(td) < 7 {
				for i := len(td); i <= 7; i++ {
					td = append(td, "")
				}
			}
			td[1] = fmt.Sprintf("%s %s", td[1], td[2])
			for i := 2; i <= 5; i++ {
				td[i] = td[i+1]
			}
			// td[2] = td[3]
			// td[3] = td[4]
			// td[4] = td[5]
			// td[5] = td[6]
			// slice に変換
			td = td[0:6]
		}
		data = append(data, td)
	}
	printTable(data)
}

func GetChild(nodes []*cdp.Node, bs *strings.Builder) {
	for _, node := range nodes {
		if node.LocalName == "tr" {
			bs.WriteString("\n")
		}
		if node.LocalName == "th" {
			continue
		}
		if node.NodeType == cdp.NodeTypeText {
			s := fmt.Sprintf("%s@", node.NodeValue)
			bs.WriteString(s)
		}
		if node.ChildNodeCount > 0 {
			GetChild(node.Children, bs)
		}
	}
}

func printTable(data [][]string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"先物名", "現在値", "変動率", "始値", "高値", "安値"})
	for _, v := range data {
		table.Append(v)
	}
	table.Render()
}
