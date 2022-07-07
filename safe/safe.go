package safe

import (
	"context"
	"errors"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"sync"
	"time"
)

type routineCtx func(ctx context.Context)

type GoPool struct {
	wg     sync.WaitGroup
	ctx    context.Context
	cancel context.CancelFunc
}

func newPool(parentCtx context.Context) *GoPool {
	ctx, cancelFunc := context.WithCancel(parentCtx)
	return &GoPool{
		ctx:    ctx,
		cancel: cancelFunc,
	}
}

func (p *GoPool) Execute(goroutine routineCtx) {
	p.wg.Add(1)
	defer p.wg.Done()
	go func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err)
			}
		}()
		goroutine(p.ctx)
	}()

}

func (p *GoPool) ExecuteWithRecover(goroutine routineCtx, customerRecoverFunc func(err interface{})) {
	p.wg.Add(1)
	defer p.wg.Done()
	go func() {
		defer func() {
			if err := recover(); err != nil {
				if customerRecoverFunc != nil {
					customerRecoverFunc(err)
				} else {
					printStack(err)
				}
			}
			goroutine(p.ctx)
		}()
	}()
}

func (p *GoPool) Stop() {
	p.cancel()
	p.wg.Wait()
}

func printStack(err interface{}) {
	fmt.Printf("err is: %+v\n", err)
}

var (
	paramError = errors.New("param error")
)

const defaultExpireTime = 10

type ConcurrentJob struct {
	Id  uint64
	Job func()
}

type ConcurrentGoroutinePool struct {
	routinePool  *GoPool
	workerNumber int
	queue        chan *ConcurrentJob
	servingMap   sync.Map
	uuid         string
}

func NewConcurrentGoroutinePool(ctx context.Context, workerNumber, channelSize int) (*ConcurrentGoroutinePool, error) {
	p := newPool(ctx)
	if workerNumber <= 0 || channelSize <= 0 {
		return nil, paramError
	}

	pool := &ConcurrentGoroutinePool{
		routinePool:  p,
		workerNumber: workerNumber,
		queue:        make(chan *ConcurrentJob, channelSize),
		uuid:         uuid.NewV4().String(),
	}
	pool.start()
	return pool, nil

}

func (p *ConcurrentGoroutinePool) start() {
	for i := 0; i < p.workerNumber; i++ {
		p.routinePool.Execute(
			func(ctx context.Context) {
				for j := range p.queue {
					id := j.Id
					j.Job()
					p.servingMap.Delete(id)
				}
			})
	}
}

func (p *ConcurrentGoroutinePool) CommitJob(task *ConcurrentJob) bool {
	id := task.Id
	if value, ok := p.servingMap.Load(id); ok {
		now := uint32(time.Now().Unix())
		expireTime := value.(uint32)
		// 10 seconds
		if now-expireTime < defaultExpireTime {
			return false
		}
	}
	p.servingMap.Store(id, uint32(time.Now().Unix()))
	p.queue <- task
	return true
}
