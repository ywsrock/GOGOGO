package main

import (
	"fmt"
)

// User 構造
type User struct {
	id   int
	name string
}

// NewUser 構造体コンストラクタ 普通はポインタを返す
func NewUser(id int, name string) *User {
	// new キーワードでポインタ作成 == &user
	u := new(User)
	u.id = id
	u.name = name
	return u
}

// 構造体メソッド
func (u *User) getId() int {
	return u.id
}

func (u *User) getName() string {
	return u.name
}

func main() {
	fmt.Println("-----構造体-----")
	// 構造体ポインタ作成
	newUser := NewUser(1, "struct")
	fmt.Println(newUser)
	// メソッド呼び出す
	uid := newUser.getId()
	uname := newUser.getName()
	fmt.Println(uid, uname)
}
