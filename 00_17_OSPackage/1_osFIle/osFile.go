package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	//ファイル新規作成
	f, err := os.Create("new.txt")
	if err != nil {
		//文字書き込む
		log.Fatal("ファイル作成失敗")
	}

	//ファイル情報取得
	fInfo, _ := f.Stat()
	fmt.Println(fInfo.Name())
	fmt.Println(fInfo.Size())

	//データの書き込む
	f.Write([]byte("Hello World!\n"))
	f.WriteString("string-------------")
	f.WriteAt([]byte("GOlang"), 7)

	//ファイル読み込
	bs := make([]byte, 128)
	_, err1 := f.Read(bs)

	if err1 != nil {
		fmt.Println(string(bs))
		for _, b := range bs {
			fmt.Println(b)
		}
	}
}
