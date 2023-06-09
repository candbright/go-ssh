package ssh

import (
	"bytes"
	"errors"
	"github.com/candbright/go-log/log"
	"os"
	"os/exec"
	"strings"
)

type LocalSession struct {
	logger *log.Logger
}

func (s *LocalSession) Run(name string, arg ...string) error {
	_, err := s.Output(name, arg...)
	return err
}

func (s *LocalSession) Output(name string, arg ...string) ([]byte, error) {
	c := exec.Command(name, arg...)
	var errBuffer bytes.Buffer
	c.Stderr = &errBuffer
	output, err := c.Output()
	if err != nil {
		Fail(s.logger, name, arg, errBuffer.String(), err)
		return nil, err
	} else {
		Success(s.logger, name, arg, string(output))
		return output, nil
	}
}

func (s *LocalSession) CombinedOutput(name string, arg ...string) ([]byte, error) {
	output, err := exec.Command(name, arg...).CombinedOutput()
	if err != nil {
		Fail(s.logger, name, arg, string(output), err)
		return nil, err
	} else {
		Success(s.logger, name, arg, string(output))
		return output, nil
	}
}

func (s *LocalSession) OutputGrep(cmdList []Cmd) ([]byte, error) {
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
	c := exec.Command(name, arg...)
	var errBuffer bytes.Buffer
	c.Stderr = &errBuffer
	output, err := c.Output()
	if err != nil {
		Fail(s.logger, name, arg, errBuffer.String(), err)
		return nil, err
	} else {
		Success(s.logger, name, arg, string(output))
		return output, nil
	}
}

func (s *LocalSession) Exists(path string) (bool, error) {
	return Exists(path), nil
}

func (s *LocalSession) ReadFile(fileName string) ([]byte, error) {
	return os.ReadFile(fileName)
}

func (s *LocalSession) ReadDir(dir string) ([]FileInfo, error) {
	dirs, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	files := make([]FileInfo, len(dirs))
	for i, fileInfo := range dirs {
		files[i] = FileInfo{Name: fileInfo.Name(), Path: fileInfo.Name()}
	}
	return files, nil
}

func (s *LocalSession) MakeDirAll(path string, perm os.FileMode) error {
	return os.MkdirAll(path, perm)
}

func (s *LocalSession) Remove(name string) error {
	return os.Remove(name)
}

func (s *LocalSession) RemoveAll(path string) error {
	return os.RemoveAll(path)
}

func (s *LocalSession) Create(name string) error {
	_, err := os.Create(name)
	return err
}

func (s *LocalSession) WriteString(name string, data string, mode ...string) error {
	flag := os.O_RDWR | os.O_CREATE | os.O_TRUNC
	if len(mode) == 1 && mode[0] == ">>" {
		flag = os.O_RDWR | os.O_CREATE | os.O_APPEND
	}
	fileHandler, err := os.OpenFile(name, flag, os.ModePerm)
	if err != nil {
		return err
	}
	_, err = fileHandler.WriteString(data)
	if err != nil {
		return err
	}
	return nil
}
