package main

// import は別々で管理
import (
	"fmt"
)

// メソッドは別ファイルで定義
// go run main.go scope.go(分割したファイルを全部指定必要)
func main() {
	str1 := test()
	fmt.Println(str1)

	str2 := doTest()
	fmt.Println(str2)
}
