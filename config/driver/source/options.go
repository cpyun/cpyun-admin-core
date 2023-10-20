package source

import "context"

type Options struct {
	Context context.Context
}
type Option interface {
	apply(*Options)
}

type OptionFunc func(*Options)

func (fn OptionFunc) apply(o *Options) {
	fn(o)
}

func NewOptions(opts ...Option) Options {
	options := Options{
		Context: context.Background(),
	}
	for _, opt := range opts {
		opt.apply(&options)
	}

	return options
}
