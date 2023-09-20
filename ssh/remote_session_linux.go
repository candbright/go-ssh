package ssh

import (
	"bytes"
	"errors"
	"fmt"
	"golang.org/x/crypto/ssh"
	"golang.org/x/text/encoding/simplifiedchinese"
	"io"
	"os"
	"strconv"
	"strings"
	"time"
)

type RemoteSession struct {
	client   *ssh.Client
	writer   io.Writer
	ip       string
	port     uint16
	user     string
	password string
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
		sshKeyPath := "/root/.ssh/id_rsa"
		keyAuth, err := publicKeyAuth(sshKeyPath)
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
	cmdStrList := make([]string, len(cmdList))
	for i, cmd := range cmdList {
		cmdStrList[i] = cmd.Name
		for _, arg := range cmd.Arg {
			cmdStrList[i] += " " + arg
		}
	}
	name := "bash"
	arg := []string{"-c", strings.Join(cmdStrList, " | ")}
	output, err := s.Output(name, arg...)
	if err != nil {
		return nil, err
	} else {
		return output, nil
	}
}

func (s *RemoteSession) Exists(path string) (bool, error) {
	var err error
	var output []byte
	if strings.HasSuffix(path, "/") {
		//dir
		output, err = s.Output("find", path, "-prune")
	} else {
		//file
		output, err = s.Output("find", Dir(path), "-name", FileName(path))
	}
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
	return s.Output("cat", fileName)
}

func (s *RemoteSession) ReadDir(dir string) ([]FileInfo, error) {
	output, err := s.Output("ls", "-AF", dir)
	if err != nil {
		return nil, err
	}
	files := make([]FileInfo, 0)
	for _, file := range strings.Split(string(output), "\n") {
		if strings.HasSuffix(file, "/") {
			files = append(files, FileInfo{Name: file, Path: dir + file})
		} else {
			files = append(files, FileInfo{Name: file, Path: dir + "/" + file})
		}
	}
	return files, nil
}

func (s *RemoteSession) MakeDirAll(path string, perm os.FileMode) error {
	return s.Run("mkdir", "-p", path, "-m", strconv.FormatUint(uint64(perm), 10))
}

func (s *RemoteSession) Remove(name string) error {
	return s.Run("rm", "-f", name)
}

func (s *RemoteSession) RemoveAll(path string) error {
	return s.Run("rm", "-r", "-f", path)
}

func (s *RemoteSession) Create(name string) error {
	exists, err := s.Exists(Dir(name))
	if err != nil {
		return err
	}
	if !exists {
		err = s.MakeDirAll(Dir(name), 0666)
		if err != nil {
			return err
		}
	}
	return s.Run("touch", name)
}

func (s *RemoteSession) WriteString(name string, data string, mode ...string) error {
	flag := ">"
	if len(mode) == 1 && mode[0] == ">>" {
		flag = ">>"
	}
	return s.Run("echo", fmt.Sprintf(`"%s"`, data), flag, name)
}
