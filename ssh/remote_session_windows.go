package ssh

import (
	"bytes"
	"fmt"
	"github.com/pkg/errors"
	"golang.org/x/text/encoding/simplifiedchinese"
	"os"
	"strings"
)

func (s *RemoteSession) Run(name string, arg ...string) error {
	_, err := s.Output(name, arg...)
	return err
}

func (s *RemoteSession) Output(name string, arg ...string) ([]byte, error) {
	sess, err := s.client.NewSession()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer sess.Close()
	var errBuffer bytes.Buffer
	sess.Stderr = &errBuffer
	output, err := sess.Output(Command(name, arg...))
	errStr, _ := simplifiedchinese.GBK.NewDecoder().String(errBuffer.String())
	if err != nil {
		s.fail(name, arg, errStr, err)
		return nil, errors.WithStack(err)
	} else {
		s.success(name, arg, string(output))
		return output, nil
	}
}

func (s *RemoteSession) CombinedOutput(name string, arg ...string) ([]byte, error) {
	sess, err := s.client.NewSession()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer sess.Close()
	output, err := sess.CombinedOutput(Command(name, arg...))
	if err != nil {
		s.fail(name, arg, string(output), err)
		return nil, errors.WithStack(err)
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
	var err error
	var output []byte
	output, err = s.Output("type", fileName)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func (s *RemoteSession) ReadDir(dir string) ([]FileInfo, error) {
	output, err := s.Output("dir", "/a", "/b", dir)
	if err != nil {
		return nil, err
	}
	files := make([]FileInfo, 0)
	for _, file := range strings.Split(string(output), "\n") {
		if file == "" {
			continue
		}
		if strings.HasSuffix(file, "\\") {
			files = append(files, FileInfo{Name: file, Path: dir + file})
		} else {
			files = append(files, FileInfo{Name: file, Path: dir + "\\" + file})
		}
	}
	return files, nil
}

func (s *RemoteSession) MakeDirAll(path string, perm os.FileMode) error {
	return s.Run("mkdir", strings.ReplaceAll(path, "/", "\\"))
}

func (s *RemoteSession) Remove(name string) error {
	return s.Run("del", name)
}

func (s *RemoteSession) RemoveAll(path string) error {
	return s.Run("rd", "/s", "/q", path)
}

func (s *RemoteSession) Create(name string) error {
	return s.Run("type", "nul", ">", name)
}

func (s *RemoteSession) WriteString(name string, data string, mode ...string) error {
	flag := ">"
	if len(mode) == 1 && mode[0] == ">>" {
		flag = ">>"
	}
	return s.Run("echo", fmt.Sprintf(`'%s'`, data), flag, name)
}
