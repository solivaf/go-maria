package docker

import (
	"github.com/pelletier/go-toml"
	"github.com/solivaf/go-maria/internal/pkg/command"
	"github.com/solivaf/go-maria/internal/pkg/command/git"
	"github.com/solivaf/go-maria/internal/pkg/file"
)

const (
	dockerCompose = "docker-compose"
	dockerCommand = "docker"
)

type DockerService struct {
	Tree      *toml.Tree
	Commander command.Commander
}

type Docker interface {
	ReleaseNewImage() (string, error)
	TagVersion(imageId, gitTag string) (string, error)
	PushTag(tagName string) (string, error)
}

func CreateDocker(tree *toml.Tree) Docker {
	return &DockerService{Tree: tree, Commander: &command.CommanderUnix{}}
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
	buildDirectory := docker.Tree.Get(file.DockerBuildDirectoryKey).(string)
	isDockerCompose := docker.Tree.Get(file.DockerComposeKey).(bool)

	commandName := docker.getCommandName(isDockerCompose)
	message, err := docker.Commander.Execute(commandName, "build", buildDirectory)

	return message, err
}

func (docker *DockerService) TagVersion(imageId, latestGitTag string) (string, error) {
	tagName := docker.buildTagName(latestGitTag)
	return docker.Commander.Execute(dockerCommand, "tag", imageId, tagName)
}

func (docker *DockerService) PushTag(tagName string) (string, error) {
	return docker.Commander.Execute(dockerCommand, "push", tagName)
}

func (docker *DockerService) buildTagName(latestGitTag string) string {
	organization := docker.Tree.Get(file.DockerOrganizationKey).(string)
	imageName := docker.Tree.Get(file.DockerImageNameKey).(string)
	return organization + "/" + imageName + ":" + latestGitTag
}

func (docker *DockerService) getCommandName(isDockerCompose bool) string {
	if isDockerCompose {
		return dockerCompose
	}
	return dockerCommand
}
