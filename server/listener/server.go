/*
 * @Author: lwnmengjing
 * @Date: 2021/6/8 2:04 下午
 * @Last Modified by: lwnmengjing
 * @Last Modified time: 2021/6/8 2:04 下午
 */

package listener

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"sync"

	log "github.com/cpyun/gyopls-core/logger"
	"github.com/cpyun/gyopls-core/server"
)

type Server struct {
	name string
	//ctx     context.Context
	srv     *http.Server
	mux     sync.Mutex
	opts    options
	started bool
}

// New 实例化
func New(name string, opts ...Option) server.Runnable {
	s := &Server{
		name: name,
		opts: setDefaultOption(),
	}

	s.applyOptions(opts...)
	return s
}

// Options 设置参数
func (e *Server) applyOptions(opts ...Option) {
	for _, o := range opts {
		o(&e.opts)
	}
}

func (e *Server) String() string {
	return e.name
}

// Start 开始
func (e *Server) Start(ctx context.Context) (err error) {
	e.mux.Lock()
	defer e.mux.Unlock()

	e.started = true
	e.srv = &http.Server{
		Addr:         e.opts.addr,
		Handler:      e.opts.handler,
		ReadTimeout:  e.opts.readTimeout,
		WriteTimeout: e.opts.writeTimeout,
		BaseContext: func(_ net.Listener) context.Context {
			return ctx
		},
	}
	if e.opts.endHook != nil {
		e.srv.RegisterOnShutdown(e.opts.endHook)
	}

	go func() {
		if e.opts.cert != nil {
			err = e.srv.ListenAndServeTLS(e.opts.cert.certFile, e.opts.cert.keyFile)
		} else {
			err = e.srv.ListenAndServe()
		}
		if err != nil && err != http.ErrServerClosed {
			log.Errorf("%s Server start error: %s", e.name, err.Error())
		}
	}()

	go func() {
		if err = e.Shutdown(ctx); err != nil && err != context.Canceled {
			log.Errorf("%s server shutdown error: %s", e.name, err.Error())
		}
	}()

	if e.opts.startedHook != nil {
		e.opts.startedHook()
	}

	fmt.Printf("- [%s] Server listening on %s \r\n", e.name, e.srv.Addr)
	return nil
}

// Attempt 判断是否可以启动
func (e *Server) Attempt() bool {
	return !e.started
}

// Shutdown 停止
func (e *Server) Shutdown(ctx context.Context) error {
	<-ctx.Done()
	return e.srv.Shutdown(ctx)
}
