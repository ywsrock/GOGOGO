package worker

import (
	task "minikabu/Task"
	"sync"
)

type Dispatcher struct {
	WorkPool chan *Worker
	Message  chan interface{}
	Workers  []*Worker
	Quit     chan struct{}
	Wg       *sync.WaitGroup
}

type Worker struct {
	dispatcher *Dispatcher
	job        chan interface{}
	quit       chan struct{}
}

func (d *Dispatcher) Start() {
	for _, w := range d.Workers {
		if w != nil {
			w.Start()
		}
	}
	go func() {
		for {
			select {
			case m := <-d.Message:
				(<-d.WorkPool).job <- m
			case <-d.Quit:
				return
			}
		}
	}()
}

func (w *Worker) Start() {
	go func() {
		w.dispatcher.WorkPool <- w
		for {
			select {
			case m := <-w.job:
				task.StartJob(m)
				w.dispatcher.Wg.Done()
			case <-w.quit:
				return
			}
		}
	}()
}

func NewWorker(poolSize int, taskSize int) *Dispatcher {
	d := &Dispatcher{
		WorkPool: make(chan *Worker, poolSize),
		Message:  make(chan interface{}),
		Quit:     make(chan struct{}),
		Wg:       &sync.WaitGroup{},
	}

	ws := make([]*Worker, taskSize)
	for j := 0; j < cap(d.WorkPool); j++ {
		w := &Worker{
			dispatcher: d,
			job:        make(chan interface{}),
			quit:       make(chan struct{}),
		}
		ws[j] = w
	}
	d.Workers = ws
	return d
}
