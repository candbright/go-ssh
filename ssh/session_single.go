package ssh

import (
	"os"
	"sync"
)

type SingleSession struct {
	lock *sync.Mutex
	session
}

func (s SingleSession) Run(name string, arg ...string) error {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.session.Run(name, arg...)
}

func (s SingleSession) Output(name string, arg ...string) ([]byte, error) {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.session.Output(name, arg...)
}

func (s SingleSession) CombinedOutput(name string, arg ...string) ([]byte, error) {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.session.CombinedOutput(name, arg...)
}

func (s SingleSession) OutputGrep(cmdList []Cmd) ([]byte, error) {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.session.OutputGrep(cmdList)
}

func (s SingleSession) Exists(path string) (bool, error) {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.session.Exists(path)
}

func (s SingleSession) ReadFile(fileName string) ([]byte, error) {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.session.ReadFile(fileName)
}

func (s SingleSession) ReadDir(dir string) ([]FileInfo, error) {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.session.ReadDir(dir)
}

func (s SingleSession) MakeDirAll(path string, perm os.FileMode) error {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.session.MakeDirAll(path, perm)
}

func (s SingleSession) Remove(name string) error {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.session.Remove(name)
}

func (s SingleSession) RemoveAll(path string) error {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.session.RemoveAll(path)
}

func (s SingleSession) Create(name string) error {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.session.Create(name)
}

func (s SingleSession) WriteString(name string, data string, mode ...string) error {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.session.WriteString(name, data, mode...)
}
