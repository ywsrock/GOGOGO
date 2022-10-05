package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("io base")
	// file 情報
	_, err := os.Stat("test.txt")
	var curF *os.File
	// 存在チェック
	if os.IsNotExist(err) {
		// ファイル新規作成
		f, _ := os.Create("test.txt")
		curF = f
	} else {
		// 存在の場合開く
		curF, _ = os.Open("test.txt")
	}
	defer curF.Close()

	curF.WriteString("123456")
	// ファイル開く
	// curFile, err1 := os.Open("test.txt")
	bs := make([]byte, 128)
	n, _ := curF.Read(bs)
	fmt.Println(n)
	fmt.Println(string(bs))

}
