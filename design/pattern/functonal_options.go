package pattern

type Options struct {
	ID   string
	Name string
}

type Option func(*Options)

func WithID(id string) Option {
	return func(o *Options) {
		o.ID = id
	}
}
func WithName(name string) Option {
	return func(o *Options) {
		o.Name = name
	}
}

func NewOptions(opts ...Option) *Options {
	o := &Options{}
	for _, opt := range opts {
		opt(o)
	}
	return o
}
