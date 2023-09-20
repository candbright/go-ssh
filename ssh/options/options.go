package options

import (
	"io"
	"os"
)

type Options struct {
	Writer   io.Writer
	Local    bool
	Single   bool
	Ip       string
	Port     uint16
	User     string
	Password string
}

func Default() Options {
	return Options{
		Writer: os.Stdout,
		Local:  true,
		Single: true,
		Ip:     LocalIp,
		Port:   22,
	}
}

type Option interface {
	Set(opt *Options) error
}
