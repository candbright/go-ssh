package ssh

import (
	"os"
	"sync"
)

type SingleSession struct {
	lock *sync.Mutex
	Session
}

func (s SingleSession) Run(name string, arg ...string) error {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.Session.Run(name, arg...)
}

func (s SingleSession) Output(name string, arg ...string) ([]byte, error) {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.Session.Output(name, arg...)
}

func (s SingleSession) CombinedOutput(name string, arg ...string) ([]byte, error) {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.Session.CombinedOutput(name, arg...)
}

func (s SingleSession) OutputGrep(cmdList []Cmd) ([]byte, error) {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.Session.OutputGrep(cmdList)
}

func (s SingleSession) Exists(path string) (bool, error) {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.Session.Exists(path)
}

func (s SingleSession) ReadFile(fileName string) ([]byte, error) {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.Session.ReadFile(fileName)
}

func (s SingleSession) ReadDir(dir string) ([]FileInfo, error) {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.Session.ReadDir(dir)
}

func (s SingleSession) MakeDirAll(path string, perm os.FileMode) error {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.Session.MakeDirAll(path, perm)
}

func (s SingleSession) Remove(name string) error {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.Session.Remove(name)
}

func (s SingleSession) RemoveAll(path string) error {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.Session.RemoveAll(path)
}

func (s SingleSession) Create(name string) error {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.Session.Create(name)
}

func (s SingleSession) WriteString(name string, data string, mode ...string) error {
	s.lock.Lock()
	defer s.lock.Unlock()
	return s.Session.WriteString(name, data, mode...)
}
