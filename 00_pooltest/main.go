package main

import (
	"fmt"
	"sync"
	"time"
)

type Dispatcher struct {
	Pool    chan *Worker
	Worker  []*Worker
	Message chan interface{}
	Quit    chan struct{}
	wg      *sync.WaitGroup
}

type Worker struct {
	d       *Dispatcher
	message chan interface{}
	quit    chan struct{}
}

func (d *Dispatcher) Start() {
	for _, w := range d.Worker {
		w.Start()
	}
	go func() {
		for {
			select {
			case m := <-d.Message:
				(<-d.Pool).message <- m
			case <-d.Quit:
				return
			}
		}
	}()
}

func (w *Worker) Start() {
	go func() {
		w.d.Pool <- w
		for {
			select {
			case m := <-w.message:
				Run(m)
				w.d.wg.Done()
			case <-w.quit:
				return
			}
		}
	}()
}

func Run(s interface{}) {
	time.Sleep(3)
	fmt.Println(s)
}

func main() {
	poolSize := 10
	workSize := 20

	d := Dispatcher{
		Pool:    make(chan *Worker, poolSize),
		Message: make(chan interface{}),
		Quit:    make(chan struct{}),
		wg:      &sync.WaitGroup{},
	}

	w := make([]*Worker, workSize)
	for j := 0; j < cap(w); j++ {
		wk := Worker{
			d:       &d,
			message: make(chan interface{}),
			quit:    make(chan struct{}),
		}
		w[j] = &wk
	}
	d.Worker = w

	d.Start()

	ar := [...]string{"a", "b", "c"}
	for _, s := range ar {
		d.Message <- s
		d.wg.Add(1)
	}

	d.wg.Wait()
}
