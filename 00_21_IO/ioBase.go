package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("IO　BABASE")
	_, err := os.Stat("test.txt")
	var curf *os.File
	if os.IsNotExist(err) {
		//ファイル新規作成
		f, _ := os.Create("test.txt")
		curf = f
	} else {
		//存在の場合開く
		curf, _ = os.Open("test.txt")
	}

	curf.WriteString("123456")
	//ファイル開く
	// curFile, err1 := os.Open("test.txt")
	bs := make([]byte, 128)
	n, _ := curf.Read(bs)
	fmt.Println(n)
	fmt.Println(string(bs))

}
