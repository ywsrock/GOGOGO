package main

import (
	"fmt"
	"reflect"
	"sync"
	"time"
)

const PoolSize = 10
const MaxJobs = 20

// task
func PrintString(i string) {
	time.Sleep(time.Second * 1)
	fmt.Println("実行-->", i)
}

// task mmanager
type Dispatcher struct {
	WorkerPool chan *Worker
	Workers    []*Worker
	Jobs       chan interface{}
	quit       chan struct{}
	wg         *sync.WaitGroup
}

// Worker
type Worker struct {
	dp   *Dispatcher
	job  chan interface{}
	quit chan struct{}
}

// Manager Start
func (d *Dispatcher) start() {
	// worker 初期化
	for _, worker := range d.Workers {
		// worker 起動
		worker.start()
	}

	go func() {
		for {
			select {
			case job := <-d.Jobs:
				t := reflect.TypeOf(job)
				fmt.Println(t.String())
				if task, ok := job.(string); ok {
					(<-d.WorkerPool).job <- task
				}
			case <-d.quit:
				fmt.Println("--------------")
				return
			}
		}
	}()
}

// Worker start
func (w *Worker) start() {
	go func() {
		for {
			w.dp.WorkerPool <- w
			select {
			case m := <-w.job:
				if str, ok := m.(string); ok {
					PrintString(str)
				}
				w.dp.wg.Done()
			case <-w.quit:
				fmt.Println("--------------")
				return
			}
		}
	}()

}

func main() {
	d := new(Dispatcher)
	d.WorkerPool = make(chan *Worker, PoolSize)
	d.Jobs = make(chan interface{}, MaxJobs)
	d.quit = make(chan struct{})
	d.wg = &sync.WaitGroup{}

	for i := 0; i < cap(d.WorkerPool); i++ {
		w := Worker{
			dp:   d,
			job:  make(chan interface{}),
			quit: make(chan struct{}),
		}
		d.Workers = append(d.Workers, &w)
	}

	d.start()

	tasks := make([]string, 100)
	for i := range tasks {
		tasks[i] = fmt.Sprintf("Hello---->%d", i)
		d.wg.Add(1)
		d.Jobs <- tasks[i]

	}
	d.wg.Wait()
}
