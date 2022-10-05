package worker

import (
	"00_goCollyMinikabu/process"
	"sync"
)

// Dispatcher作成
type dispatcher struct {
	// WorkPool
	workerPool chan *Worker
	// Message　queue
	messageQueue chan interface{}
	// Workers
	Workers []*Worker
	// Quit
	Quit chan struct{}
	// wait
	wg *sync.WaitGroup
}

// Worker
type Worker struct {
	// dp
	dp *dispatcher
	// job
	job chan interface{}
	// Quit
	quit chan struct{}
}

// Workerの初期化、タスクの受付、タスクの配布
func (d *dispatcher) Start() {
	for _, worker := range d.Workers {
		worker.start()
	}

	go func() {
		for {
			select {
			case m := <-d.messageQueue:
				(<-d.workerPool).job <- m
			case <-d.Quit:
				return
			}
		}
	}()
}

// タスクの受付、タスク処理
func (w *Worker) start() {
	go func() {
		for {
			w.dp.workerPool <- w
			select {
			case job := <-w.job:
				if task, ok := job.(string); ok {
					// タスク処理
					process.ProcessTask(task)
				}
				// カウントダウン
				w.dp.wg.Done()
			case <-w.quit:
				return
			}
		}
	}()
}

func NewWorker(poolSize int, taskSize int) *dispatcher {
	d := &dispatcher{
		workerPool:   make(chan *Worker, poolSize),
		messageQueue: make(chan interface{}, taskSize),
		Quit:         make(chan struct{}),
		wg:           &sync.WaitGroup{},
	}
	// worker
	for i := 0; i < cap(d.workerPool); i++ {
		w := Worker{
			dp:   d,
			job:  make(chan interface{}),
			quit: make(chan struct{}),
		}
		d.Workers = append(d.Workers, &w)
	}

	return d
}

func (d *dispatcher) Add(tasks []string) {
	for _, task := range tasks {
		d.wg.Add(1)
		d.messageQueue <- task
	}
}

func (d *dispatcher) AddOne(task string) {
	d.wg.Add(1)
	d.messageQueue <- task
}

func (d *dispatcher) Wait() {
	d.wg.Wait()
}
