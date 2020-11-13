package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//コマンドラインから値を受け取る
	scanner := bufio.NewScanner(os.Stdin)

	//入力を受け取る
	for scanner.Scan() {
		//受けた値の出力
		fmt.Println("受けた値:", scanner.Text())
	}

	//エラーの場合
	if err := scanner.Err(); err != nil {
		fmt.Println("処理エラー")
	}
}
