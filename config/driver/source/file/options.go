package file

import (
	"context"
	"github.com/cpyun/gyopls-core/config/driver/source"
)

type filePathKey struct{}

func WithPath(p string) source.Option {
	return source.OptionFunc(func(o *source.Options) {
		if o.Context == nil {
			o.Context = context.Background()
		}
		o.Context = context.WithValue(o.Context, filePathKey{}, p)
	})
}
