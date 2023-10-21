package ssh

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type CmdLog struct {
	Time      string `json:"time,omitempty"`
	Cmd       string `json:"cmd,omitempty"`
	ErrString string `json:"error,omitempty"`
	Output    string `json:"output,omitempty"`
	Host      string `json:"host,omitempty"`
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
