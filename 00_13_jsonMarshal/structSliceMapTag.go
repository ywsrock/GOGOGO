package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// User info
type User struct {
	ID   int    `json:"id"`
	Name string `json:"Name"`
}

func main() {
	fmt.Println("-------構造体------")
	// スライス作成
	user := make([]User, 3)
	user = append(user, User{1, "test"})
	fmt.Println(user)
	// 構造体ポインタ形式
	var pUser []*User
	pUser = append(pUser, &User{2, "test2"})
	pUser = append(pUser, &User{3, "test3"})
	pUser = append(pUser, &User{4, "test4"})
	pUser = append(pUser, &User{5, "test5"})
	pUser = append(pUser, &User{6, "test6"})
	fmt.Println(pUser[0].ID)
	for _, v := range pUser {
		fmt.Println(v.ID, v.Name)
	}

	// マップ型構造体
	mu := map[int]User{
		1: {1, "test"},
		2: {2, "test1"},
	}
	fmt.Println(mu[1])
	for k, v := range mu {
		fmt.Println(k, v.Name)
	}

	// タグの出力
	ut := User{1, "tag"}
	t := reflect.TypeOf(ut)
	fmt.Println("tag---->", t)
	for i := 0; i < t.NumField(); i++ {
		fmt.Println(t.Field(i))
	}
	// 構造体JSON出力
	uJson, _ := json.Marshal(ut)
	fmt.Println(string(uJson))
}
