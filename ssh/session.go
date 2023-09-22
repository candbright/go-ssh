package ssh

import (
	"github.com/candbright/go-ssh/ssh/options"
	"os"
	"sync"
)

func AddLocalIp(ips ...string) {
	options.AddLocalIp(ips...)
}

type Cmd struct {
	Name string
	Arg  []string
}

type Session interface {
	Run(name string, arg ...string) error
	Output(name string, arg ...string) ([]byte, error)
	CombinedOutput(name string, arg ...string) ([]byte, error)
	OutputGrep(cmdList []Cmd) ([]byte, error)
	Exists(path string) (bool, error)
	ReadFile(fileName string) ([]byte, error)
	ReadDir(dir string) ([]FileInfo, error)
	MakeDirAll(path string, perm os.FileMode) error
	Remove(name string) error
	RemoveAll(path string) error
	Create(name string) error
	WriteString(name string, data string, mode ...string) error
}

type session struct {
	Session
}

func NewSession(opt ...options.Option) (Session, error) {
	o := options.Default()
	var err error
	for _, option := range opt {
		err = option.Set(&o)
		if err != nil {
			return nil, err
		}
	}
	var s Session
	if o.Local {
		s = &LocalSession{writer: o.Writer}
	} else {
		s = &RemoteSession{
			writer:     o.Writer,
			ip:         o.Ip,
			port:       o.Port,
			user:       o.User,
			password:   o.Password,
			sshKeyPath: o.SshKeyPath,
		}
	}
	if o.Single {
		s = &SingleSession{
			lock:    &sync.Mutex{},
			Session: s,
		}
	}
	return &session{s}, nil
}
