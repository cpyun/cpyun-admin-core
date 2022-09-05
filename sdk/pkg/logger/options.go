/*
 * @Author: lwnmengjing
 * @Date: 2021/6/10 10:26 上午
 * @Last Modified by: lwnmengjing
 * @Last Modified time: 2021/6/10 10:26 上午
 */

package logger

type OptionFunc func(*options)

type options struct {
	driver string
	path   string
	level  string
	stdout string
	cap    uint
}

func setDefault() options {
	return options{
		driver: "default",
		path:   "temp/logs",
		level:  "warn",
		stdout: "default",
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

func WithCap(n uint) OptionFunc {
	return func(o *options) {
		o.cap = n
	}
}
