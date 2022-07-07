package safe

//
//import (
//	"context"
//	"fmt"
//	"sync"
//)
//
//type routineCtx func(ctx context.Context)
//
//type Pool struct {
//	waitGroup sync.WaitGroup
//	ctx       context.Context
//	cancel    context.CancelFunc
//}
//
//func NewPool(parentCtx context.Context) *Pool {
//	c, cancelFunc := context.WithCancel(parentCtx)
//	return &Pool{
//		ctx:    c,
//		cancel: cancelFunc,
//	}
//}
//
//func (p *Pool) GoWithRecover(goroutine func(), customerRecover func(err interface{})) {
//	go func() {
//		defer func() {
//			if err := recover(); err != nil {
//				if customerRecover != nil {
//					customerRecover(err)
//				} else {
//					printStack(err)
//				}
//			}
//		}()
//		goroutine()
//	}()
//}
//
//func printStack(err interface{}) {
//	// todo log
//	fmt.Printf("err is: %+v\n", err)
//}
//
//func defaultRecover(err interface{}) {
//	printStack(err)
//}
//
//func (p *Pool) Go(routine routineCtx) {
//	p.waitGroup.Add(1)
//	p.GoWithRecover(func() {
//		defer p.waitGroup.Done()
//		routine(p.ctx)
//	}, defaultRecover)
//}
//
//func (p *Pool) Stop() {
//	p.cancel()
//	p.waitGroup.Wait()
//}
