package flag

import (
	"encoding/json"
	"errors"
	"github.com/cpyun/cpyun-admin-core/config/driver/source"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"time"
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

	var changes map[string]any
	sets.VisitAll(func(flag *pflag.Flag) {
		changes[flag.Name] = flag.Value.String()
	})
	b, _ := json.Marshal(changes)

	cs := &source.ChangeSet{
		Format:    "json",
		Source:    f.String(),
		Timestamp: time.Now(),
		Data:      b,
	}
	cs.Checksum = cs.Sum()

	return cs, err
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
