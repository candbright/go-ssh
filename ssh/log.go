package ssh

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"time"
)

type CmdLog struct {
	Time      string `json:"time,omitempty"`
	Cmd       string `json:"cmd,omitempty"`
	ErrString string `json:"error,omitempty"`
	Output    string `json:"output,omitempty"`
}

func (cmdLog *CmdLog) ToJson() string {
	b, err := json.Marshal(cmdLog)
	if err != nil {
		return fmt.Sprintf("%v", cmdLog)
	}
	var out bytes.Buffer
	err = json.Indent(&out, b, "", "  ")
	if err != nil {
		return fmt.Sprintf("%v", cmdLog)
	}
	return out.String()
}

func Success(writer io.Writer, name string, arg []string, output string) {
	if writer != nil {
		cmdLog := &CmdLog{
			Cmd:    Command(name, arg...),
			Output: output,
			Time:   time.Now().Format("2006-01-02 15:04:05"),
		}
		_, _ = writer.Write([]byte(fmt.Sprint(cmdLog.ToJson(), "\n")))
	}
}

func Fail(writer io.Writer, name string, arg []string, output string, err error) {
	if writer != nil {
		cmdLog := &CmdLog{
			Cmd:       Command(name, arg...),
			ErrString: err.Error(),
			Output:    output,
			Time:      time.Now().Format("2006-01-02 15:04:05"),
		}
		_, _ = writer.Write([]byte(fmt.Sprint(cmdLog.ToJson(), "\n")))
	}
}
