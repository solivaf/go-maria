package release

import (
	"github.com/pelletier/go-toml"
	"github.com/solivaf/go-maria/internal/pkg/command/docker"
	"github.com/solivaf/go-maria/internal/pkg/command/git"
	"github.com/solivaf/go-maria/internal/pkg/file"
	"gopkg.in/urfave/cli.v2"
	"strconv"
	"strings"
	"sync"
)

const (
	majorIndex     = 0
	minorIndex     = 1
	patchIndex     = 2
	snapshotSuffix = "-SNAPSHOT"
	versionPrefix  = "v"
	zeroVersion    = "0"
)

type ReleaseService struct {
	docker.Docker
	FileTree *toml.Tree
}

type Release interface {
	SkipPush(context *cli.Context) bool
	ReleaseMajor(push bool) error
	ReleaseMinor(push bool) error
	ReleasePatch(push bool) error
}

func CreateRelease(tree *toml.Tree) Release {
	r := &ReleaseService{FileTree: tree}
	if tree.Has(file.DockerKey) {
		r.Docker = &docker.DockerService{Tree: tree.Get(file.DockerKey).(*toml.Tree)}
	}
	return r
}

func (r ReleaseService) SkipPush(context *cli.Context) bool {
	return context.NArg() > 0 && context.Args().First() == "skip-push"
}

func (r *ReleaseService) ReleaseMajor(push bool) error {
	return r.newVersion(majorIndex, push)
}

func (r *ReleaseService) ReleaseMinor(push bool) error {
	return r.newVersion(minorIndex, push)
}

func (r *ReleaseService) ReleasePatch(push bool) error {
	return r.newVersion(patchIndex, push)
}

func (r *ReleaseService) newVersion(versionIndex int, push bool) error {
	if err := r.version(versionIndex); err != nil {
		return err
	}

	if err := r.prepareToNextRelease(versionIndex); err != nil {
		return err
	}

	if push {
		err := r.pushVersion()
		return err
	}
	return nil
}

func (r *ReleaseService) pushVersion() error {
	var wg sync.WaitGroup
	var err error
	wg.Add(1)
	go pushGit(&wg, err)
	if r.Docker != nil {
		wg.Add(1)
		go r.pushDocker(&wg, err)
	}
	wg.Wait()
	return err
}

func (r *ReleaseService) pushDocker(wg *sync.WaitGroup, err error) {
	defer wg.Done()
	if _, _err := r.Docker.ReleaseNewImage(); _err != nil {
		err = _err
	}
}

func pushGit(wg *sync.WaitGroup, err error) {
	defer wg.Done()
	if _err := git.Push(); _err != nil {
		err = _err
	}
}

func (r *ReleaseService) version(index int) error {
	versionReleased := r.getUpdatedVersionFromTomlFile(index, false)
	if err := r.updateTomlFileVersion(versionReleased); err != nil {
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

func (r *ReleaseService) prepareToNextRelease(index int) error {
	versionReleased := r.getUpdatedVersionFromTomlFile(index, true)
	if err := r.updateTomlFileVersion(versionReleased); err != nil {
		return err
	}

	commitMessage := git.PrepareVersionToNextReleaseMessage(versionReleased)
	return git.CommitChanges(commitMessage)
}

func (r *ReleaseService) updateTomlFileVersion(version string) error {
	module := r.FileTree.Get(file.ModuleKey).(*toml.Tree)
	module.Set(file.ModuleVersionKey, version)
	r.FileTree.Set(file.ModuleKey, module)

	f := file.OpenFile(file.GetAbsolutePath())
	if _, err := r.FileTree.WriteTo(f); err != nil {
		return err
	}

	return nil
}

func (r *ReleaseService) getUpdatedVersionFromTomlFile(index int, isSnapshot bool) string {
	version := file.GetVersionFromTomlFile(r.FileTree)
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
		numbers[patchIndex] = zeroVersion
		numbers[minorIndex] = zeroVersion
	}
	if index == minorIndex {
		numbers[patchIndex] = zeroVersion
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
