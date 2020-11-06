package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

//User info
type User struct {
	ID   int    "json:id"
	Name string "json:Name"
}

func main() {
	fmt.Println("-------構造体------")
	//スライス作成
	user := make([]User, 3)
	user = append(user, User{1, "test"})
	fmt.Println(user)
	// 構造体ポインタ形式
	var puser []*User
	puser = append(puser, &User{2, "test2"})
	puser = append(puser, &User{3, "test3"})
	puser = append(puser, &User{4, "test4"})
	puser = append(puser, &User{5, "test5"})
	puser = append(puser, &User{6, "test6"})
	fmt.Println(puser[0].ID)
	for _, v := range puser {
		fmt.Println(v.ID, v.Name)
	}

	//マップ型構造体
	mu := map[int]User{
		1: {1, "test"},
		2: {2, "test1"},
	}
	fmt.Println(mu[1])
	for k, v := range mu {
		fmt.Println(k, v.Name)
	}

	//タグの出力
	ut := User{1, "tag"}
	t := reflect.TypeOf(ut)
	fmt.Println("tag---->", t)
	for i := 0; i < t.NumField(); i++ {
		fmt.Println(t.Field(i))
	}
	//構造体JSON出力
	ujson, _ := json.Marshal(ut)
	fmt.Println(string(ujson))
}
