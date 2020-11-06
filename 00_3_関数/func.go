package main

import (
	"fmt"
)

//関数引数,戻り値
func sum(p1, p2 int) int {
	return p1 + p2
}

// 複数引数、複数戻り値
func sub(p1 int, p2 int) (string, int) {
	return "計算結果=", p1 - p2
}

//エラー処理
func div(a int, b int) (int, string) {
	if b == 0 {
		return 0, "error"
	}
	return a / b, ""
}

//戻り値は変数名指定
func getValue() (name string) {
	name = "Function"
	return 	
}

// 上記と同じの意味
func getValue1() string {
	var name string
	name = "上記と同じ"
	return name
}

//関数返す関数
func retFunction() func(int,int) int {
	return func(x,y int) int {
		return x +y
	}
}

//閉包処理 GO コンパイラがクロージャ内からのローカル変数参照を検出する場合、
// クロージャと変数結び付けいた変数として処理します。
func closure() func(string) string{
	var str string
	return func (contactStr string) string {
		//ローカル変数strに参照すると閉包成立
		str += contactStr
		return str
	}
}

func main() {
	sum := sum(3, 4)
	fmt.Println(sum)

	str, v := sub(6, 2)
	fmt.Println(str, v)

	// 戻り値廃棄の場合
	_, v1 := sub(3, 2)
	fmt.Println(v1)

	//エラーが発生処理方法の一つ
	d, err := div(1, 0)
	fmt.Println(err,d)

	d1, err1 := div(1, 1)
	fmt.Println(err1,d1)

	//戻り値は変数名指定
	geV := getValue()
	fmt.Println(geV)
	getV1 := getValue1()
	fmt.Println(getV1)

	//無名関数
	f := func(x, y int) int {return x +y}
	fmt.Println(f(3,4))

	// 上記の同じ書き方
	var f1 func(int,int) int
	f1 = func(x,y int) int {return x + y}
	fmt.Println(f1(6,1))

	// 関数を返す関数
    retFunc1 := retFunction()
	reNum := retFunc1(1,2)
	fmt.Println(reNum)
	retNum1 := retFunction()(2,3)
	fmt.Println(retNum1)

	//クロージャ 文字列連結して出力
	closure1 := closure()
	var strclosure string
	strclosure = closure1("aa")
	strclosure = closure1("bb")
	strclosure = closure1("cc")
	fmt.Println(strclosure)
}