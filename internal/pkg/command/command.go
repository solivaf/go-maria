package command

import (
	"bytes"
	"os/exec"
)

func GetStdOutAndStdErr(cmd *exec.Cmd) (*bytes.Buffer, *bytes.Buffer) {
	var out, outErr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &outErr
	return &out, &outErr
}
