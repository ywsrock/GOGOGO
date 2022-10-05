package main

import "fmt"

// 新しいタイプ定義
type Callback func(i int) int

// 関数を計算
func sumIntFunc(ins []int, callBack Callback) int {
	var sum int
	// 累積加算
	for _, i := range ins {
		sum += i
	}

	return callBack(sum)
}

// コールバック関数定義
func multi(num int) int {

	return num * 2
}

func main() {
	// スライス作成
	var is []int = []int{1, 2, 3, 4, 5}
	var ret = sumIntFunc(is, multi)
	fmt.Println(ret)
}
