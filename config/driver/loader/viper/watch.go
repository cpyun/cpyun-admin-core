package viper

import (
	"github.com/cpyun/cpyun-admin-core/config/driver/loader"
	"github.com/cpyun/cpyun-admin-core/config/driver/reader"
	"github.com/cpyun/cpyun-admin-core/config/driver/source"
	"time"
)

type watcher struct {
	exit    chan bool
	path    []string
	value   reader.Value
	reader  reader.Reader
	version string
	//updates chan updateValue
}

func (w *watcher) Next() (*loader.Snapshot, error) {
	cs := &source.ChangeSet{
		Data:      nil,
		Format:    w.reader.String(),
		Source:    "viper",
		Timestamp: time.Now(),
	}

	return &loader.Snapshot{
		ChangeSet: cs,
		Version:   w.version,
	}, nil
}
