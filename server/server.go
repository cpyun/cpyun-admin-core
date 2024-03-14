/*
 * @Author: lwnmengjing
 * @Date: 2021/6/7 5:43 下午
 * @Last Modified by: lwnmengjing
 * @Last Modified time: 2021/6/7 5:43 下午
 */

package server

import (
	"context"
	"errors"
	"fmt"
	"sync"

	"github.com/cpyun/gyopls-core/logger"
)

type Server struct {
	services               map[string]Runnable
	mutex                  sync.Mutex
	errChan                chan error
	waitForRunnable        sync.WaitGroup
	internalCtx            context.Context
	internalCancel         context.CancelFunc
	internalProceduresStop chan struct{}
	shutdownCtx            context.Context
	shutdownCancel         context.CancelFunc
	logger                 *logger.Helper
	opts                   options
}

// New 实例化
func New(opts ...OptionFunc) *Server {
	s := &Server{
		services:               make(map[string]Runnable),
		errChan:                make(chan error),
		internalProceduresStop: make(chan struct{}),
	}
	s.opts = setDefaultOptions()
	s.withOptions(opts...)
	return s
}

func (e *Server) withOptions(opts ...OptionFunc) {
	for _, opt := range opts {
		opt(&e.opts)
	}
}

// Add 添加 runnable
func (e *Server) Add(r ...Runnable) {
	if e.services == nil {
		e.services = make(map[string]Runnable)
	}
	for i := range r {
		e.services[r[i].String()] = r[i]
	}
}

// Start 启动 runnable
func (e *Server) Start(ctx context.Context) (err error) {
	e.mutex.Lock()
	defer e.mutex.Unlock()
	e.internalCtx, e.internalCancel = context.WithCancel(context.Background())
	defer func() {
		err = e.shutdownStopComplete(err)
	}()
	e.errChan = make(chan error)

	for k := range e.services {
		if !e.services[k].Attempt() {
			//先判断是否可以启动
			return errors.New("can't accept new runnable as stop procedure is already engaged")
		}
	}
	//按顺序启动
	for k := range e.services {
		e.startRunnable(e.services[k])
	}
	e.waitForRunnable.Wait()
	select {
	case <-ctx.Done():
		return nil
	case err = <-e.errChan:
		return err
	}
}

func (e *Server) startRunnable(r Runnable) {
	e.waitForRunnable.Add(1)
	go func() {
		defer e.waitForRunnable.Done()
		if err := r.Start(e.internalCtx); err != nil {
			e.errChan <- err
		}
	}()
}

func (e *Server) shutdownStopComplete(err error) error {
	stopComplete := make(chan struct{})
	defer close(stopComplete)
	stopErr := e.engageStopProcedure(stopComplete)
	if stopErr != nil {
		if err != nil {
			err = fmt.Errorf("%s, %w", stopErr.Error(), err)
		} else {
			err = stopErr
		}
	}
	return err
}

func (e *Server) engageStopProcedure(stopComplete <-chan struct{}) error {
	if e.opts.gracefulShutdownTimeout > 0 {
		e.shutdownCtx, e.shutdownCancel = context.WithTimeout(context.Background(), e.opts.gracefulShutdownTimeout)
	} else {
		e.shutdownCtx, e.shutdownCancel = context.WithCancel(context.Background())
	}
	defer e.shutdownCancel()
	close(e.internalProceduresStop)
	e.internalCancel()

	go func() {
		for {
			select {
			case err, ok := <-e.errChan:
				if ok {
					e.logger.Error(err, "error received after stop sequence was engaged")
				}
			case <-stopComplete:
				return
			}
		}
	}()

	return e.waitForRunnableToEnd()
}

func (e *Server) waitForRunnableToEnd() error {
	if e.opts.gracefulShutdownTimeout == 0 {
		go func() {
			e.waitForRunnable.Wait()
			e.shutdownCancel()
		}()
	}
	select {
	case <-e.shutdownCtx.Done():
		if err := e.shutdownCtx.Err(); err != nil && err != context.Canceled && err != context.DeadlineExceeded {
			return fmt.Errorf(
				"failed waiting for all runnables to end within grace period of %s: %w",
				e.opts.gracefulShutdownTimeout, err)
		}
	}

	return nil
}

func (e *Server) Shutdown(ctx context.Context) (err error) {
	return e.shutdownStopComplete(err)
}
