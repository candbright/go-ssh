package ssh

import (
	"bytes"
	"errors"
	"fmt"
	"golang.org/x/crypto/ssh"
	"golang.org/x/text/encoding/simplifiedchinese"
	"io"
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
		Timeout:         time.Second * 10,
		User:            s.user,
		Auth:            auth,
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
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

func (s *RemoteSession) Run(name string, arg ...string) error {
	_, err := s.Output(name, arg...)
	return err
}

func (s *RemoteSession) Output(name string, arg ...string) ([]byte, error) {
	err := s.Connect()
	if err != nil {
		return nil, err
	}
	sess, err := s.client.NewSession()
	if err != nil {
		return nil, err
	}
	defer sess.Close()
	var errBuffer bytes.Buffer
	sess.Stderr = &errBuffer
	output, err := sess.Output(Command(name, arg...))
	errStr, _ := simplifiedchinese.GBK.NewDecoder().String(errBuffer.String())
	if err != nil {
		Fail(s.writer, name, arg, errStr, err)
		return nil, err
	} else {
		Success(s.writer, name, arg, string(output))
		return output, nil
	}
}

func (s *RemoteSession) CombinedOutput(name string, arg ...string) ([]byte, error) {
	err := s.Connect()
	if err != nil {
		return nil, err
	}
	sess, err := s.client.NewSession()
	if err != nil {
		return nil, err
	}
	defer sess.Close()
	output, err := sess.CombinedOutput(Command(name, arg...))
	if err != nil {
		Fail(s.writer, name, arg, string(output), err)
		return nil, err
	} else {
		Success(s.writer, name, arg, string(output))
		return output, nil
	}
}

func (s *RemoteSession) OutputGrep(cmdList []Cmd) ([]byte, error) {
	if cmdList == nil {
		return nil, errors.New("execute cmd grep failed, cmd list is nil")
	}
	//TODO
	return nil, nil
}

func (s *RemoteSession) Exists(path string) (bool, error) {
	var err error
	var output []byte
	output, err = s.Output("dir", "/b", path)
	if err != nil {
		return false, err
	}
	if len(output) != 0 {
		return true, nil
	} else {
		return false, nil
	}
}

func (s *RemoteSession) ReadFile(fileName string) ([]byte, error) {
	//TODO
	return nil, nil
}

func (s *RemoteSession) ReadDir(dir string) ([]FileInfo, error) {
	//TODO
	return nil, nil
}

func (s *RemoteSession) MakeDirAll(path string, perm os.FileMode) error {
	//TODO
	return nil
}

func (s *RemoteSession) Remove(name string) error {
	//TODO
	return nil
}

func (s *RemoteSession) RemoveAll(path string) error {
	//TODO
	return nil
}

func (s *RemoteSession) Create(name string) error {
	//TODO
	return nil
}

func (s *RemoteSession) WriteString(name string, data string, mode ...string) error {
	//TODO
	return nil
}
