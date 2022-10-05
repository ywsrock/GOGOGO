package main

import (
	"fmt"
	"log"
	"os"
)

func pDefer() {
	fmt.Println("Exit指定の場合、defer 実行しません")
}

func main() {
	defer pDefer()
	// コマンドライン引数
	if len(os.Args) > 0 {
		for _, v := range os.Args {
			fmt.Println("引数--->", v)
		}
	}

	// Log Fatal ログ出力　+ 終了コード１
	_, fErr := os.Open("123")
	if fErr != nil {
		log.Fatal("ファイル存在しません")
	}

	// 終了コード指定する場合、deferは実行しません。
	fmt.Println("終了コード")

	// 終了コード指定
	// os.Exit(1)
}
