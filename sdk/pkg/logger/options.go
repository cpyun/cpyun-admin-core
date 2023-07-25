package logger

import "time"

type OptionFunc func(*options)

type options struct {
	driver     string
	path       string
	level      string
	stdout     string
	timeFormat string //时间格式
	cap        uint
	maxAge     uint
	maxBackups uint
	compress   bool
}

func setDefault() options {
	return options{
		driver:     "default",
		path:       "temp/logs/warn.log",
		level:      "warn",
		stdout:     "default",
		timeFormat: time.RFC3339,
		cap:        0,
		maxAge:     0,
		maxBackups: 0,
		compress:   false,
	}
}

func WithType(s string) OptionFunc {
	return func(o *options) {
		o.driver = s
	}
}

func WithPath(s string) OptionFunc {
	return func(o *options) {
		o.path = s
	}
}

func WithLevel(s string) OptionFunc {
	return func(o *options) {
		o.level = s
	}
}

func WithStdout(s string) OptionFunc {
	return func(o *options) {
		o.stdout = s
	}
}

func WithTimeFormat(s string) OptionFunc {
	return func(o *options) {
		o.timeFormat = s
	}
}

func WithCap(n uint) OptionFunc {
	return func(o *options) {
		o.cap = n
	}
}

func WithMaxAge(n uint) OptionFunc {
	return func(o *options) {
		o.maxAge = n
	}
}

func WithMaxBackups(n uint) OptionFunc {
	return func(o *options) {
		o.maxBackups = n
	}
}

func WithCompress(n bool) OptionFunc {
	return func(o *options) {
		o.compress = n
	}
}
