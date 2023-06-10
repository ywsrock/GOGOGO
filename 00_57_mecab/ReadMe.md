## Golangで形態素解析

Golangでも簡単で日本語の形態素解析ができます。（環境構築はGoogle先生にお願いします。）





## コードの説明

## パッケージと定数

```go
package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/go-gota/gota/dataframe"
	"github.com/go-gota/gota/series"
	"github.com/shogo82148/go-mecab

"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)
```

上記のコードは、`main` パッケージを使用しています。また、以下のパッケージもインポートされています。

- `fmt`：標準出力のためのパッケージ
- `log`：エラーログのためのパッケージ
- `os`：ファイル操作のためのパッケージ
- `strings`：文字列操作のためのパッケージ
- `github.com/go-gota/gota/dataframe`：データフレーム操作のためのパッケージ
- `github.com/go-gota/gota/series`：シリーズ操作のためのパッケージ
- `github.com/shogo82148/go-mecab`：形態素解析のためのパッケージ
- `golang.org/x/text/encoding/japanese`：日本語のエンコーディングのためのパッケージ
- `golang.org/x/text/transform`：テキストの変換のためのパッケージ

```go
const BOSEOS = "BOS/EOS"
```

`BOSEOS` は定数として定義されています。この定数は、形態素解析結果の特徴の先頭にある特別なタグを表しています。



## 構造体

```go
type NodeInfo struct {
	Surface  string
	POS      string
	POS1     string
	POS2     string
	BaseForm string
	PN       float64
}
```

`NodeInfo` 構造体は、形態素解析結果のノード情報を格納するために使用されます。各フィールドは以下のような情報を保持します。

- `Surface`：表層形（文字列）
- `POS`：品詞（文字列）
- `POS1`：品詞細分類1（文字列）
- `POS2`：品詞細分類2（文字列）
- `BaseForm`：基本形（文字列）
- `PN`：PN値（浮動小数点数）

```go
type OutInfo struct {
	id   string
	text string
	Pn   float64
}
```

`OutInfo` 構造体は、解析結果を格納するために使用されます。各フィールドは以下のような情報を保持します。

- `id`：ID（文字列）
- `text`：テキスト（文字列）
- `Pn`：PN値（浮動小数点数）



## `main` 関数

```go
func main() {
	tagger, _ := mecab.New(map[string]string{"output-format-type": "wakati"})
	defer tagger.Destroy()
	text := "見えない"

	dicList, err := parseToNode(tagger, text)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(dicList)

	dfc := dataframe.LoadStructs(dicList

)
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
```

`main` 関数はプログラムのエントリーポイントです。以下の処理を行います。

1. `mecab.New` 関数を使用して形態素解析器（`tagger`）を作成します。また、`defer` ステートメントを使用して関数終了時に `tagger.Destroy()` を呼び出し、解放処理を行います。
2. 解析対象のテキスト（`text`）を定義します。
3. `parseToNode` 関数を呼び出して形態素解析を行い、解析結果を `dicList` に格納します。
4. 解析結果を表示します。
5. `dataframe.LoadStructs` 関数を使用してデータフレーム（`dfc`）を作成します。
6. データフレームの `PN` 列の平均値を計算し、`meand` に代入します。
7. 解析結果を `OutInfo` のスライス（`oinfs`）に追加します。
8. `oinfs` の要素をループして、テキストと PN 値を表示します。



## 形態素解析関数

```go
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
```

`parseToNode` 関数は、形態素解析を行い、解析結果を `NodeInfo` のスライスとして返します。引数として形態素解析器（`m`）と解析対象のテキスト（`text`）を受け取ります。

1. 空の `dicList` スライスを作成します。
2. `m.ParseToNode` 関数を使用してテキストを解析し、解析結果の最初のノードを取得します。
3. ノードが存在する限り、以下の処理を繰り返します。
   -

 ノードの特徴を取得し、`BOSEOS` で始まっていないかを確認します。
   - ノードの特徴をカンマで分割し、各フィールドに割り当てます。
   - `NodeInfo` 構造体を作成し、解析結果の情報をフィールドに格納します。
   - `dicList` に `NodeInfo` を追加します。
4. 解析結果の `dicList` を返します。

## PN分析関数

```go
func parsePN(v string) float64 {
	f, err := os.Open("pn_ja.csv")
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}

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
```

`parsePN` 関数は、PN値を分析するために使用されます。引数として解析対象の文字列（`v`）を受け取り、その文字列に対応するPN値を返します。

1. `"pn_ja.csv"` ファイルを開きます。処理が終了したら `defer` ステートメントでファイルを閉じます。
2. 日本語のシフトJISエンコーディングを使用して、ファイルをテキストとして読み込むための `transform.NewReader` を作成します。
3. `dataframe.ReadCSV` 関数を使用して、CSVファイルをデータフレーム（`df`）として読み込みます。
4. データフレームのフィルタリングを行い、列 `"A"` の値が解析対象の文字列と等しい行のみを取得します。
5. 取得したデータフレーム（`dfl`）の行数が0であれば、PN値は0として返します。
6. `dfl.Col("D").Val(0)` を使用して、列 `"D"` の最初の値を取得します。
7. 値が浮動小数点数であれば、その値をPN値として返します。そうでなければ、0を返します。



以上、です！