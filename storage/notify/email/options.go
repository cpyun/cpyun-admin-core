package email

import (
	"gopkg.in/gomail.v2"
	"io"
)

type Option func(*options)

type options struct {
	started bool
	from    string
	to      string
	cc      string
	bcc     string
	msg     *gomail.Message
}

func setDefaultOptions() options {
	return options{
		started: false,
		msg:     gomail.NewMessage(),
	}
}

func WithStarted(started bool) Option {
	return func(o *options) {
		o.started = started
	}
}

func WithFrom(from string) Option {
	return func(o *options) {
		o.msg.SetHeader("From", from)
		o.from = from
	}
}

func WithTo(to ...string) Option {
	return func(o *options) {
		o.msg.SetHeader("To", to...)
	}
}

func WithCc(cc ...string) Option {
	return func(o *options) {
		o.msg.SetHeader("Cc", cc...)
	}
}

func WithBcc(bcc ...string) Option {
	return func(o *options) {
		o.msg.SetHeader("Bcc", bcc...)
	}
}

func WithSubject() Option {
	return func(o *options) {
		o.msg.SetHeader("Subject", "test")
	}
}

// WithBody contentType, body string
// func WithBody(contentType, body string) Option
//
func WithBody(contentType, body string) Option {
	return func(o *options) {
		o.msg.SetBody(contentType, body)
	}
}

// WithAttr 本地附件
func WithAttr(filename string) Option {
	return func(o *options) {
		o.msg.Attach(filename)
	}
}

func WithAttrWriter(filename string, src io.Reader) Option {
	return func(o *options) {
		o.msg.Attach(filename,
			gomail.SetCopyFunc(func(w io.Writer) error {
				_, err := io.Copy(w, src)
				return err
			}))
	}
}

// WithEmbed 本地图片
func WithEmbed(filename string) Option {
	return func(o *options) {
		o.msg.Attach(filename)
	}
}

func WithEmbedWriter(filename string, src io.Reader) Option {
	return func(o *options) {
		o.msg.Embed(filename,
			gomail.SetCopyFunc(func(w io.Writer) error {
				_, err := io.Copy(w, src)
				return err
			}))
	}
}

func WithMsg(msg *gomail.Message) Option {
	return func(o *options) {
		o.msg = msg
	}
}
