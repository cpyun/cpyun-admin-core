package loader

import (
	"context"
	"github.com/cpyun/gyopls-core/config/driver/reader"
	"github.com/cpyun/gyopls-core/config/driver/source"
)

type Options struct {
	Reader reader.Reader
	Source []source.Source

	// for alternative data
	Context context.Context
}

type OptionFunc func(o *Options)
