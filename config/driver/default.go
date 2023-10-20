package driver

import "github.com/cpyun/cpyun-admin-core/config/driver/source"

type config struct {
	opts Options
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
	val, err := c.opts.Reader.Values(&source.ChangeSet{})
	if err != nil {
		return err
	}

	// 绑定实体
	if c.opts.Entity != nil {
		_ = val.Scan(c.opts.Entity)
		//WithBind(c.opts.Entity)
	}

	return nil
}

//
func (c *config) Options() Options {
	return c.opts
}

func (c *config) run() {

	//if c.opts.Entity != nil {
	//	WithBind(c.opts.Entity)
	//	c.opts.Entity.OnChange()
	//}
	return
}
