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

type If interface {
	getname() string
}

type ImpIf struct {
	name string
}

func (I *ImpIf) getname() string{
	return I.name
}



//interface の実装とき、interfaceで定義されたメソッド名を同じにすると
//実装を見なす
func main() {
	err := RaiseError()
	fmt.Println(err.Error())

	ii := &ImpIf{name: "12345"}
	fmt.Println(ii.getname())

}
