package options

import (
	"github.com/candbright/go-log/log"
	options_log "github.com/candbright/go-log/options"
	"io"
)

type filename struct {
	LogPath string
}

func (o filename) Set(opt *Options) error {
	logger, err := log.New(options_log.Path(o.LogPath))
	if err != nil {
		return err
	}
	opt.Logger = logger
	return nil
}

func LogPath(path string) Option {
	return filename{
		LogPath: path,
	}
}

type writer struct {
	LogOutput io.Writer
}

func (o writer) Set(opt *Options) error {
	logger, err := log.New(options_log.Writer(o.LogOutput))
	if err != nil {
		return err
	}
	opt.Logger = logger
	return nil
}

func LogWriter(output io.Writer) Option {
	return writer{
		LogOutput: output,
	}
}

type logger struct {
	Logger *log.Logger
}

func (o logger) Set(opt *Options) error {
	opt.Logger = o.Logger
	return nil
}

func Logger(log *log.Logger) Option {
	return logger{
		Logger: log,
	}
}
