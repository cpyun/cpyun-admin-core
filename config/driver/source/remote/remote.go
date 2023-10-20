package remote

import (
	"github.com/cpyun/cpyun-admin-core/config/driver/source"
	"github.com/spf13/viper"
)

type remote struct {
	opts source.Options
}

func (r *remote) Read() (*source.ChangeSet, error) {
	provider := r.opts.Context.Value(remoteProvider{}).(remoteProvider)

	err := viper.AddRemoteProvider(provider.name, provider.endpoint, provider.path)
	if err != nil {
		return nil, err
	}

	err = viper.ReadRemoteConfig()
	return nil, err
}

func (r *remote) Watch() (source.Watcher, error) {
	return nil, nil
}

func (r *remote) String() string {
	return "remote"
}

func NewSourceRemote(opts ...source.Option) source.Source {
	options := source.NewOptions(opts...)

	return &remote{
		opts: options,
	}
}
