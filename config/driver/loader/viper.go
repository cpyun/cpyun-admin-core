package loader

import (
	"errors"
	"github.com/cpyun/cpyun-admin-core/config/driver/source"
	"sync"
)

type viper struct {
	rwMtx   sync.RWMutex
	sources []source.Source
	sets    []*source.ChangeSet
}

func (v *viper) Load(sources ...source.Source) error {
	if len(sources) == 0 {
		return errors.New("source is empty")
	}

	for _, s := range sources {
		set, err := s.Read()
		if err != nil {
			continue
		}

		//
		v.rwMtx.Lock()
		v.sources = append(v.sources, s)
		v.sets = append(v.sets, set)
		v.rwMtx.Unlock()

		go s.Watch()
	}

	return nil
}

func NewLoaderViper() Loader {
	return &viper{}
}
