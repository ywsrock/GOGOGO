package main

import (
	"fmt"
)

func main() {
	// 普通IF
	if 1 < 2 {
		fmt.Println("1<2")
	}

	i := 1
	j := 2

	// if Else if　形式
	if i == j {
		fmt.Println("i == j")

	} else if i < j {
		fmt.Println("i < j")
	} else {
		fmt.Println("i > j")
	}

	// 簡易文付き
	if i, j := 3, 4; i < j {
		fmt.Println("簡易文付き")
	}

	// For
	arr := [3]int{1, 2, 3}
	count := len(arr)
	for i := 0; i < count; i++ {
		fmt.Println(arr[i])
	}

	// 無限的Loop break Loop終了
	var a int
	for {
		fmt.Printf("current i = %v \n", a)
		if a == 3 {
			break
		}
		a += 1
	}

	// 10以内の奇数出力
	for n := 0; n < 10; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println("現在の値--->", n)
		n++
	}

	// 条件指定ループ
	var b int
	for b < 10 {
		fmt.Printf("条件指定 i < 10 %d \n", b)
		b++
	}

	// 範囲指定
	var slic []int = []int{1, 2, 3, 4, 5}
	slic[0] = 0
	slic[1] = 1
	slic[2] = 2

	for i := range slic {
		fmt.Println(i)
	}

	// switch 値の指定（複数も可能）
	s := 1
	switch s {
	case 2, 3, 4:
		fmt.Println("2,3,4")
	case 5, 6, 7:
		fmt.Println("5,6,7")
	default:
		fmt.Println("default")
	}

	// switch 範囲の指定 式の指定
	switch {
	case 1 <= s && s < 5:
		fmt.Println("1<s<5")
	case 5 <= s && s < 8:
		fmt.Println("5 < s < 8")
	default:
		fmt.Println("範囲の指定初期値")
	}

	// 簡易文付き
	switch x := 2; x {
	case 1, 2, 3:
		fmt.Println("x = 2")
	default:
		fmt.Println("簡易文付きdefault")
	}

	// 型のアサーション
	var typeAs interface{}
	typeAs = 1

	fmt.Println(typeAs)
	fmt.Println(typeAs.(int))

	switch typeAs.(type) {
	case string:
		fmt.Println("string")
	case int, uint:
		fmt.Println("int uint")
	default:
		fmt.Println("unKnow")
	}

	// GOTO　文 label　は　関数内に閉じたものです。関数の末尾に記載する
	var ar []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, v := range ar {
		if v%2 == 0 {
			goto label
		}
		fmt.Println("GOTO 現在の処理=", v)
	}

label:
	fmt.Println("GOTO　文")
	// label付き文
Loop:
	for {
		for {
			fmt.Println("処理開始")
			break Loop
		}
		// 処理は通らない
		fmt.Println("通らない")
	}
}
