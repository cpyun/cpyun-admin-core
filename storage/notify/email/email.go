package email

import (
	"context"
	"gopkg.in/gomail.v2"
)

type Email struct {
	client  *gomail.Dialer
	opts    options
	started bool
}

func NewEmail(client *gomail.Dialer, opts ...OptionFunc) *Email {
	e := &Email{
		client: client,
		opts:   setDefaultOptions(),
	}
	e.opts.from = client.Username

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
	return !e.opts.disable
}

func (e *Email) applyOptions(opts ...OptionFunc) {
	for _, opt := range opts {
		opt(&e.opts)
	}
}

func (e *Email) sendMsg(_ context.Context) error {
	e.started = true

	return e.client.DialAndSend(e.opts.msg)
}
