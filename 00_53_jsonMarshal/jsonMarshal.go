package main

import (
	"encoding/json"
	"fmt"
	"log"
)
type User struct {
	Id   int
	Name string
}

func main() {
  // u := new(User)
  // ポインタ作成
    u := &User {
    Id : 1,
    Name : "test",
  }

  // ＪＳＯＮ変換
  bs, err := json.Marshal(u)
  if err != nil {
    log.Fatal(err)
  }
  fmt.Println(string(bs))
  
  var p *User
  
  er1 := json.Unmarshal(bs,p)
  if er1 != nil {
    log.Fatal(er1)
  }

  fmt.Println(p)


}