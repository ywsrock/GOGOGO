package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}

	// ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Microsecond)
	// ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second))
	defer cancel()

	wg.Add(1)
	go child(ctx, &wg)

	wg.Wait()
}

func child(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("child timeout")
			return
		default:
			fmt.Println("child is running")
		}
	}
}
