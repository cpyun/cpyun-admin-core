package env

import (
	"github.com/cpyun/cpyun-admin-core/config/driver/source"
	"github.com/spf13/viper"
	"strings"
)

type env struct {
	prefixes         []string
	strippedPrefixes []string
	replacer         *strings.Replacer
	opts             source.Options
}

func (e *env) Read() (*source.ChangeSet, error) {
	// 按照前缀读取环境变量
	if len(e.prefixes) > 0 {
		viper.SetEnvPrefix(e.prefixes[0])
	}
	// 使用替代符替换
	viper.SetEnvKeyReplacer(e.replacer)

	// 自动加载环境变量
	viper.AutomaticEnv()

	return nil, nil
}

func (e *env) Watch() (source.Watcher, error) {
	return nil, source.ErrWatcherStopped
}

func (e *env) String() string {
	return "env"
}

func NewSourceEnv(opts ...source.Option) source.Source {
	options := source.NewOptions(opts...)
	replacer := strings.NewReplacer(".", "_")

	var pre []string
	if p, ok := options.Context.Value(prefixKey{}).([]string); ok {
		pre = p
	}

	if r, ok := options.Context.Value(replaceKey{}).(*strings.Replacer); ok {
		replacer = r
	}

	return &env{
		prefixes: pre,
		replacer: replacer,
		opts:     options,
	}
}
