package options

const LocalIp = "127.0.0.1"

var localIps = []string{LocalIp}

type host struct {
	Local    bool
	Ip       string
	Port     uint16
	User     string
	Password string
}

func (o host) Set(opt *Options) error {
	opt.Local = o.Local
	opt.Ip = o.Ip
	opt.Port = o.Port
	opt.User = o.User
	opt.Password = o.Password
	return nil
}

func AddLocalIp(ips ...string) {
	if ips == nil {
		return
	}
	for _, ip := range ips {
		addLocalIp(ip)
	}
}

func addLocalIp(ip string) {
	if localIps == nil {
		localIps = make([]string, 0)
	}
	for _, localIp := range localIps {
		if localIp == ip {
			return
		}
	}
	localIps = append(localIps, ip)
}

func LocalHost() Option {
	return host{
		Local: true,
		Ip:    LocalIp,
		Port:  22,
		User:  "root",
	}
}

func RemoteHost(ip string, port uint16) Option {
	return RemoteHostPWD(ip, port, "", "")
}

func RemoteHostPWD(ip string, port uint16, user string, password string) Option {
	if ip == "" {
		ip = LocalIp
	}
	if port == 0 {
		port = 22
	}
	if user == "" {
		user = "root"
	}
	for _, localIp := range localIps {
		if ip == localIp {
			return LocalHost()
		}
	}
	return host{
		Local:    false,
		Ip:       ip,
		Port:     port,
		User:     user,
		Password: password,
	}
}
