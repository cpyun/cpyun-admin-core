package loader

import (
	"context"
	"github.com/cpyun/cpyun-admin-core/config/driver/reader"
	"github.com/cpyun/cpyun-admin-core/config/driver/source"
)

type Options struct {
	Reader reader.Reader
	Source []source.Source

	// for alternative data
	Context context.Context
}

type OptionFunc func(o *Options)
