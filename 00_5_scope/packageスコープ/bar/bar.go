package bar

import (
	f "fmt"
)
//頭大文字 他のパッケージから参照できる
func DoTest() string {
   return "Foo から参照"
}

// 定数定義
const (
	m = 1 
	N = 2 // 他のパッケージから参照できる
)

func main() {
	str1 := DoTest()
	f.Println(str1)
	f.Println(m,N)
}