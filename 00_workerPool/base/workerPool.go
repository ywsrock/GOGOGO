package main

import (
	"fmt"
	"time"
)

func main() {

	// 起動するWokerの数
	var wokerNum int = 5
	var numJobs int = 3

	//複数ルーチンの場合、終了まちの場合は　WatiGroupを利用する
	//var wg sync.WaitGroup

	jobs := make(chan int, numJobs)
	res := make(chan int, numJobs)

	for i := 1; i <= wokerNum; i++ {
		//Woker起動、Jobがないので、処理しない　Block状態
		//wg.Add(1)
		go woker(i, jobs, res)
		//go woker(i, jobs, res,&wg)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	//JobsにJob設定完了、Close
	close(jobs)

	//処理結果出力
	for a := 1; a <= numJobs; a++ {
		<-res
	}

	//wg.Wait()
}

func woker(id int, jobs <-chan int, result chan<- int) {
	//func woker(id int, jobs <-chan int, result chan<- int,wg *sync.WaitGroup) {
	//defer wg.Done()
	for job := range jobs {
		fmt.Println("start worker->", id)
		time.Sleep(2 * time.Second)
		fmt.Println("end worker->", job)
		result <- job * 2
	}
}
