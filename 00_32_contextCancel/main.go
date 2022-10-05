package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func child(ctx context.Context, num int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("child context cancel")
			return
		default:
			fmt.Println(num)
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	wg := sync.WaitGroup{}
	ctx, cancel := context.WithCancel(context.Background())
	t := time.After(6 * time.Second)

	wg.Add(2)

	go child(ctx, 1, &wg)

	go func() {
		defer wg.Done()
		for {
			select {
			case <-t:
				cancel()
				fmt.Println("cancel!")
				return
			default:
			}
		}
	}()

	wg.Wait()
}
