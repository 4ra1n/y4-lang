package pool

import "sync"

type Job func()

type Pool struct {
	jobs chan Job
	stop chan struct{}
	wg   sync.WaitGroup
}

func NewPool(size int) *Pool {
	p := &Pool{
		jobs: make(chan Job),
		stop: make(chan struct{}),
	}

	p.wg.Add(size)
	for i := 0; i < size; i++ {
		go p.worker()
	}

	return p
}

func (p *Pool) worker() {
	defer p.wg.Done()

	for {
		select {
		case job := <-p.jobs:
			job()
		case <-p.stop:
			return
		}
	}
}

func (p *Pool) AddJob(job Job) {
	p.jobs <- job
}

func (p *Pool) StopAll() {
	close(p.stop)
}

func (p *Pool) Wait() {
	p.wg.Wait()
}
