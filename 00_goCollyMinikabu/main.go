package main

import (
	"00_goCollyMinikabu/worker"
)

const PoolSize = 10
const MaxTask = 20

func main() {
	d := worker.NewWorker(PoolSize, MaxTask)
	//Worker dispatcher start
	d.Start()
	//タスク作成
	tasks := []string{"1570","1571"}
	d.Add(tasks)
	//終了待ち
	d.Wait()
}
