package main

import (
	"00_goCollyMinikabu/worker"
	"fmt"
	"os"
)

const PoolSize = 10
const MaxTask = 20

func main() {
	quit := make(chan os.Signal)
	d := worker.NewWorker(PoolSize, MaxTask)
	// Worker dispatcher start
	d.Start()
	// タスク作成
	tasks := []string{"1570", "1571"}
	d.Add(tasks)

	go func() {
		var key string
		for {
			fmt.Scanln(&key)
			if key == "exit" {
				quit <- os.Interrupt
				break
			}
			d.AddOne(key)
		}
	}()

	// 終了待ち
	d.Wait()
	<-quit
}
