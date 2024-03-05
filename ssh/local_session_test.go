package ssh

import (
	"github.com/candbright/go-ssh/ssh/options"
	"testing"
)

func TestLocalSession_Exists(t *testing.T) {
	s, err := NewSession(options.LocalHost(), options.Single(false))
	if err != nil {
		t.Fatal(err)
	}
	exist, err := s.Exists("D:\\test\\test.txt")
	if err != nil {
		t.Fatal(err)
	}
	if exist {
		t.Log("true")
	} else {
		t.Log("false")
	}
}

func TestLocalSession_ReadFile(t *testing.T) {
	s, err := NewSession(options.LocalHost(), options.Single(false))
	if err != nil {
		t.Fatal(err)
	}
	output, err := s.ReadFile("C:\\Windows\\System32\\drivers\\etc\\hosts")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(output))
}

func TestLocalSession_ReadDir(t *testing.T) {
	s, err := NewSession(options.LocalHost(), options.Single(false))
	if err != nil {
		t.Fatal(err)
	}
	infos, err := s.ReadDir("C:\\Windows\\System32\\drivers\\etc")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(infos)
}

func TestLocalSession_MakeDirAll(t *testing.T) {
	s, err := NewSession(options.LocalHost(), options.Single(false))
	if err != nil {
		t.Fatal(err)
	}
	err = s.MakeDirAll("D:\\AAA\\BBB\\CCC", 0644)
	if err != nil {
		t.Fatal(err)
	}
}

func TestLocalSession_Remove(t *testing.T) {
	s, err := NewSession(options.LocalHost(), options.Single(false))
	if err != nil {
		t.Fatal(err)
	}
	err = s.Remove("D:\\AAA\\BBB\\CCC\\A.txt")
	if err != nil {
		t.Fatal(err)
	}
}

func TestLocalSession_RemoveAll(t *testing.T) {
	s, err := NewSession(options.LocalHost(), options.Single(false))
	if err != nil {
		t.Fatal(err)
	}
	err = s.RemoveAll("D:\\AAA")
	if err != nil {
		t.Fatal(err)
	}
}

func TestLocalSession_Create(t *testing.T) {
	s, err := NewSession(options.LocalHost(), options.Single(false))
	if err != nil {
		t.Fatal(err)
	}
	err = s.Create("D:\\AAA\\BBB\\CCC\\A.txt")
	if err != nil {
		t.Fatal(err)
	}
}

func TestLocalSession_WriteString(t *testing.T) {
	s, err := NewSession(options.LocalHost(), options.Single(false))
	if err != nil {
		t.Fatal(err)
	}
	err = s.WriteString("D:\\AAA\\BBB\\CCC\\A.txt", "write data")
	if err != nil {
		t.Fatal(err)
	}
}
