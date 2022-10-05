package main

import (
	"fmt"
	"sync"
)

// var wg sync.WaitGroup

func child(done chan struct{}, num int, wg *sync.WaitGroup) chan int {
	out := make(chan int)

	go func() {
		defer wg.Done()
	Loop:
		for {
			select {
			case <-done:
				break Loop
			case out <- num:
			}
		}
		close(out)
		fmt.Println("child close")
	}()
	return out
}

func main() {
	done := make(chan struct{})
	wg := sync.WaitGroup{}

	out := child(done, 1, &wg)

	wg.Add(1)
	for i := 0; i < 5; i++ {
		fmt.Println(<-out)
	}
	close(done)
	wg.Wait()
}
