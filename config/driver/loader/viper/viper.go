package viper

import (
	"errors"
	"fmt"
	"github.com/cpyun/cpyun-admin-core/config/driver/loader"
	"github.com/cpyun/cpyun-admin-core/config/driver/reader"
	readerViper "github.com/cpyun/cpyun-admin-core/config/driver/reader/viper"
	"github.com/cpyun/cpyun-admin-core/config/driver/source"
	"strings"
	"sync"
	"time"
)

type viper struct {
	exit    chan bool
	sources []source.Source
	sets    []*source.ChangeSet
	snap    *loader.Snapshot
	values  reader.Values

	rwMux sync.RWMutex
	opts  loader.Options
}

func (v *viper) Load(sources ...source.Source) error {
	if len(sources) == 0 {
		return errors.New("source is empty")
	}

	var vErrors []string
	for i, s := range sources {
		set, err := s.Read()
		if err != nil {
			vErrors = append(vErrors, fmt.Sprintf("error loading source %s: %v", s, err))
			continue
		}
		//
		v.rwMux.Lock()
		v.sources = append(v.sources, s)
		v.sets = append(v.sets, set)
		v.rwMux.Unlock()

		go v.watcher(i, s)
	}

	// 更新v.snap
	if err := v.reload(); err != nil {
		vErrors = append(vErrors, err.Error())
	}

	if len(vErrors) > 0 {
		return errors.New(strings.Join(vErrors, "|"))
	}

	return nil
}

func (v *viper) watcher(idx int, s source.Source) {
	w, err := s.Watch()
	if err != nil {
		return
	}
	defer w.Stop()

	v.watch(idx, w)
	return
}

func (v *viper) watch(idx int, sw source.Watcher) {
	for {
		cs, err := sw.Next()
		if err != nil || cs == nil {
			continue
		}

		// 更新后赋值
		v.sets[idx] = cs
		_ = v.reload()
	}
}

func (v *viper) reload() error {
	v.rwMux.Lock()
	defer v.rwMux.Unlock()

	set, err := v.opts.Reader.Merge(v.sets...)
	if err != nil {
		return err
	}
	v.values, _ = v.opts.Reader.Values(set)
	v.snap = &loader.Snapshot{
		ChangeSet: set,
		Version:   genVer(),
	}

	return nil
}

func (v *viper) Snapshot() (*loader.Snapshot, error) {
	v.rwMux.RLock()
	snap := loader.Copy(v.snap)
	v.rwMux.RUnlock()
	return snap, nil
}

// Watch for changes
func (v *viper) Watch(path ...string) (loader.Watcher, error) {
	w := &watcher{
		path:    path,
		reader:  v.opts.Reader,
		version: v.snap.Version,
	}
	return w, nil
}

func (v *viper) String() string {
	return "viper"
}

func NewLoaderViper(opts ...loader.OptionFunc) loader.Loader {
	options := loader.Options{
		Reader: readerViper.NewReaderViper(),
	}
	for _, o := range opts {
		o(&options)
	}

	return &viper{
		opts: options,
		//sources: options.Source,
	}
}

func genVer() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}
