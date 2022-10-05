package main

import (
	"log"
	"sync"
	"time"
)

func main() {
	doTask()
	log.Println("finished")
}

const concurrency = 2 // 最大同時並列実行数

func doTask() {
	numbers := []int{1, 2, 3, 4, 5, 6}

	// 完了するまで、待機
	var wg sync.WaitGroup
	sem := make(chan struct{}, concurrency) // concurrency数のバッファ

	for _, num := range numbers {
		// chan　に送信
		sem <- struct{}{}

		// 終了待ち数
		wg.Add(1)
		// ルーチン実行
		go func(n int) {
			defer wg.Done()
			// 処理が終わったらチャネルを解放
			defer func() { <-sem }()

			fnAction(n)

			// 処理をわかりやすくするため
			time.Sleep(1 * time.Second)
		}(num)
	}
	wg.Wait()
}

func fnAction(n int) {
	log.Printf("do fnAction. num: %d \n", n)
}
