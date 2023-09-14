/*
 * @Author: lwnmengjing
 * @Date: 2021/6/8 2:15 下午
 * @Last Modified by: lwnmengjing
 * @Last Modified time: 2021/6/8 2:15 下午
 */

package listener

import (
	"context"
	"net/http"
	"time"
)

// Option 参数设置类型
type Option func(*options)

type options struct {
	addr         string
	cert         *cert
	handler      http.Handler
	readTimeout  time.Duration
	writeTimeout time.Duration
	ctx          context.Context
	startedHook  func()
	endHook      func()
}

type cert struct {
	certFile, keyFile string
}

func setDefaultOption() options {
	return options{
		addr:         ":8080",
		handler:      http.NotFoundHandler(),
		readTimeout:  60 * time.Second,
		writeTimeout: 60 * time.Second,
	}
}

// WithAddr 设置addr
func WithAddr(s string) Option {
	return func(o *options) {
		o.addr = s
	}
}

// WithTlsOption 设置cert
func WithTlsOption(certFile, keyFile string) Option {
	return func(o *options) {
		o.cert = &cert{
			certFile: certFile,
			keyFile:  keyFile,
		}
	}
}

// WithHandler 设置handler
func WithHandler(handler http.Handler) Option {
	return func(o *options) {
		o.handler = handler
	}
}

// WithReadTimeout 设置读超时
func WithReadTimeout(d int) Option {
	return func(o *options) {
		o.readTimeout = time.Duration(d)
	}
}

// WithWriteTimeout 设置写超时
func WithWriteTimeout(d int) Option {
	return func(o *options) {
		o.writeTimeout = time.Duration(d)
	}
}

// WithStartedHook 设置启动回调函数
func WithStartedHook(f func()) Option {
	return func(o *options) {
		o.startedHook = f
	}
}

// WithEndHook 设置结束回调函数
func WithEndHook(f func()) Option {
	return func(o *options) {
		o.endHook = f
	}
}

// WithCert 设置cert
//
// Deprecated: Set tls cert.
// punctuation properly. Use WithTlsOption instead.
func WithCert(s string) Option {
	return func(o *options) {
		o.cert = &cert{
			certFile: s,
		}
	}
}

// WithKey 设置key
//
// Deprecated: Set tls key.
// punctuation properly. Use WithKey instead.
func WithKey(s string) Option {
	return func(o *options) {
		o.cert = &cert{
			keyFile: s,
		}
	}
}
