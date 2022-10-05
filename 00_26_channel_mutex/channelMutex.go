package main

import (
	"fmt"
	// "time"
	"sync"
)

var wg *sync.WaitGroup

func subGo(num int, c chan int) {

	// time.Sleep(time.Second * 1)
	fmt.Println("subGO", num)
	c <- num
	// wg.Done()
}
func main() {
	// wg = new(sync.WaitGroup)
	// wg.Add(4)
	// channel make
	ch := make(chan int, 3)
	mutex := new(sync.Mutex)
	// for i:=0;i<3;i++ {
	go subGo(2, ch)

	mutex.Lock()
	go subGo(3, ch)
	mutex.Unlock()
	go subGo(4, ch)
	go subGo(5, ch)
	re1, re2, re3, re4 := <-ch, <-ch, <-ch, <-ch
	fmt.Println("開始")
	fmt.Println(re1, re2, re3, re4)
	//   }
	// wg.Wait()

}
