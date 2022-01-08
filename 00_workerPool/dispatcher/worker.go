package main

import (
	"fmt"
	"hash/fnv"
	"math/rand"
	"sync"
	"time"
)

/**
dispatcher : ワーカーのインスタンス作成、ワーカーとワーカープールを接続する
workers : 処理待っているJobを受け取り、処理する
collector : Job到着を待って、Workerに割り当てます
*/

var letterRunes = []rune("abcdefghigklmnopqrstuvwxyzABCDEFGHIGKLMNOPQRSTUVWXYZ")

//create random string
func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

//create list of jobs
func CreateJobs(amount int) []string {
	var jobs []string
	for i := 0; i < amount; i++ {
		jobs = append(jobs, RandStringRunes(8))
	}
	return jobs
}

func doWork(word string) {
	h := fnv.New32()
	h.Write([]byte(word))
	time.Sleep(1 * time.Second)
	fmt.Printf("created hash [%d] from word [%s]\n", h.Sum32(), word)
}

type dispatcher struct {
	pool    chan *worker     // Worker Pool --実行可能なworkerしか持っている
	queue   chan interface{} //dispatcher 受信したメッセンジ
	quit    chan struct{}    //終了条件
	workers []*worker        //起動するWorker 起動されたWorker全部持っている
	wg      sync.WaitGroup
}

type worker struct {
	dispatcher *dispatcher
	data       chan interface{} // Worker　受信したメッセージ
	quit       chan struct{}    //　Worker 処理終了指示
}

/**
dispatcher 起動を経由で、Workerを起動
*/
func (d *dispatcher) start() {
	for _, wk := range d.workers {
		//Dispatch 所有しているWorkerを起動 worker インスタンス作成
		wk.start()
	}
	go func() {
		for {
			select {
			case v := <-d.queue:
				//worker Pool 中から、worker を取得し、メッセージをWorkerに渡す
				(<-d.pool).data <- v
			case <-d.quit:
				return
			}
		}
	}()
}

func (w *worker) start() {
	//起動したWorkerをPoolに設定する
	go func() {
		for {
			w.dispatcher.pool <- w
			select {
			case v := <-w.data:
				if str, ok := v.(string); ok {
					doWork(str)
				}
				//WaitGroup Countdown
				w.dispatcher.wg.Done()
			case <-w.quit:
				return
			}
		}
	}()
}

// Dispathcer に　メッセージを追加
func (d *dispatcher) Add(word string) {
	d.wg.Add(1)
	d.queue <- word
}

func (d *dispatcher) wait() {
	d.wg.Wait()
}

func main() {
	var poolMax = 5
	var jobMax = 10

	//Dispatcher 新規作成
	d := &dispatcher{
		pool:  make(chan *worker, poolMax),
		queue: make(chan interface{}, jobMax),
		quit:  make(chan struct{}),
	}

	// Worker の数
	d.workers = make([]*worker, cap(d.pool))
	for i := 0; i < cap(d.pool); i++ {
		wk := worker{
			dispatcher: d,
			data:       make(chan interface{}),
			quit:       make(chan struct{}),
		}
		d.workers[i] = &wk
	}

	d.start()

	jobs := CreateJobs(100)
	for _, job := range jobs {
		//fmt.Println(job)
		d.Add(job)
	}
	d.wait()
}
