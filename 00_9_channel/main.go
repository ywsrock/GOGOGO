package main

import (
	"fmt"
	"time"
)

// ルーチン作成
func receive(ch <-chan int, name string) {
	for {
		i, ok := <-ch
		if ok == false {
			break

		}
		fmt.Println(name, i)
	}
	fmt.Println("done!")
}

func main() {

	ch := make(chan int, 20)

	go receive(ch, "routine01")
	go receive(ch, "routine02")
	go receive(ch, "routine03")

	i := 0
	for i <= 100 {
		ch <- i
		i++
	}
	close(ch)
	time.Sleep(3 * time.Second)
}
