package ssh

import (
	"github.com/candbright/go-ssh/options"
	"testing"
)

func TestSingleSession_Exists(t *testing.T) {
	singleSession, err := NewSession(options.EnableSingle())
	if err != nil {
		t.Fatal(err)
	}
	exists, err := singleSession.Exists("D:\\test\\test.txt")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(exists)
}
