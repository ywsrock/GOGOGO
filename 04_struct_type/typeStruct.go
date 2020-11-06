package main

import "fmt"

// 新しいタイプ定義
type Callback func(i int) int

// 関数を計算
func sumInts(ints []int,callBack Callback) int {
	var sum int
	// 累積加算
	for _,i := range ints {
		sum += i
	}

	return callBack(sum)
}

// コールバック関数定義
func kake2(num int) int {
	
	return num * 2
}

func main() {
	// スライス作成
	var ints []int= []int{1, 2, 3, 4, 5,}
	var ret = sumInts(ints,kake2)
	fmt.Println(ret)
}