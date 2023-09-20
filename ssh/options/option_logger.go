package options

import (
	"io"
	"os"
	"path"
)

type filename struct {
	LogPath string
}

func (o filename) Set(opt *Options) error {
	err := os.MkdirAll(path.Dir(o.LogPath), 0750)
	if err != nil && !os.IsExist(err) {
		return err
	}
	f, err := os.OpenFile(o.LogPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		return err
	}
	opt.Writer = f
	return nil
}

func LogPath(path string) Option {
	return filename{
		LogPath: path,
	}
}

type writer struct {
	Writer io.Writer
}

func (o writer) Set(opt *Options) error {
	opt.Writer = o.Writer
	return nil
}

func LogWriter(w io.Writer) Option {
	return writer{
		Writer: w,
	}
}
