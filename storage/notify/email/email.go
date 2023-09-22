package email

import (
	"context"
	"gopkg.in/gomail.v2"
)

type Email struct {
	client *gomail.Dialer
	opts   options
}

func NewEmail(client *gomail.Dialer, opts ...Option) *Email {
	e := &Email{
		client: client,
		opts:   setDefaultOptions(),
	}
	opts = append(opts, WithFrom(client.Username))

	e.applyOptions(opts...)
	return e
}

func (e *Email) String() string {
	return "email"
}

func (e *Email) Start(ctx context.Context) error {
	return e.sendMsg(ctx)
}

// Attempt 是否可以使用
func (e *Email) Attempt() bool {
	return e.opts.started
}

func (e *Email) applyOptions(opts ...Option) {
	for _, opt := range opts {
		opt(&e.opts)
	}
}

func (e *Email) sendMsg(_ context.Context) error {
	e.opts.started = true

	return e.client.DialAndSend(e.opts.msg)
}
