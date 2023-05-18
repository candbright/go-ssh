package options

import (
	"github.com/candbright/go-log/log"
)

type Options struct {
	Logger   *log.Logger
	Local    bool
	Single   bool
	Ip       string
	Port     uint16
	User     string
	Password string
}

func Default() Options {
	return Options{
		Logger: log.Instance(),
		Local:  true,
		Single: true,
		Ip:     LocalIp,
		Port:   22,
	}
}

type Option interface {
	Set(opt *Options) error
}
