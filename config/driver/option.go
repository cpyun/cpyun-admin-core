package driver

import (
	"context"
	"github.com/cpyun/cpyun-admin-core/config/driver/loader"
	"github.com/cpyun/cpyun-admin-core/config/driver/reader"
	"github.com/cpyun/cpyun-admin-core/config/driver/source"
)

type Option interface {
	apply(o *Options)
}

type OptionFunc func(o *Options)

func (fn OptionFunc) apply(o *Options) {
	fn(o)
}

type Options struct {
	Source []source.Source
	Loader loader.Loader
	Reader reader.Reader

	Context context.Context

	Entity Entity
}

func setDefaultOptions() Options {
	return Options{
		Loader: loader.NewLoaderViper(),
		Reader: reader.NewReaderViper(),
	}
}

// WithSource source 数据源
func WithSource(s ...source.Source) Option {
	return OptionFunc(func(o *Options) {
		o.Source = append(o.Source, s...)
	})
}

// WithEntity 实体
func WithEntity(e Entity) Option {
	return OptionFunc(func(o *Options) {
		o.Entity = e
	})
}
