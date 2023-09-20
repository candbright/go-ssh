package ssh

import (
	"github.com/candbright/go-ssh/ssh/options"
	"testing"
)

func TestRemoteSession_Exists(t *testing.T) {
	s, err := NewSession(options.LocalHost(), options.Single(false))
	if err != nil {
		t.Fatal(err)
	}
	output, err := s.Output("cmd", "/c", "dir", "/b", "D:\\test\\test.txt")
	if err != nil {
		t.Fatal(err)
	}
	if len(output) != 0 {
		t.Log("true")
	} else {
		t.Log("false")
	}
}
