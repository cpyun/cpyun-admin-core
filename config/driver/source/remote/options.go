package remote

import (
	"context"
	"github.com/cpyun/gyopls-core/config/driver/source"
)

// provider
type remoteProvider struct {
	name, endpoint, path string
}

func WithProvider(name, endpoint, path string) source.Option {
	return source.OptionFunc(func(o *source.Options) {
		if o.Context == nil {
			o.Context = context.Background()
		}
		o.Context = context.WithValue(o.Context, remoteProvider{}, remoteProvider{name, endpoint, path})
	})
}
