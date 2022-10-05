package main

import (
	"fmt"
)

func main() {
	// 変数種類:値型、参照型変数(スライス、マップ、チャンネル、ポインタ)
	// 変数の定義：明示的、暗黙的
	// 明示的な定義
	var i int
	i = 1
	fmt.Println(i)

	var str string
	str = "GOLANG"
	fmt.Println(str)

	var isTrue bool
	isTrue = false
	fmt.Println(isTrue)

	var v1 int = 11
	fmt.Print(v1)

	// 暗黙的な定義,型を自動推論
	f := 1
	fmt.Println(f)
	b := true
	fmt.Printf("%T", b)

	// ブロック単位で定義

	// 同じ型な場合
	var a1, b1, c1 int
	a1, b1, c1 = 1, 2, 3
	fmt.Printf("%v %v %v", a1, b1, c1)

	// 違い型の場合
	var (
		e int
		g string
		t float64
	)
	e = 1
	g = "block"
	t = 0.1
	fmt.Printf("%v %v %v", e, g, t)

	var (
		e1 = 1
		e2 = 2
		e3 = 3
	)
	fmt.Println(e1, e2, e3)

	// 変数の代入（同じ型しかできない。暗黙的な型変換一切しません）
	var kk1 int = 1
	var kk2 int
	kk2 = kk1
	fmt.Println(kk2)
}
