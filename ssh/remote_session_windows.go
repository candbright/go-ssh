package ssh

import (
	"bytes"
	"errors"
	"golang.org/x/text/encoding/simplifiedchinese"
	"os"
)

func (s *RemoteSession) Run(name string, arg ...string) error {
	_, err := s.Output(name, arg...)
	return err
}

func (s *RemoteSession) Output(name string, arg ...string) ([]byte, error) {
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
		s.fail(name, arg, errStr, err)
		return nil, err
	} else {
		s.success(name, arg, string(output))
		return output, nil
	}
}

func (s *RemoteSession) CombinedOutput(name string, arg ...string) ([]byte, error) {
	sess, err := s.client.NewSession()
	if err != nil {
		return nil, err
	}
	defer sess.Close()
	output, err := sess.CombinedOutput(Command(name, arg...))
	if err != nil {
		s.fail(name, arg, string(output), err)
		return nil, err
	} else {
		s.success(name, arg, string(output))
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
