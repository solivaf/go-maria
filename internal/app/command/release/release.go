package release

import (
	"github.com/pelletier/go-toml"
	"github.com/solivaf/go-maria/internal/app/file"
	"github.com/solivaf/go-maria/internal/app/git"
	"gopkg.in/urfave/cli.v2"
	"strconv"
	"strings"
)

const (
	majorIndex     = 0
	minorIndex     = 1
	patchIndex     = 2
	snapshotSuffix = "-SNAPSHOT"
	versionPrefix  = "v"
)

func MajorVersion(skipPush bool) error {
	return newVersion(majorIndex, skipPush)
}

func MinorVersion(skipPush bool) error {
	return newVersion(minorIndex, skipPush)
}

func PatchVersion(skipPush bool) error {
	return newVersion(patchIndex, skipPush)
}

func SkipPush(c *cli.Context) bool {
	return c.NArg() > 0 && c.Args().First() == "skip-push"
}

func newVersion(versionIndex int, skipPush bool) error {
	tomlFile := file.LoadTomlFile(file.GetAbsolutePath())
	if err := version(tomlFile, versionIndex); err != nil {
		return err
	}

	if err := prepareToNextRelease(tomlFile, versionIndex); err != nil {
		return err
	}

	if !skipPush {
		return git.Push()
	}
	return nil
}

func version(tomlFile *toml.Tree, index int) error {
	versionReleased := getUpdatedVersionFromTomlFile(tomlFile, index, false)
	if err := updateTomlFileVersion(tomlFile, versionReleased); err != nil {
		return err
	}
	commitMessage := git.ReleaseVersionCommitMessage(versionReleased)
	if err := git.CommitChanges(commitMessage); err != nil {
		return err
	}
	if err := git.CreateTag(versionReleased); err != nil {
		return err
	}
	return nil
}

func prepareToNextRelease(tomlFile *toml.Tree, index int) error {
	versionReleased := getUpdatedVersionFromTomlFile(tomlFile, index, true)
	if err := updateTomlFileVersion(tomlFile, versionReleased); err != nil {
		return err
	}

	commitMessage := git.PrepareVersionToNextReleaseMessage(versionReleased)
	return git.CommitChanges(commitMessage)
}

func updateTomlFileVersion(tomlFile *toml.Tree, version string) error {
	module := tomlFile.Get("module").(*toml.Tree)
	module.Set("version", version)
	tomlFile.Set("module", module)

	f := file.OpenFile(file.GetAbsolutePath())
	if _, err := tomlFile.WriteTo(f); err != nil {
		return err
	}

	return nil
}

func getUpdatedVersionFromTomlFile(tomlFile *toml.Tree, index int, isSnapshot bool) string {
	version := file.GetVersionFromTomlFile(tomlFile)
	updatedVersion := getUpdatedVersionByIndex(version, index, isSnapshot)
	return updatedVersion
}

func getUpdatedVersionByIndex(version string, index int, isSnapshot bool) string {
	versionUpdated := updateVersion(version, index, isSnapshot)
	return versionUpdated
}

func updateVersion(version string, index int, isSnapshot bool) string {
	if !isSnapshot {
		version = strings.Replace(version, snapshotSuffix, "", -1)
		return version
	}
	numbers := getNumbersInVersion(version)
	minorString := numbers[index]
	incrementedVersion := incrementVersion(minorString)
	numbers[index] = incrementedVersion
	setZeroValues(index, numbers)
	return prepareVersionToNextRelease(version, numbers)
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

func prepareVersionToNextRelease(version string, numbersInVersion []string) string {
	versionJoined := strings.Join(numbersInVersion, ".")

	versionJoined += snapshotSuffix
	if strings.HasPrefix(version, versionPrefix) {
		return versionPrefix + versionJoined
	}
	return versionJoined
}

func getNumbersInVersion(version string) []string {
	if strings.HasPrefix(version, versionPrefix) {
		version = strings.Split(version, versionPrefix)[1]
	}
	if strings.HasSuffix(version, snapshotSuffix) {
		version = strings.TrimSuffix(version, snapshotSuffix)
	}
	versionSplitted := strings.Split(version, ".")
	return versionSplitted
}

func incrementVersion(version string) string {
	versionInt, err := strconv.Atoi(version)
	if err != nil {
		panic(err)
	}
	versionInt += 1
	return strconv.Itoa(versionInt)
}
