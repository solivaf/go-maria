package git

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
	"sync"
)

type Git struct {
}

func (g *Git) Push() error {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		if err := pushCommits(); err != nil {
			fmt.Println(err.Error())
		}
	}()

	go func() {
		defer wg.Done()
		if lastTag, err := getLastTag(); err == nil {
			if err := pushTag(lastTag); err != nil {
				fmt.Println(err.Error())
			}
		}
	}()

	wg.Wait()
	return nil
}

func (g *Git) CommitChanges(message string) error {
	if err := addUntrackedFiles(); err != nil {
		panic(err)
	}
	if err := commitLocally(message); err != nil {
		panic(err)
	}
	return nil
}

func (g *Git) CreateTag(tagName string) error {
	if err := createTag(tagName); err != nil {
		return err
	}
	return nil
}

func createTag(tagName string) error {
	command := exec.Command("git", "tag", tagName)
	return command.Run()
}

func commitLocally(message string) error {
	command := exec.Command("git", "commit", "-m", message)
	return command.Run()
}

func addUntrackedFiles() error {
	command := exec.Command("git", "add", ".")
	return command.Run()
}

func pushCommits() error {
	command := exec.Command("git", "push", "origin", "master")
	return command.Run()
}

func getLastTag() (string, error) {
	cmd := exec.Command("git", "describe", "--tags", "--abbrev=0")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()

	return strings.Replace(out.String(), "\n", "", -1), err
}

func pushTag(tagName string) error {
	command := exec.Command("git", "push", "origin", tagName)
	var out bytes.Buffer
	command.Stderr = &out
	return command.Run()
}
