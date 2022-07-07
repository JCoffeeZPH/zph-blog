package utils

import (
	"errors"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type Task struct {
	// 参数个数和类型不确定
	HandlerFunc func(v ...interface{})
	Params      []interface{}
}

type Pool struct {
	capacity       uint64
	runningWorkers uint64
	status         int64
	chTask         chan *Task
	sync.Mutex
	PanicHandler func(interface{})
}

const (
	RUNNING = 1
	STOP    = 0
)

var errInvalidPoolCap = errors.New("invalid pool cap")

func NewPool(capacity uint64) (*Pool, error) {
	if capacity <= 0 {
		return nil, errInvalidPoolCap
	}
	pool := &Pool{
		capacity: capacity,
		status:   RUNNING,
		chTask:   make(chan *Task, capacity),
	}
	return pool, nil
}

func (p *Pool) run() {
	p.incrRunningWorkers()

	go func() {
		defer func() {
			p.decRunningWorkers()
			if err := recover(); err != nil {
				if p.PanicHandler != nil {
					p.PanicHandler(err)
				} else {
					fmt.Printf("err is: %+v\n", err)
				}
			}
			p.checkWorker()
		}()
		for {
			select {
			case task, ok := <-p.chTask:
				if !ok {
					return
				}
				task.HandlerFunc(task.Params...)
			}
		}
	}()
}

func (p *Pool) incrRunningWorkers() {
	atomic.AddUint64(&p.runningWorkers, 1)
}

func (p *Pool) decRunningWorkers() {
	atomic.AddUint64(&p.runningWorkers, ^uint64(0))
}

func (p *Pool) GetRunningWorkers() uint64 {
	return atomic.LoadUint64(&p.runningWorkers)
}

func (p *Pool) GetCap() uint64 {
	return p.capacity
}

func (p *Pool) setStatus(status int64) bool {
	p.Lock()
	defer p.Unlock()

	if p.status == status {
		return false
	}

	p.status = status
	return true
}

func (p *Pool) Put(task *Task) error {
	p.Lock()
	defer p.Unlock()

	if p.status == STOP {
		return errInvalidPoolCap
	}

	if p.GetRunningWorkers() < p.GetCap() {
		p.run()
	}

	if p.status == RUNNING {
		p.chTask <- task
	}

	return nil
}

func (p *Pool) Close() {
	if !p.setStatus(STOP) {
		return
	}
	if len(p.chTask) > 0 {
		time.Sleep(1e6)
	}

	close(p.chTask)
}

func (p *Pool) checkWorker() {
	p.Lock()
	defer p.Unlock()

	if p.GetRunningWorkers() == 0 && len(p.chTask) > 0 {
		p.run()
	}
}
