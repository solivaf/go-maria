package file

import (
	"errors"
	"fmt"
	"github.com/pelletier/go-toml"
	"os"
	"path/filepath"
	"runtime"
)

const (
	fileName                = ".goversion.toml"
	DockerKey               = "docker"
	DockerBuildDirectoryKey = "buildDirectory"
	DockerComposeKey        = "dockerCompose"
	DockerOrganizationKey   = "organization"
	DockerImageNameKey      = "imageName"
	DockerTagPrefixKey      = "tagPrefix"
	ModuleKey               = "module"
	ModuleVersionKey        = "version"
)

func OpenFile(path string) *os.File {
	f, err := os.OpenFile(getFilePath(path), os.O_WRONLY|os.O_TRUNC, os.ModeAppend)
	if err != nil {
		panic(err)
	}

	return f
}

func LoadTomlFile(path string) *toml.Tree {
	f, err := toml.LoadFile(getFilePath(path))
	if err != nil {
		panic(err)
	}

	return f
}

func WriteFile(path, content string) error {
	f, err := os.Create(getFilePath(path))
	if err != nil {
		return err
	}

	_, err = f.WriteString(content)
	if err != nil {
		return err
	}
	return nil
}

func CreateInitialFile(path string) (*os.File, error) {
	if _, err := os.Open(fileName); err == nil {
		return nil, errors.New("file go.version.toml already exists in path")
	}
	f, err := os.Create(path + "/" + fileName)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Error %s creating file go.version.toml on path %s", err.Error(), path))
	}

	return f, nil
}

func GetAbsolutePath() string {
	p := os.Args[0]
	absPath, _ := filepath.Abs(filepath.Dir(p))
	return absPath
}

func GetTemplatesPath() string {
	_, f, _, _ := runtime.Caller(2)
	p, _ := filepath.Abs(filepath.Dir(f))
	p += "/../../../../templates/"
	return p
}

func GetVersionFromTomlFile(tomlFile *toml.Tree) string {
	module := tomlFile.Get(ModuleKey).(*toml.Tree)
	v := module.Get(ModuleVersionKey).(string)
	return v
}

func getFilePath(path string) string {
	return path + "/" + fileName
}
