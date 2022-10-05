package handler

import (
	"bytes"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/html"
	"log"
	"net/http"
	"pageredder/db"
	"pageredder/model"
	"strings"
)

// var history = make(map[int]*model.WordInfo)
// var mapLen = 0
// var ss string

func RouterHandler(ctx *gin.Context) {
	ws := make([]*model.WordInfo, 0)
	// 検索履歴
	history := db.FindAll()
	// 検索キーワード
	q := ctx.DefaultQuery("key", " ")
	if strings.TrimSpace(q) == "" {
		ctx.HTML(http.StatusOK, "index.html", gin.H{
			"w": model.WordInfo{Word: "", Info: "Have a Nice Day!", InfosJp: "", VoiceLink: ""},
			"h": history,
		})
		return
	}

	// 検索履歴からキーワードを検察
	w, b := db.FindFirst(q)
	if b == true {
		ws = append(ws, w)
	} else {
		// dic weblio
		ws = ReqWeblio(q, ws)
		ws[0].InfosJp = ReqGoo(q)
		ws[0].InfosEn = ""
		db.SaveWord(ws[0])
	}

	// 全部検索履歴
	history = db.FindAll()
	ctx.HTML(http.StatusOK, "index.html", gin.H{
		"w": ws[0],
		"h": history,
	})
}

/**
From dictionary Weblio
*/
func ReqWeblio(q string, ws []*model.WordInfo) []*model.WordInfo {
	// weblio
	doc := SendRequest("https://ejje.weblio.jp/content/%s", q)
	w := GetWeblioRes(doc)
	ws = append(ws, w)
	return ws
}

/**
From dictionary Goo
*/

func ReqGoo(q string) string {
	// Goo
	doc := SendRequest("https://dictionary.goo.ne.jp/word/en/%s/", q)
	w := GetGooMeansList(doc)
	return w
}

/**
Send The HTTP Request
*/
func SendRequest(gurl string, word string) *goquery.Document {
	url := fmt.Sprintf(gurl, word)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal()
	}

	c := new(http.Client)
	res, err := c.Do(req)
	if err != nil {
		log.Fatal()
	}

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	return doc
}

/**

Goquery from weblio
*/
func GetWeblioRes(doc *goquery.Document) *model.WordInfo {
	w := &model.WordInfo{}
	// query word expression selector
	// wInfo := "#summary > div.summaryM.descriptionWrp > p > span"
	wInfo := "#summary > div.summaryM.descriptionWrp > p > span.content-explanation.ej"
	// query word voice selector
	wVoiceUrl := `#summary > table.summaryTbl > tbody > tr > td.summaryC > table > tbody > 
	tr:nth-child(1) > td:nth-child(1) > i > audio > source`
	// query word selector
	keyword := "#h1Query"
	// query word phonetic symbols
	vPhoneticSymbol := "#phoneticEjjeNavi > div"
	// query word syllable
	wSyllable := "#summary > div:nth-child(4) > span:nth-child(1) > span"

	// word keyword
	doc.Find(keyword).Each(func(_ int, s *goquery.Selection) {
		w.Word = s.Text()
	})
	// query word syllable
	doc.Find(wSyllable).Each(func(_ int, s *goquery.Selection) {
		w.Syllable = s.Text()
	})
	//  word phonetic symbols
	doc.Find(vPhoneticSymbol).Each(func(_ int, s *goquery.Selection) {
		w.PhoneticSymbols = s.Text()
	})

	// word expression
	doc.Find(wInfo).Each(func(_ int, s *goquery.Selection) {
		w.Info = strings.TrimSpace(s.Text())
	})
	// word voice link
	doc.Find(wVoiceUrl).Each(func(_ int, s *goquery.Selection) {
		v, _ := s.Attr("src")
		w.VoiceLink = strings.TrimSpace(v)
	})

	return w
}

/**
Goquery from Goo
*/
func GetGooMeansList(doc *goquery.Document) string {
	var infos bytes.Buffer
	ch := make(chan string, 0)
	// infos := make([]string, 0)A
	defer close(ch)
	// query word expression selector
	wInfo := "#NR-main > section > div > div.contents-wrap-b.meanging > div > div.content-box.content-box-ej > ol:nth-child(2)"
	// #NR-main > section > div > div.contents-wrap-b.meanging > div > div.content-box.content-box-ej > ol:nth-child(2)
	// word expression
	doc.Find(wInfo).Each(func(_ int, s *goquery.Selection) {
		go func() {
			for s := range ch {
				if strings.TrimSpace(s) != "" {
					// s = strings.ReplaceAll(s,"<br>"," ")
					infos.WriteString(s)
					infos.WriteString("\n")
				}
			}
		}()

		for _, node := range s.Nodes {
			printChild(node, s, ch)
		}
		// close(ch)
	})

	return infos.String()
}

func printChild(node *html.Node, selection *goquery.Selection, ch chan string) {
	if strings.ToLower(node.Data) == "li" {
		ch <- goquery.NewDocumentFromNode(node).Text()
	}

	// if node.Type == html.TextNode {
	//	ch <- node.Data
	// }

	for n := node.FirstChild; n != nil; n = n.NextSibling {
		if strings.ToLower(n.Data) == "br" || strings.ToLower(n.Data) == "li" {
			fmt.Println()
		}
		printChild(n, selection, ch)
	}
}

func CountRowCount(i int) int {
	return i + 1
}

func String2Slice(str string) []string {
	s := strings.Split(str, "\n")
	return s
}
