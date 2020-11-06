package main

// import は別々で管理
import (
	"fmt"
)

// メソッドは別ファイルで定義
func main() {
	str1 := test()
	fmt.Println(str1)

	str2 := doTest()
	fmt.Println(str2)
}
