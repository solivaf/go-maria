package git

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"
	"strings"
	"sync"
)

const (
	releaseMessage              = "[gomaria] - releasing version "
	prepareToNextReleaseMessage = "[gomaria] - preparing for next release "
)

func Push() error {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		if _, err := pushCommits(); err != nil {
			fmt.Println(err.Error())
		}
	}()

	go func() {
		defer wg.Done()
		if lastTag, err := getLastTag(); err == nil {
			if _, err := pushTag(lastTag); err != nil {
				fmt.Println(err.Error())
			}
		}
	}()

	wg.Wait()
	return nil
}

func CommitChanges(message string) error {
	if _, err := addUntrackedFiles(); err != nil {
		return err
	}
	if _, err := commitLocally(message); err != nil {
		return err
	}
	return nil
}

func CreateTag(tagName string) error {
	if _, err := createTag(tagName); err != nil {
		return err
	}
	return nil
}

func ReleaseVersionCommitMessage(version string) string {
	return releaseMessage + version

}

func PrepareVersionToNextReleaseMessage(version string) string {
	return prepareToNextReleaseMessage + version
}

func createTag(tagName string) (string, error) {
	command := exec.Command("git", "tag", tagName)
	stdOut, stdErr := getStdOutAndStdErr(command)
	if err := command.Run(); err != nil {
		return "", errors.New(stdErr.String())
	}
	return stdOut.String(), nil
}

func commitLocally(message string) (string, error) {
	command := exec.Command("git", "commit", "-m", message)
	stdOut, stdErr := getStdOutAndStdErr(command)
	if err := command.Run(); err != nil {
		return "", errors.New(stdErr.String())
	}
	return stdOut.String(), nil
}

func addUntrackedFiles() (string, error) {
	command := exec.Command("git", "add", ".")
	stdOut, stdErr := getStdOutAndStdErr(command)
	if err := command.Run(); err != nil {
		return "", errors.New(stdErr.String())
	}
	return stdOut.String(), nil
}

func pushCommits() (string, error) {
	command := exec.Command("git", "push", "origin", "master")
	out, stdErr := getStdOutAndStdErr(command)
	if err := command.Run(); err != nil {
		return "", errors.New(stdErr.String())
	}
	return out.String(), nil
}

func getLastTag() (string, error) {
	cmd := exec.Command("git", "describe", "--tags", "--abbrev=0")
	out, outErr := getStdOutAndStdErr(cmd)
	if err := cmd.Run(); err != nil {
		return "", errors.New(outErr.String())
	}

	return strings.Replace(out.String(), "\n", "", -1), nil
}

func getStdOutAndStdErr(cmd *exec.Cmd) (bytes.Buffer, bytes.Buffer) {
	var out, outErr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &outErr
	return out, outErr
}

func pushTag(tagName string) (string, error) {
	command := exec.Command("git", "push", "origin", tagName)
	stdOut, stdErr := getStdOutAndStdErr(command)
	if err := command.Run(); err != nil {
		return "", errors.New(stdErr.String())
	}

	return stdOut.String(), nil
}
