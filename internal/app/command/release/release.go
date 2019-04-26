package release

import (
	"github.com/pelletier/go-toml"
	"github.com/solivaf/go-maria/internal/app/file"
	_git "github.com/solivaf/go-maria/internal/app/git"
	"strconv"
	"strings"
)

const majorIndex = 0
const minorIndex = 1
const patchIndex = 2

type Config struct {
	Patch       *bool
	Minor       *bool
	Major       *bool
	SkipPublish *bool
}

const (
	releaseMessage              = "[gomaria] - releasing version "
	prepareToNextReleaseMessage = "[gomaria] - preparing for next release "
)

type Command struct {
	*Config
	_git.Git
}

func (c *Command) Execute() error {
	tomlFile := file.LoadTomlFile(file.GetAbsolutePath())
	if err := c.releaseVersion(tomlFile); err != nil {
		return err
	}
	if err := c.prepareFileToNextRelease(tomlFile); err != nil {
		return err
	}
	if !*c.Config.SkipPublish {
		return c.Git.Push()
	}

	return nil
}

func (c *Command) releaseVersion(tomlFile *toml.Tree) error {
	updatedVersion := c.getUpdatedVersion(tomlFile, false)
	if err := c.updateTomlFileVersion(tomlFile, updatedVersion); err != nil {
		return err
	}

	commitMessage := releaseMessage + updatedVersion
	if err := c.Git.CommitChanges(commitMessage); err != nil {
		return err
	}
	if err := c.Git.CreateTag(updatedVersion); err != nil {
		return err
	}
	return nil
}

func (c *Command) prepareFileToNextRelease(tomlFile *toml.Tree) error {
	updatedVersion := c.getUpdatedVersion(tomlFile, true)
	if err := c.updateTomlFileVersion(tomlFile, updatedVersion); err != nil {
		return err
	}

	commitMessage := prepareToNextReleaseMessage + updatedVersion
	return c.CommitChanges(commitMessage)
}

func (c *Command) updateTomlFileVersion(tomlFile *toml.Tree, version string) error {
	module := tomlFile.Get("module").(*toml.Tree)
	module.Set("version", version)
	tomlFile.Set("module", module)

	f := file.OpenFile(file.GetAbsolutePath())
	if _, err := tomlFile.WriteTo(f); err != nil {
		return err
	}

	return nil
}

func (c *Command) getAppName(tree *toml.Tree) string {
	scm := tree.Get("module").(*toml.Tree)
	return scm.Get("name").(string)
}

func (c *Command) getUpdatedVersion(tree *toml.Tree, isSnapshot bool) string {
	module := tree.Get("module").(*toml.Tree)
	v := module.Get("version").(string)

	updatedVersion := c.getUpdatedVersionByArg(v, isSnapshot)
	return updatedVersion
}

func (c *Command) getUpdatedVersionByArg(version string, isSnapshot bool) string {
	if *c.Major {
		versionUpdated := updateVersion(version, majorIndex, isSnapshot)
		return versionUpdated
	}
	if *c.Minor {
		versionUpdated := updateVersion(version, minorIndex, isSnapshot)
		return versionUpdated
	}
	if *c.Patch {
		versionUpdated := updateVersion(version, patchIndex, isSnapshot)
		return versionUpdated
	}
	return version
}

func updateVersion(version string, index int, isSnapshot bool) string {
	if !isSnapshot {
		version = strings.Replace(version, "-SNAPSHOT", "", -1)
		return version
	}
	numbers := getNumbersInVersion(version)
	minorString := numbers[index]
	incrementedVersion := incrementVersion(minorString)
	numbers[index] = incrementedVersion
	setZeroValues(index, numbers)
	return getUpdatedVersion(version, numbers)
}

func setZeroValues(index int, numbers []string) {
	if index == majorIndex {
		numbers[patchIndex] = "0"
		numbers[minorIndex] = "0"
	}
	if index == minorIndex {
		numbers[patchIndex] = "0"
	}
}

func getUpdatedVersion(version string, numbersInVersion []string) string {
	versionJoined := strings.Join(numbersInVersion, ".")
	versionJoined += "-SNAPSHOT"
	if strings.HasPrefix(version, "v") {
		return "v" + versionJoined
	}
	return versionJoined
}

func getNumbersInVersion(version string) []string {
	if strings.HasPrefix(version, "v") {
		version = strings.Split(version, "v")[1]
	}
	if strings.HasSuffix(version, "-SNAPSHOT") {
		version = strings.TrimSuffix(version, "-SNAPSHOT")
	}
	versionSplitted := strings.Split(version, ".")
	return versionSplitted
}

func incrementVersion(version string) string {
	major, err := strconv.Atoi(version)
	if err != nil {
		panic(err)
	}
	major += 1
	return strconv.Itoa(major)
}
