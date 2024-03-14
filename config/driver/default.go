package driver

import (
	"errors"
	"github.com/cpyun/gyopls-core/config/driver/loader"
	"github.com/cpyun/gyopls-core/config/driver/reader"
	"sync"
	"time"
)

type config struct {
	values reader.Values
	snap   *loader.Snapshot
	rwMux  sync.RWMutex
	opts   Options
}

func newConfig(opts ...Option) (Config, error) {
	var c = &config{}

	if err := c.Init(opts...); err != nil {
		return nil, err
	}
	go c.run()

	return c, nil
}

//
func (c *config) Init(opts ...Option) error {
	var err error

	c.opts = setDefaultOptions()
	for _, o := range opts {
		o.apply(&c.opts)
	}

	// 读取数据源数据
	if err = c.opts.Loader.Load(c.opts.Source...); err != nil {
		return err
	}

	// 解析数据
	c.snap, err = c.opts.Loader.Snapshot()

	c.values, err = c.opts.Reader.Values(c.snap.ChangeSet)
	if err != nil {
		return err
	}

	// 绑定实体
	if c.opts.Entity != nil {
		_ = c.Scan(c.opts.Entity)
	}

	return nil
}

//
func (c *config) Options() Options {
	return c.opts
}

func (c *config) Map() interface{} {
	c.rwMux.RLock()
	defer c.rwMux.RUnlock()
	return c.values.Map()
}

func (c *config) Scan(v interface{}) error {
	c.rwMux.RLock()
	defer c.rwMux.RUnlock()
	return c.values.Scan(v)
}

func (c *config) watch(w loader.Watcher) error {
	snap, _ := w.Next()

	// 判断数据是否更新
	if c.snap.Version >= snap.Version {
		//c.rwMux.Unlock()
		return errors.New("no data updated")
	}

	c.rwMux.Lock()
	c.snap = snap
	// set values
	c.values, _ = c.opts.Reader.Values(c.snap.ChangeSet)
	if c.opts.Entity != nil {
		_ = c.values.Scan(c.opts.Entity)
		c.opts.Entity.OnChange()
	}

	c.rwMux.Unlock()
	return nil

}

func (c *config) run() {
	for {
		// 获取观察者
		w, err := c.opts.Loader.Watch()
		if err != nil {
			time.Sleep(3 * time.Second)
			continue
		}

		//
		if err = c.watch(w); err != nil {
			time.Sleep(time.Second)
			continue
		}
	}
}
