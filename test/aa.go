package main

import "fmt"

func main() {
	//ch := make(chan int,1)
	//for i :=0; i <= 10; i++ {
	//	select {
	//	case x:= <-ch :
	//		fmt.Println(x)
	//	case ch <- i:
	//		fmt.Println("test")
	//	}
	//}

	//clients := make(map[string]bool)
	//clients["a"] = false
	//clients["b"] = true
	//
	//for cli := range clients{
	//	fmt.Println(cli)
	//}

	chs := make(chan int)
	chs2 := make(chan chan <- int)

	go func(){
		chs <- 1
		chs2 <- chs
	}()

	fmt.Println ( <- chs)
	fmt.Println ( <- chs2)


}

