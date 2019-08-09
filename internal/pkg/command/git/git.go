package git

import (
	"errors"
	"github.com/solivaf/go-maria/internal/pkg/command"
	"log"
	"os/exec"
	"strings"
)

const (
	releaseMessage              = "[skip ci] [gomaria] - releasing version "
	prepareToNextReleaseMessage = "[skip ci] [gomaria] - preparing for next release "
)

func Push() error {
	message, err := pushCommits()
	if err != nil {
		log.Fatalln(err.Error())
	}
	log.Println(message)
	if lastTag, err := getLastTag(); err == nil {
		message, err := pushTag(lastTag)
		if err != nil {
			log.Fatalln(err)
		}
		log.Println(message)
	}
	return nil
}

func CommitChanges(message string) error {
	log.Println("Committing git changes")
	cmdMessage, err := addUntrackedFiles()
	log.Println(cmdMessage)
	if err != nil {
		log.Fatalln(err.Error())
		return err
	}

	cmdMessage, err = commitLocally(message)
	log.Println(cmdMessage)
	if err != nil {
		log.Fatalln(err.Error())
		return err
	}
	return nil
}

func GetLatestTag() (string, error) {
	cmd := exec.Command("git", "describe", "--abbrev=0", "--tags")
	stdOut, stdErr := command.GetStdOutAndStdErr(cmd)
	if err := cmd.Run(); err != nil {
		return "", errors.New(stdErr.String())
	}

	return stdOut.String(), nil
}

func CreateTag(tagName string) error {
	log.Println("Creating tag " + tagName)
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
	cmd := exec.Command("git", "tag", tagName)
	stdOut, stdErr := command.GetStdOutAndStdErr(cmd)
	if err := cmd.Run(); err != nil {
		return "", errors.New(stdErr.String())
	}
	return stdOut.String(), nil
}

func commitLocally(message string) (string, error) {
	cmd := exec.Command("git", "commit", "-m", message)
	stdOut, stdErr := command.GetStdOutAndStdErr(cmd)
	if err := cmd.Run(); err != nil {
		log.Fatalln(stdErr.String())
		return "", errors.New(stdErr.String())
	}
	return stdOut.String(), nil
}

func addUntrackedFiles() (string, error) {
	cmd := exec.Command("git", "add", ".")
	stdOut, stdErr := command.GetStdOutAndStdErr(cmd)
	if err := cmd.Run(); err != nil {
		return "", errors.New(stdErr.String())
	}
	return stdOut.String(), nil
}

func pushCommits() (string, error) {
	cmd := exec.Command("git", "push", "origin", "master")
	out, stdErr := command.GetStdOutAndStdErr(cmd)
	if err := cmd.Run(); err != nil {
		return "", errors.New(stdErr.String())
	}
	return out.String(), nil
}

func getLastTag() (string, error) {
	cmd := exec.Command("git", "describe", "--tags", "--abbrev=0")
	out, outErr := command.GetStdOutAndStdErr(cmd)
	if err := cmd.Run(); err != nil {
		return "", errors.New(outErr.String())
	}

	return strings.Replace(out.String(), "\n", "", -1), nil
}

func pushTag(tagName string) (string, error) {
	cmd := exec.Command("git", "push", "origin", tagName)
	stdOut, stdErr := command.GetStdOutAndStdErr(cmd)
	if err := cmd.Run(); err != nil {
		return "", errors.New(stdErr.String())
	}

	return stdOut.String(), nil
}
