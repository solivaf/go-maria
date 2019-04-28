package docker

import (
	"errors"
	"github.com/pelletier/go-toml"
)

type Config struct {
	Repository          string
	Organization        string
	Directory           string
	Registry            string
	IsDockerComposeFile bool
	ReleaseLatest       bool
}

func CreateConfig(tree *toml.Tree) (*Config, error) {
	if !tree.Has("docker") {
		return nil, errors.New("[docker] not found in .goversion.toml")
	}
	dockerTree := tree.Get("docker").(*toml.Tree)
	if isValid := isValidTree(dockerTree); !isValid {
		return nil, errors.New("missing one of required parameters: organization, repository or builDirectory")
	}
	return &Config{
		Repository:          dockerTree.Get("repository").(string),
		Organization:        dockerTree.Get("organization").(string),
		Directory:           dockerTree.Get("buildDirectory").(string),
		Registry:            dockerTree.Get("registry").(string),
		ReleaseLatest:       dockerTree.Get("releaseLatest").(bool),
		IsDockerComposeFile: dockerTree.Get("dockerComposeFile").(bool),
	}, nil
}

func isValidTree(tree *toml.Tree) bool {
	return tree.Has("repository") && tree.Has("organization") && tree.Has("buildDirectory")
}
