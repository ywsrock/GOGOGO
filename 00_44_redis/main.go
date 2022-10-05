package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)
// Redis 接続
// Redisに値を書き込み
func main() {

	conn, err := redis.Dial("tcp", ":6379")
	if err != nil {
		panic(err)
	}
	response, err := conn.Do("AUTH", "ywsrock")
	if err != nil {
		panic(err)
	}
	fmt.Println("Connected! ", response)
	defer conn.Close()

	// 値の書き込み
	r, err := conn.Do("SET", "test", "10")
	if err != nil {
		panic(err)
	}
	fmt.Println(r) // OK

	// 値の読み出し
	s, err := redis.String(conn.Do("GET", "test"))
	if err != nil {
		panic(err)
	}
	fmt.Println(s) // 10

	// 値の削除
	// redisにデータが入っているか確認したい場合
	// 以下の行はコメントアウトしてください
	d, err := conn.Do("DEL", "test")
	if err != nil {
		panic(err)
	}
	fmt.Println(d) // 1
}
