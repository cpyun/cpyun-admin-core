package flag

import (
	"context"
	"github.com/spf13/pflag"

	"github.com/cpyun/cpyun-admin-core/config/driver/source"
)

type flagSets struct{}

func WithFlagSets(set *pflag.FlagSet) source.Option {
	return source.OptionFunc(func(o *source.Options) {
		if o.Context == nil {
			o.Context = context.Background()
		}
		o.Context = context.WithValue(o.Context, flagSets{}, set)
	})
}
