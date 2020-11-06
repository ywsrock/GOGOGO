package main

import (
	"fmt"
)
func main()  {
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)
	ch3 := make(chan int, 1)

	// ch1 <- 1
	// ch2 <- 2

	// select {
	// case <- ch1:
	//    fmt.Println("ch1 rec")
	// case <- ch2:
	//    fmt.Println("ch2 rec")
	// case ch3 <- 3:
	//    fmt.Println("ch3 send")
	// default:
	// 	fmt.Println("222")
	// }
	

	go func() {
		for { 
			i := <- ch1
			ch2 <- (i * 2) 
			} 
		}()

	go func() {
		for { 
			i := <- ch2
			ch3 <- (i - 1) 
			} 
		}()
		
	n:=1
	LOOP:
	for{
		select {
		case ch1 <- n:
				n++
		case  i:= <- ch3:
				fmt.Println(i)
		default:
			if n >= 5 {
				break LOOP
			}
		}
	}
}