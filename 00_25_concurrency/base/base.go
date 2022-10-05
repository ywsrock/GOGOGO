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

func doTask() {
	numbers := []int{1, 2, 3, 4, 5, 6}
	// goRoutine完了まで待ち
	var wg sync.WaitGroup
	for _, num := range numbers {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			fnAction(n)
			// 時間かかる処理
			time.Sleep(1 * time.Second)
		}(num)
	}
	// 完了待ち
	wg.Wait()
}

func fnAction(n int) {
	log.Printf("do fnAction. num: %d \n", n)
}
