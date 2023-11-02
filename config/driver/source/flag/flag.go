package flag

import (
	"errors"
	"github.com/cpyun/cpyun-admin-core/config/driver/source"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type flag struct {
	opts source.Options
}

func (f *flag) Read() (*source.ChangeSet, error) {
	sets, ok := f.opts.Context.Value(flagSets{}).(*pflag.FlagSet)
	if !ok {
		return nil, errors.New("flag sets not found")
	}

	err := viper.BindPFlags(sets)
	return nil, err
}

func (f *flag) Watch() (source.Watcher, error) {
	return nil, source.ErrWatcherStopped
}

func (f *flag) String() string {
	return "flag"
}

func NewSourceFlag(opts ...source.Option) source.Source {
	return &flag{
		opts: source.NewOptions(opts...),
	}
}
