package main

import (
	"fmt"
)

func main() {
	//整数
	i := 1
	fmt.Println(i)

	//倫理
	flg := false
	fmt.Println(flg)

	//浮動型
	f := 1.0
	fmt.Println(f)

	//文字列型
	str := "test"
	fmt.Println(str)

	//複素数型
	c := 1.0 + 3i
	fmt.Println(c)
	real := real(c) //実数取得
	imag := imag(c) //虚数取得
	fmt.Println(real)
	fmt.Println(imag)

	//rune型 文字対応のコード表http://www3.nit.ac.jp/~tamura/ex2/ascii.html
	var rune rune
	// rune = 'a' //97
	rune = '\n'
	fmt.Println(rune)

	// 配列　明示型
	var arr1 [3]int
	arr1[0] = 1
	arr1[1] = 2
	arr1[2] = 3
	// 配列　　暗黙型
	arr2 := [4]int{1, 2, 3, 4}

	fmt.Printf("%d\n", arr1[1])
	fmt.Printf("%d\n", arr2[2])

	//    同じ要素数、同じ型だけ、交換性ができます
	// ★配列代入の場合、配列コピーになります。メモリは別々で管理されます。
	a1 := [2]int{1, 2}
	b1 := [2]int{5, 6}
	a1 = b1
	fmt.Println(a1)
	fmt.Println(b1)

	//interface{} 型　初期化時はnil状態 nil　は　null の理解してもいい
	//interace{}　java のＯｂｊｅｃｔ見たい、何でも受ける型です
	//★演算はできません
	 var inf interface{}
	 inf = 1
	 fmt.Printf("%v\n",inf)
	 if v,err := inf.(int);err{
		fmt.Printf("型%v",v)
	 }
	 inf = [6]int{1,2,3,4,}
	 fmt.Printf("%v\n",inf)
}
