package options

type sshKeyPath struct {
	SshKeyPath string
}

func (o sshKeyPath) Set(opt *Options) error {
	opt.SshKeyPath = o.SshKeyPath
	return nil
}

func SshKeyPath(path string) Option {
	return sshKeyPath{
		SshKeyPath: path,
	}
}
