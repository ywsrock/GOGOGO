package main

import (
	"encoding/json"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"log"
	"net/http"
	"os"
	"strings"
)

//	type RankingResponse struct {
//		RankingData struct {
//			RankingDetails []struct {
//				Index    int    `json:"index"`
//				Name     string `json:"name"`
//				Score    int    `json:"score"`
//				UnitName string `json:"unitName"`
//			} `json:"rankingDetails"`
//		} `json:"rankingData"`
//	}
func main() {
	// appID := "dj00aiZpPURyS0lwUVNTRU9YOSZzPWNvbnN1bWVyc2VjcmV0Jng9ZGM-" // 自分のYahooアプリケーションIDに置き換えてください
	// categoryID := "2494"                                                // カテゴリIDを指定します。ここでは「1」はすべてのカテゴリを表します。

	// url := fmt.Sprintf("https://shopping.yahooapis.jp/ShoppingWebService/V2/categoryRanking?appid=%s&category_id=%s&gender=female&generation=20", appID, categoryID)

	url := `https://shopping.yahooapis.jp/ShoppingWebService/V2/categoryRanking?appid=dj00aiZpPURyS0lwUVNTRU9YOSZzPWNvbnN1bWVyc2VjcmV0Jng9ZGM-&category_id=2494&gender=female&generation=20`
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	var response RankingResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		log.Fatal(err)
	}

	// v := make(map[string][]int)
	labels := make([]string, 0)
	// va := make([]int, len(response.CategoryRanking.RankingData))

	items := make(map[string]opts.BarData, 0)

	bar := charts.NewBar()
	// bar.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
	// 	// set some global options like Title/Legend/ToolTip or anything else
	// 	// create a new bar instance
	// 	Title:    "My first bar chart generated by go-echarts",
	// 	Subtitle: "It's extremely easy to use, right?",
	// }))
	for _, d := range response.CategoryRanking.RankingData[0:1] {
		// fmt.Printf("name :%s---> rank:%v\n", d.Name, d.Rank)
		// va = append(va,d.Rank)
		n := strings.Split(d.Name, " ")[0]
		// v[n] = va

		labels = append(labels, n)
		// items = append(items, opts.BarData{Value: d.Rank})
		items[n] = opts.BarData{Value: d.Rank}
		// // Put data into instance
		// da := make([]opts.BarData, 0)
		// da = append(da, opts.BarData{Value: d.Rank})
		// bar.AddSeries(n, da).SetXAxis("1")
	}
	// bar = bar.SetXAxis(labels)
	for key, val := range items {
		bar.AddSeries(key, []opts.BarData{val})
	}

	// f := func() {
	// 	for i, name := range labels {
	//
	// 	}
	// }

	// bar.SetXAxis(labels).AddSeries("", items)
	// AddSeries("Category B", generateBarItems())
	// Where the magic happens
	f, _ := os.Create("bar.html")
	defer f.Close()

	bar.Render(f)

}
