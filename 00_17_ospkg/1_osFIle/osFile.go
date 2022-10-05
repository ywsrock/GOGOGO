package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	// ファイル新規作成
	f, err := os.Create("new.txt")
	if err != nil {
		// 文字書き込む
		log.Fatal("ファイル作成失敗")
	}

	// ファイル情報取得
	fInfo, _ := f.Stat()
	fmt.Println(fInfo.Name())
	fmt.Println(fInfo.Size())

	// データの書き込む
	f.Write([]byte("Hello World!\n"))
	f.WriteString("string-------------")
	f.WriteAt([]byte("Golang"), 7)

	// ファイルを開く
	of, _ := os.Open("new.txt")
	// ファイル読み込
	bs := make([]byte, 128)
	_, err1 := of.Read(bs)

	if err1 != nil {
		log.Fatal("read error")
	}
	fmt.Println(string(bs))

	// File
	defer f.Close()
	defer of.Close()
}
