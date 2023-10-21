package ssh

import (
	"fmt"
	"io"
	"time"
)

type LocalSession struct {
	writer io.Writer
}

func (s *LocalSession) success(name string, arg []string, output string) {
	if s.writer != nil {
		cmdLog := &CmdLog{
			Cmd:    Command(name, arg...),
			Output: output,
			Time:   time.Now().Format("2006-01-02 15:04:05"),
		}
		_, _ = s.writer.Write([]byte(fmt.Sprint(cmdLog.ToJson(), "\n")))
	}
}

func (s *LocalSession) fail(name string, arg []string, output string, err error) {
	if s.writer != nil {
		cmdLog := &CmdLog{
			Cmd:       Command(name, arg...),
			ErrString: err.Error(),
			Output:    output,
			Time:      time.Now().Format("2006-01-02 15:04:05"),
		}
		_, _ = s.writer.Write([]byte(fmt.Sprint(cmdLog.ToJson(), "\n")))
	}
}
