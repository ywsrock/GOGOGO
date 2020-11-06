package main

import "fmt"

//エラーインターフェース
type error interface {
	Error() string
}

//MyError エラー情報表示構造体
type MyError struct {
	ErrMsg  string
	ErrCode int
}

//インターフェースのメソッド実現
func (m *MyError) Error() string {
	return m.ErrMsg
}

//RaiseError エラー作成関数 ★戻り値の型はインターフェース型、戻り理は実現構造
func RaiseError() error {
	return &MyError{ErrMsg: "Error implement", ErrCode: 10}
}

//interface の実装とき、interfaceで定義されたメソッド名を同じにすると
//実装を見なす
func main() {
	err := RaiseError()
	fmt.Println(err.Error())
}
