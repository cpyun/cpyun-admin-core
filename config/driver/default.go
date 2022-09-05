package driver

type config struct {
	opts Options
}

func newConfig(opts ...OptionFunc) (Config, error) {
	var c = &config{}

	err := c.Init(opts...)
	if err != nil {
		return nil, err
	}

	return c, nil
}

//
func (c *config) Init(opts ...OptionFunc) error {
	c.opts = Options{}

	for _, o := range opts {
		o(&c.opts)
	}

	return nil
}

//
func (c *config) Options() Options {
	return c.opts
}
