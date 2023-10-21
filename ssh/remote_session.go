package ssh

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"io"
	"net"
	"os"
	"time"
)

type RemoteSession struct {
	client     *ssh.Client
	writer     io.Writer
	ip         string
	port       uint16
	user       string
	password   string
	sshKeyPath string
}

func (s *RemoteSession) success(name string, arg []string, output string) {
	if s.writer != nil {
		cmdLog := &CmdLog{
			Cmd:    Command(name, arg...),
			Output: output,
			Time:   time.Now().Format("2006-01-02 15:04:05"),
			Host:   s.ip,
		}
		_, _ = s.writer.Write([]byte(fmt.Sprint(cmdLog.ToJson(), "\n")))
	}
}

func (s *RemoteSession) fail(name string, arg []string, output string, err error) {
	if s.writer != nil {
		cmdLog := &CmdLog{
			Cmd:       Command(name, arg...),
			ErrString: err.Error(),
			Output:    output,
			Time:      time.Now().Format("2006-01-02 15:04:05"),
			Host:      s.ip,
		}
		_, _ = s.writer.Write([]byte(fmt.Sprint(cmdLog.ToJson(), "\n")))
	}
}

func (s *RemoteSession) Reconnect() error {
	if s.client != nil {
		err := s.client.Close()
		if err != nil {
			return err
		}
		s.client = nil
	}
	err := s.Connect()
	if err != nil {
		return err
	}
	return nil
}

func (s *RemoteSession) Connect() error {
	if s.client != nil {
		return nil
	}
	var auth []ssh.AuthMethod
	if s.password != "" {
		auth = []ssh.AuthMethod{ssh.Password(s.password)}
	} else {
		keyAuth, err := publicKeyAuth(s.sshKeyPath)
		if err != nil {
			return err
		}
		auth = []ssh.AuthMethod{keyAuth}
	}
	sshCfg := &ssh.ClientConfig{
		Timeout: time.Second * 10,
		User:    s.user,
		Auth:    auth,
		HostKeyCallback: ssh.HostKeyCallback(func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		}),
	}
	addr := fmt.Sprintf("%s:%d", s.ip, s.port)
	sess, err := ssh.Dial("tcp", addr, sshCfg)
	if err != nil {
		return err
	}
	s.client = sess
	return nil
}

func publicKeyAuth(kPath string) (ssh.AuthMethod, error) {
	key, err := os.ReadFile(kPath)
	if err != nil {
		return nil, err
	}
	singer, err := ssh.ParsePrivateKey(key)
	if err != nil {
		return nil, err
	}
	return ssh.PublicKeys(singer), nil
}

func (s *RemoteSession) Close() error {
	if s.client != nil {
		return s.client.Close()
	}
	return nil
}
