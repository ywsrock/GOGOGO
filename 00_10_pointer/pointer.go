package main

import (
	"fmt"
	f "fmt"
)

func main() {
	// 配列の作成
	i := [3]int{1, 2, 3}

	// ポインタ取得
	p := &i[1]
	// 値の変更
	f.Printf("Type=%T\n", p)
	//値の変更理
	i[1] = 4
	f.Printf("i[1]=%d\n", i[1])

	// アドレス経由で、値の設定
	*p = 6
	f.Printf("i[1]=%d\n", i[1])

	//new キーワードで、任意型のポインタを作成することができます
	var ni int
	fmt.Printf("ni =%T\n", ni)
	type nip int
	pni := new(nip)
	fmt.Printf("pni=%T\n", pni)

	// 構造体
	type myUser struct {
		id   int
		name string
	}
	user := myUser{1, "test"}
	// 基本型出力
	fmt.Println(user.id)
	user.id = 2
	fmt.Println(user.id)
	//ポインタで変更
	puser := &user
	fmt.Println("userID=", puser.id)
	// ポインタ経由で変更
	puser.id = 3
	fmt.Println("userid = ", puser.id)

}
