package main

import (
	"fmt"
)

// メイン実行前に実行（複数存在可）
func init() {
	fmt.Println("Main 開始前に実行")
}

// defer プログラム終了後に実行
func fileClose() {
	fmt.Println("ファイルのクローズ処理")
}

// サブルーチン
func sub() {
	for i := 0; i < 10; i++ {
		fmt.Println("Sub......")
	}
}

func main() {
	fmt.Println("--------Main開始-----")

	// GO routine
	go sub()
	for i := 0; i < 10; i++ {
		fmt.Println("Main......")
	}
	// 上記と同じ
	go func() {
		fmt.Println("sub main().......")
	}()
	// 関数終了時実行される(最後の定義されたdeferは最優先実行)
	defer fmt.Println("defer 関数終了しました.1")
	defer fmt.Println("defer 関数終了しました.2")
	defer fmt.Println("defer 関数終了しました.3")
	defer fmt.Println("defer 関数終了しました.4")

	// 無名関数
	defer func() {
		fmt.Println("defer 無名関数")
	}()

	// 関数実行
	defer fileClose()

	for _, v := range [3]int{1, 2, 3} {
		fmt.Println(v)
	}

	// panic　ランダムを強制的中止(例外throw)
	fmt.Println("実行開始")
	// recover とpanic 組み合わせは原則です
	defer func() {
		if x := recover(); x != nil {
			fmt.Println("-------------", x)
		}
	}()
	panic("強制中止")
	fmt.Println("panic 通らない")
}
