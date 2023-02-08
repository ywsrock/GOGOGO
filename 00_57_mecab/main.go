package main

import (
	"fmt"
	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
	"github.com/shogo82148/go-mecab"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
	"log"
	"os"
	"strings"
)

const BOSEOS = "BOS/EOS"

type NodeInfo struct {
	Surface  string
	POS      string
	POS1     string
	POS2     string
	BaseForm string
	PN       float64
}

type OutInfo struct {
	id   string
	text string
	Pn   float64
}

func main() {
	// 形態素解析
	tagger, _ := mecab.New(map[string]string{"output-format-type": "wakati"})
	defer tagger.Destroy()
	text := "見えない"
	// 解析
	dic_list, err := parseToNode(tagger, text)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dic_list)
	// dataframe 作成
	dfc := dataframe.LoadStructs(dic_list)
	// 平均
	meand := dfc.Col("PN").Mean()
	oinfs := []OutInfo{}
	oinfs = append(oinfs, OutInfo{
		id:   "1",
		text: text,
		Pn:   meand,
	})

	for _, v := range oinfs {
		fmt.Printf("%s\t%f\n", v.text, v.Pn)
	}
}

// 形態素解析
func parseToNode(m mecab.MeCab, text string) ([]NodeInfo, error) {
	dicList := make([]NodeInfo, 0)
	node, err := m.ParseToNode(text)
	if err != nil {
		return nil, err
	}
	for ; !node.IsZero(); node = node.Next() {
		if !strings.HasPrefix(node.Feature(), BOSEOS) {
			fts := strings.Split(node.Feature(), ",")
			nodeInfo := NodeInfo{
				Surface:  node.Surface(),
				POS:      fts[0],
				POS1:     fts[1],
				POS2:     fts[2],
				BaseForm: fts[6],
				PN:       parsePN(node.Surface()),
			}
			dicList = append(dicList, nodeInfo)
		}
	}
	return dicList, nil
}

// PN分析
func parsePN(v string) float64 {
	// PN_JA read
	f, err := os.Open("pn_ja.csv")
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}

	// 文字コード変換
	r := transform.NewReader(f, japanese.ShiftJIS.NewDecoder())
	df := dataframe.ReadCSV(r)
	dfl := df.Filter(
		dataframe.F{Colname: "A", Comparator: series.Eq, Comparando: v},
	)

	if dfl.Nrow() == 0 {
		return 0
	}
	if i, ok := dfl.Col("D").Val(0).(float64); ok {
		return i
	}
	return 0
}
