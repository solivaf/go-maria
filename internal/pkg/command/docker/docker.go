package docker

import (
	"errors"
	"github.com/pelletier/go-toml"
	"github.com/solivaf/go-maria/internal/pkg/command"
	"github.com/solivaf/go-maria/internal/pkg/command/git"
	"os/exec"
)

type DockerService struct {
	dockerTree *toml.Tree
}

type Docker interface {
	ReleaseImage() (string, error)
	TagVersion(imageId, gitTag string) (string, error)
}

func (docker *DockerService) ReleaseNewImage() (string, error) {
	imageId, err := docker.buildImage()
	if err != nil {
		return imageId, err
	}

	latestGitTag, err := git.GetLatestTag()
	if err != nil {
		return latestGitTag, err
	}

	tagName, err := docker.TagVersion(imageId, latestGitTag)
	if err != nil {
		return tagName, err
	}

	msg, err := docker.PushTag(tagName)

	return msg, err
}

func (docker *DockerService) buildImage() (imageId string, err error) {
	buildDirectory := docker.dockerTree.Get("buildDirectory").(string)
	isDockerCompose := docker.dockerTree.Get("dockerComposeFile").(bool)

	var dockerCommand string
	if isDockerCompose {
		dockerCommand = "docker-compose"
	} else {
		dockerCommand = "docker"
	}
	cmd := exec.Command(dockerCommand, "build", buildDirectory)
	stdOut, stdErr := command.GetStdOutAndStdErr(cmd)

	if err := cmd.Run(); err != nil {
		return "", errors.New(stdErr.String())
	}

	return stdOut.String(), nil
}

func (docker *DockerService) TagVersion(imageId, latestGitTag string) (string, error) {
	tagName := docker.buildTagName(latestGitTag)

	cmd := exec.Command("docker", "tag", imageId, tagName)
	stdOut, stdErr := command.GetStdOutAndStdErr(cmd)
	if err := cmd.Run(); err != nil {
		return "", errors.New(stdErr.String())
	}

	return stdOut.String(), nil
}

func (docker *DockerService) PushTag(tagName string) (string, error) {
	cmd := exec.Command("docker", "push", tagName)
	stdOut, stdErr := command.GetStdOutAndStdErr(cmd)
	if err := cmd.Run(); err != nil {
		return "", errors.New(stdErr.String())
	}
	return stdOut.String(), nil
}

func (docker *DockerService) buildTagName(latestGitTag string) string {
	organization := docker.dockerTree.Get("organization").(string)
	repository := docker.dockerTree.Get("repository").(string)
	if docker.dockerTree.Has("tagPrefix") {
		prefix := docker.dockerTree.Get("tagPrefix").(string)
		return organization + "/" + repository + ":" + prefix + "-" + latestGitTag
	}
	return organization + "/" + repository + ":" + latestGitTag
}
