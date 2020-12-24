package main

import (
	"fmt"
)

func recive(ch <-chan int) {
	for {
		//受信
		i := <-ch
		fmt.Println(i)
	}
}

func main() {
	fmt.Println("チャンネル")

	//チャンネル作成
	ch := make(chan int)
	go recive(ch)
	i := 0
	for i < 10 {
		//チャンネルに送信
		ch <- i
		i++
	}

	//buffer 指定
	ch1 := make(chan int, 3)
	go func() {
		//送信
		ch1 <- 123
		ch1 <- 223
		ch1 <- 323
	}()
	// 受信
	ii := <-ch1
	fmt.Println(ii)
}
