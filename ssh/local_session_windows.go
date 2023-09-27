package ssh

import (
	"bytes"
	"errors"
	"golang.org/x/text/encoding/simplifiedchinese"
	"io"
	"os"
	"os/exec"
	"strings"
)

type LocalSession struct {
	writer io.Writer
}

func (s *LocalSession) Run(name string, arg ...string) error {
	_, err := s.Output(name, arg...)
	return err
}

func (s *LocalSession) Output(name string, arg ...string) ([]byte, error) {
	args := make([]string, len(arg)+5)
	args[0] = "/c"
	args[1] = name
	copy(args[2:], arg)
	c := exec.Command("cmd", args...)
	var errBuffer bytes.Buffer
	c.Stderr = &errBuffer
	output, err := c.Output()
	errStr, _ := simplifiedchinese.GBK.NewDecoder().String(errBuffer.String())
	if err != nil {
		Fail(s.writer, name, arg, errStr, err)
		return nil, err
	} else {
		Success(s.writer, name, arg, string(output))
		return output, nil
	}
}

func (s *LocalSession) CombinedOutput(name string, arg ...string) ([]byte, error) {
	args := make([]string, len(arg)+5)
	args[0] = "/c"
	args[1] = name
	copy(args[2:], arg)
	output, err := exec.Command("cmd", args...).CombinedOutput()
	if err != nil {
		Fail(s.writer, name, arg, string(output), err)
		return nil, err
	} else {
		Success(s.writer, name, arg, string(output))
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
	name := "cmd"
	arg := []string{"/c", strings.Join(cmdStrList, " | ")}
	c := exec.Command(name, arg...)
	var errBuffer bytes.Buffer
	c.Stderr = &errBuffer
	output, err := c.Output()
	errStr, _ := simplifiedchinese.GBK.NewDecoder().String(errBuffer.String())
	if err != nil {
		Fail(s.writer, name, arg, errStr, err)
		return nil, err
	} else {
		Success(s.writer, name, arg, string(output))
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
	file, err := os.Create(name)
	defer func() {
		_ = file.Close()
	}()
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
	defer func() {
		_ = fileHandler.Close()
	}()
	_, err = fileHandler.WriteString(data)
	if err != nil {
		return err
	}
	return nil
}
