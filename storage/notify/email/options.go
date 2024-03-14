package email

import (
	"gopkg.in/gomail.v2"
	"io"
)

type OptionFunc func(*options)

type options struct {
	disable bool
	from    string
	to      []string
	cc      []string
	bcc     []string
	subject string
	msg     *gomail.Message
}

func setDefaultOptions() options {
	return options{
		disable: false,
		msg:     gomail.NewMessage(),
	}
}

func WithDisable(disable bool) OptionFunc {
	return func(o *options) {
		o.disable = disable
	}
}

func WithFrom(from string) OptionFunc {
	return func(o *options) {
		o.msg.SetHeader("From", from)
		o.from = from
	}
}

func WithTo(to ...string) OptionFunc {
	return func(o *options) {
		o.msg.SetHeader("To", to...)
		o.to = append(o.to, to...)
	}
}

func WithCc(cc ...string) OptionFunc {
	return func(o *options) {
		o.msg.SetHeader("Cc", cc...)
		o.cc = append(o.cc, cc...)
	}
}

func WithBcc(bcc ...string) OptionFunc {
	return func(o *options) {
		o.msg.SetHeader("Bcc", bcc...)
		o.bcc = append(o.bcc, bcc...)
	}
}

func WithSubject(s string) OptionFunc {
	return func(o *options) {
		o.msg.SetHeader("Subject", s)
		o.subject = s
	}
}

// WithBody contentType, body string
// func WithBody(contentType, body string) OptionFunc
//
func WithBody(contentType, body string) OptionFunc {
	return func(o *options) {
		o.msg.SetBody(contentType, body)
	}
}

// WithAttr 本地附件
func WithAttr(filename string, src ...io.Reader) OptionFunc {
	return func(o *options) {
		var settings []gomail.FileSetting
		for i := 0; i < len(src); i++ {
			settings = append(settings, gomail.SetCopyFunc(func(w io.Writer) error {
				_, err := io.Copy(w, src[i])
				return err
			}))
		}

		o.msg.Attach(filename, settings...)
	}
}

// WithEmbed 图片
// 	eg: WithEmbed("./1.jpg") 本地图片
func WithEmbed(filename string, src ...io.Reader) OptionFunc {
	return func(o *options) {
		var settings []gomail.FileSetting
		for i := 0; i < len(src); i++ {
			settings = append(settings, gomail.SetCopyFunc(func(w io.Writer) error {
				_, err := io.Copy(w, src[i])
				return err
			}))
		}

		o.msg.Embed(filename, settings...)
	}
}
