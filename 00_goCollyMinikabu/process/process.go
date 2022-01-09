package process

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	"log"
	"regexp"
	"sort"
)


type stockInfo struct {
	Label string
	Value string
}
var result = make(map[int]*stockInfo)


/**
　タスク処理
*/
func ProcessTask(taskID interface{}) {
	url := fmt.Sprintf("https://minkabu.jp/stock/%s",taskID)
	GetStockInfo(url, "123")
}

func GetStockInfo(url string, id string) {
	var targetDom = `#sh_field_body > div > div:nth-child(1) > div > div > div.md_box > div > div.ly_row.ly_gutters > dl > div`
	var targetDom2 ="#layout > div:nth-child(4) > div.contents_wrapper.ly_container > div.md_stockBoard > div > div:nth-child(1) > div > div > div.ly_col.ly_colsize_5.md_box.ly_row.md_target_box_group > div.ly_col.ly_colsize_9.md_target_box"
	var targetCurrentValue = `#layout > div:nth-child(4) > div.contents_wrapper.ly_container > div.md_stockBoard > div > div:nth-child(1) > div > div > div.ly_col.ly_colsize_7.md_box.ly_row.ly_gutters > div.ly_col.ly_colsize_5 > div > div:nth-child(2)`


	// Instantiate default collector
	c := colly.NewCollector()

	//// Called before a request
	//c.OnRequest(func(request *colly.Request) {
	//	fmt.Println("Visiting", request.URL.String())
	//})
	c.OnHTML(targetCurrentValue, func(e *colly.HTMLElement) {
		fmt.Println("Visiting", e.Request.URL.String())
		//fmt.Println(e.Text)
		re := regexp.MustCompile(` `)
		re2 := regexp.MustCompile(`(?m)^\s`)

		str := re.ReplaceAllString(e.Text,"")
		str = re2.ReplaceAllString(str,"")
		fmt.Println(str)
	})

	//Called right after OnResponse if the received content is HTML
	c.OnHTML(targetDom, func(e *colly.HTMLElement) {
		//val := e.ChildText(`li >dt`)

		e.ForEach("li > dt", func(i int, element *colly.HTMLElement) {
			//fmt.Println(element.Text)
			result[i] = &stockInfo{
				Label: element.Text,
				Value: "----",
			}
		})

		e.ForEach("li > dd", func(i int, element *colly.HTMLElement) {
			if info, ok := result[i]; ok {
				info.Value = element.Text
				//fmt.Println(info.Label)
			}
		})

		sortKey(result)
		////関数定義
		//funcMap := template.FuncMap{
		//	"sortKey":sortKey,
		//	"getMapV":getValueByKey,
		//}
		//
		//t := template.Must(template.New("resulttext.tmpl").Funcs(funcMap).ParseFiles("./template/resulttext.tmpl"))
		//
		//if ok := t.Execute(os.Stdout,result);ok != nil {
		//	log.Fatalln(ok)
		//}
	})

	//予想情報
	c.OnHTML(targetDom2, func(e *colly.HTMLElement) {
		re := regexp.MustCompile(` `)
		re2 := regexp.MustCompile(`(?m)^\s`)
		str := re.ReplaceAllString(e.Text,"")
		str = re2.ReplaceAllString(str,"")
		fmt.Println(str)
	})

	// Called if error occurred during the request
	c.OnError(func(_ *colly.Response, err error) {
		log.Println(err)
	})

	c.Visit(url)
}

func sortKey(result map[int]*stockInfo)  []int {

	var keys = make([]int, 0, len(result))
	for k := range result {
		keys = append(keys,k)
	}
	sort.Ints(keys)

	for k := range  keys {
		fmt.Printf("%s:%s\n",result[k].Label,result[k].Value)
	}
	return keys
}

func getValueByKey(e interface{},key int) *stockInfo{
	//fmt.Printf("-----%v",e)
	if m,ok := e.(map[int]*stockInfo); ok{
		return m[key]
	}
	return nil
}