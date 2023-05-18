package options

type single struct {
	Single bool
}

func (o single) Set(opt *Options) error {
	opt.Single = o.Single
	return nil
}

func EnableSingle() Option {
	return Single(true)
}

func Single(isSingle bool) Option {
	return single{
		Single: isSingle,
	}
}
