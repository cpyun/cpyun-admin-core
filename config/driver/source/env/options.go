package env

import (
	"context"
	"github.com/cpyun/gyopls-core/config/driver/source"
	"strings"
)

type strippedPrefixKey struct{}
type prefixKey struct{}
type replaceKey struct{}

func WithStrippedPrefix(p ...string) source.Option {
	return source.OptionFunc(func(o *source.Options) {
		if o.Context == nil {
			o.Context = context.Background()
		}
		o.Context = context.WithValue(o.Context, strippedPrefixKey{}, appendUnderscore(p))
	})
}

func WithPrefix(p ...string) source.Option {
	return source.OptionFunc(func(o *source.Options) {
		if o.Context == nil {
			o.Context = context.Background()
		}
		o.Context = context.WithValue(o.Context, prefixKey{}, p)
	})
}

func WithReplace(oldNew ...string) source.Option {
	return source.OptionFunc(func(o *source.Options) {
		if o.Context == nil {
			o.Context = context.Background()
		}

		replace := strings.NewReplacer(oldNew...)
		o.Context = context.WithValue(o.Context, replaceKey{}, replace)
	})
}

func appendUnderscore(prefixes []string) []string {
	//nolint:prealloc
	var result []string
	for _, p := range prefixes {
		if !strings.HasSuffix(p, "_") {
			result = append(result, p+"_")
			continue
		}

		result = append(result, p)
	}

	return result
}
