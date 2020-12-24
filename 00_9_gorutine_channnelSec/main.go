package main

import (
	"fmt"
	"time"
)

//ルーチン作成
func recive(ch <-chan int, name string) {
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

	go recive(ch, "rutine01")
	go recive(ch, "rutine02")
	go recive(ch, "rutine03")

	i := 0
	for i <= 100 {
		ch <- i
		i++
	}
	close(ch)
	time.Sleep(3 * time.Second)
}
