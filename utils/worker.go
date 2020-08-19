package utils

import (
	"sync"
)

type Work interface {
	Execute() error
}

type WorkerPool struct {
	MaxWorker int
	wg        sync.WaitGroup
	workChan  chan Work
	done      chan bool
}

func NewWorkerPool(maxWorker int) *WorkerPool {
	return &WorkerPool{
		MaxWorker: maxWorker,
		wg:        sync.WaitGroup{},
		done:      make(chan bool, 0),
		workChan:  make(chan Work, 0),
	}
}
func (p *WorkerPool) Add(w Work) {
	p.wg.Add(1)
	p.workChan <- w
}

func (p *WorkerPool) SendDone() {
	p.wg.Done()
}

func (p *WorkerPool) Stop() {
	p.done <- false
}

// RunWithBlock need call with SendDone
func (p *WorkerPool) RunWithBlock() {
	p.wg.Add(1)
	p.run()
	go func() {
		p.wg.Wait()
		p.done <- true
	}()
	<-p.done
	close(p.workChan)
	close(p.done)
}

func (p *WorkerPool) run() {
	for i := 0; i < p.MaxWorker; i++ {
		go func(i int) {
			for w := range p.workChan {
				_ = w.Execute()
				p.wg.Done()
			}
		}(i)
	}
}
