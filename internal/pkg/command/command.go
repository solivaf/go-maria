package command

import (
	"bytes"
	"errors"
	"log"
	"os/exec"
)

type Commander interface {
	Execute(name string, args ...string) (string, error)
}

type CommanderUnix struct {}

func (c *CommanderUnix) Execute(name string, args ...string) (message string, err error) {
	cmd := exec.Command(name, args...)
	stdOut, stdErr := GetStdOutAndStdErr(cmd)
	if err := cmd.Run(); err != nil {
		log.Print(err.Error())
	}
	return stdOut.String(), errors.New(stdErr.String())
}

func GetStdOutAndStdErr(cmd *exec.Cmd) (*bytes.Buffer, *bytes.Buffer) {
	var out, outErr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &outErr
	return &out, &outErr
}
