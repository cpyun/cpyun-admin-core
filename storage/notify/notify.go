package notify

import (
	"context"
	"fmt"
	"sync"
)

type Notify struct {
	services    map[string]Provider
	mutex       sync.Mutex
	errChan     chan error
	internalCtx context.Context
}

func NewNotify() Manager {
	return &Notify{
		services: make(map[string]Provider),
		errChan:  make(chan error),
	}
}

func (e *Notify) Add(r ...Provider) {
	for i := range r {
		e.services[r[i].String()] = r[i]
	}
}

func (e *Notify) Run(ctx context.Context) (err error) {
	e.mutex.Lock()
	defer e.mutex.Unlock()

	e.internalCtx = ctx
	e.errChan = make(chan error)

	for k := range e.services {
		if !e.services[k].Attempt() {
			//先判断是否可以启动
			return fmt.Errorf("[%s] disabled startup, unable to send messages", e.services[k].String())
		}
	}
	//按顺序启动
	for k := range e.services {
		go e.startSend(e.services[k])
	}

	select {
	case <-ctx.Done():
		return nil
	case err = <-e.errChan:
		return err
	}
}

func (e *Notify) startSend(r Provider) {
	if err := r.Start(e.internalCtx); err != nil {
		e.errChan <- err
	}
}
