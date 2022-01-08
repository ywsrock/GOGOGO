package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("")

	var what string
	//ショートカットキー(Mac)：Option + Enter
	what = "1234"
	fmt.Println(what)

	getUrl("1234")

	//ショートカットキー(Mac)：Command + Shift + A
	//■ コマンド検索(Find Action)

	//コピー履歴から貼り付け(Paste from History)
	//ショートカットキー(Mac)：Command + Shift + V

	//■ プロジェクトツールウィンドウの表示/非表示(Structure)
	//ショートカットキー(Mac)：Command + 1

	//■ ストラクチャーツールウィンドウの表示/非表示
	//ショートカットキー(Mac)：Command + 7

	//■ 現在開いてるファイルのプロジェクト/パッケージへジャンプ(Select In)
	//ショートカットキー(Mac)：Option + F1

	//■ コードをフォーマットする(Reformat Code)
	//ショートカットキー(Mac)：Command + Option + L

	//■ 選択範囲の型を表示(Type Info) -------選択範囲の戻り値
	//ショートカットキー(Mac)：Ctrl + Shift + P

	//■ 変数の抽出(Introduce Variable)
	//ショートカットキー(Mac)：Command + Option + V 変数
	//ショートカットキー(Mac)：Command + Option + M　Method
	//ショートカットキー(Mac)：Command + Option + C 　定数

	//■ 最近開いたファイルを表示(Recent Files)
	//ショートカットキー(Mac)：Command + E

	//■ 選択箇所をif、while、forなどで囲む(Surround With)
	//ショートカットキー(Mac)：Command + Option + T

	//■ 現在の行を複製(Duplicate Line or Selection)
	//ショートカットキー(Mac)：Command + D

	//■ 新しい行を追加(Start New Line)
	//ショートカットキー：Shift + Enter

	//■ 宣言へジャンプ(Go to Declaration or Usages)
	//ショートカットキー(Mac)：Command + B

	//■ 実装クラスへジャンプ(Go to Implementations)
	//ショートカットキー(Mac)：Command + Option + B

	//■ 型の継承階層を表示する(Type Hierarchy)
	//ショートカットキー：Ctrl + H


	//■ 変数が使われている箇所へジャンプ(Show Usages)
	//ショートカットキー(Mac)：Command + Option + F7

	//■ 呼び出し先を調べる(Call Hierarchy)
	//ショートカットキー(Mac)：Option + Ctrl + H


	//■ シグネチャの変更(Change Signature)
	//ショートカットキー(Mac)：Command (+ fn) + F6


	//■ テストクラスと実装クラスの相互ジャンプ(Go to Test)
	//ショートカットキー(Mac)：Command + Shift + T

	//
	//■ 折り畳み(Collapse)と展開 (Expand)
	//ショートカットキー(Mac)：
	//
	//折り畳み(Collapse): Command + - , Command + Shift + -



}

func getUrl(url string) string {
	if req, err := http.NewRequest("GET",
		"http://www.google.com", nil); err != nil {
		return ""
	} else {
		res, _ := http.DefaultClient.Do(req)
		bs, _ := ioutil.ReadAll(res.Body)

		return string(bs)

	}

}
