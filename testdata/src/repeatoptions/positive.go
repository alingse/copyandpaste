package repeatoptions

type Cfg struct {
	name   string
	width  int
	height int
}

func NewCfg(name string, opts ...func(c *Cfg)) *Cfg {
	cfg := &Cfg{
		name: name,
	}
	for _, opt := range opts {
		opt(cfg)
	}
	return cfg
}

func WithWidth(w int) func(c *Cfg) {
	return func(c *Cfg) {
		c.width = w
	}
}

func WithHeight(h int) func(c *Cfg) {
	return func(c *Cfg) {
		c.height = h
	}
}

func Demo(w, h int) {
	_ = NewCfg("hello",
		WithWidth(w),
		WithHeight(h),
		WithWidth(w), // want `repeat option`
	)

	opts := []func(c *Cfg){
		WithWidth(w),
		WithHeight(h),
		WithHeight(h), // want `repeat option`
	}

	_ = append(opts,
		WithWidth(w),
		WithHeight(h),
		WithWidth(w), // want `repeat option`
	)
}
