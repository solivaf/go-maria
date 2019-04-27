package release

import (
	"github.com/pelletier/go-toml"
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"
)

func TestIncrementVersion(t *testing.T) {
	versionIncremented := incrementVersion("1")

	assert.Equal(t, "2", versionIncremented)
}

func TestIncrementVersionPanicking(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			assert.Error(t, r.(error))
		}
	}()
	incrementVersion("abc")
}

func TestGetNumbersInVersionWithVersionPrefixAndSnapshotSuffix(t *testing.T) {
	numbersArray := getNumbersInVersion("v1.0.3-SNAPSHOT")

	assert.Equal(t, 3, len(numbersArray))
	assert.Equal(t, "1", numbersArray[0])
	assert.Equal(t, "0", numbersArray[1])
	assert.Equal(t, "3", numbersArray[2])
}

func TestGetNumbersInVersionWithVersionPrefix(t *testing.T) {
	numbersArray := getNumbersInVersion("v1.0.3")

	assert.Equal(t, 3, len(numbersArray))
	assert.Equal(t, "1", numbersArray[0])
	assert.Equal(t, "0", numbersArray[1])
	assert.Equal(t, "3", numbersArray[2])
}

func TestGetNumbersInVersionWithSnapshotSuffix(t *testing.T) {
	numbersArray := getNumbersInVersion("1.0.3-SNAPSHOT")

	assert.Equal(t, 3, len(numbersArray))
	assert.Equal(t, "1", numbersArray[0])
	assert.Equal(t, "0", numbersArray[1])
	assert.Equal(t, "3", numbersArray[2])
}

func TestGetVersionPreparedToNextRelease(t *testing.T) {
	numbers := []string{"1", "0", "5"}
	version := "v1.0.4"

	versionPrepared := prepareVersionToNextRelease(version, numbers)

	assert.Equal(t, "v1.0.5-SNAPSHOT", versionPrepared)
}

func TestGetVersionPreparedToNextReleaseWithoutPefix(t *testing.T) {
	numbers := []string{"1", "0", "5"}
	version := "1.0.4"

	versionPrepared := prepareVersionToNextRelease(version, numbers)

	assert.Equal(t, "1.0.5-SNAPSHOT", versionPrepared)
}

func TestSetZeroValuesByMajorIndex(t *testing.T) {
	numbers := []string{"1", "0", "5"}
	setZeroValues(majorIndex, numbers)

	assert.Equal(t, "1", numbers[0])
	assert.Equal(t, "0", numbers[1])
	assert.Equal(t, "0", numbers[2])
}

func TestSetZeroValuesByMinorIndex(t *testing.T) {
	numbers := []string{"1", "3", "5"}
	setZeroValues(minorIndex, numbers)

	assert.Equal(t, "1", numbers[0])
	assert.Equal(t, "3", numbers[1])
	assert.Equal(t, "0", numbers[2])
}

func TestUpdateVersionWithMajorIndexAndSnapshotFalse(t *testing.T) {
	version := "v1.0.0-SNAPSHOT"
	versionUpdated := updateVersion(version, majorIndex, false)

	assert.Equal(t, "v1.0.0", versionUpdated)
}

func TestUpdateVersionWithMajorIndexAndSnapshotTrue(t *testing.T) {
	version := "v1.0.0-SNAPSHOT"
	versionUpdated := updateVersion(version, majorIndex, true)

	assert.Equal(t, "v2.0.0-SNAPSHOT", versionUpdated)
}

func TestUpdateVersionWithMinorIndexAndSnapshotTrue(t *testing.T) {
	version := "v1.0.0-SNAPSHOT"
	versionUpdated := updateVersion(version, minorIndex, true)

	assert.Equal(t, "v1.1.0-SNAPSHOT", versionUpdated)
}

func TestUpdateVersionWithPatchIndexAndSnapshotTrue(t *testing.T) {
	version := "v1.0.1-SNAPSHOT"
	versionUpdated := updateVersion(version, patchIndex, true)

	assert.Equal(t, "v1.0.2-SNAPSHOT", versionUpdated)
}

func TestGetUpdatedVersionByIndexAndMajorIndex(t *testing.T) {
	version := "v1.0.4-SNAPSHOT"
	updatedVersion := getUpdatedVersionByIndex(version, majorIndex, true)

	assert.Equal(t, "v2.0.0-SNAPSHOT", updatedVersion)
}

func TestGetUpdatedVersionByIndexAndMinorIndex(t *testing.T) {
	version := "v1.0.4-SNAPSHOT"
	updatedVersion := getUpdatedVersionByIndex(version, minorIndex, true)

	assert.Equal(t, "v1.1.0-SNAPSHOT", updatedVersion)
}

func TestGetUpdatedVersionByIndexAndPatchIndex(t *testing.T) {
	version := "v1.0.4-SNAPSHOT"
	updatedVersion := getUpdatedVersionByIndex(version, patchIndex, true)

	assert.Equal(t, "v1.0.5-SNAPSHOT", updatedVersion)
}

func TestGetUpdatedVersionFromTomlFileMajorIndex(t *testing.T) {
	absPath, _ := filepath.Abs("../../../../testdata/.goversion.toml")
	file, _ := toml.LoadFile(absPath)
	fileVersion := file.Get("module").(*toml.Tree).Get("version").(string)
	version := getUpdatedVersionFromTomlFile(file, majorIndex, true)

	assert.NotEqual(t, fileVersion, version)
	assert.Equal(t, "v1.0.0-SNAPSHOT", version)
}

func TestGetUpdatedVersionFromTomlFileMinorIndex(t *testing.T) {
	absPath, _ := filepath.Abs("../../../../testdata/.goversion.toml")
	file, _ := toml.LoadFile(absPath)
	fileVersion := file.Get("module").(*toml.Tree).Get("version").(string)
	version := getUpdatedVersionFromTomlFile(file, minorIndex, true)

	assert.NotEqual(t, fileVersion, version)
	assert.Equal(t, "v0.3.0-SNAPSHOT", version)
}

func TestGetUpdatedVersionFromTomlFilePatchIndex(t *testing.T) {
	absPath, _ := filepath.Abs("../../../../testdata/.goversion.toml")
	file, _ := toml.LoadFile(absPath)
	fileVersion := file.Get("module").(*toml.Tree).Get("version").(string)
	version := getUpdatedVersionFromTomlFile(file, patchIndex, true)

	assert.NotEqual(t, fileVersion, version)
	assert.Equal(t, "v0.2.1-SNAPSHOT", version)
}

func TestUpdateTomFileVersion(t *testing.T) {
	absPath, _ := filepath.Abs("../../../../testdata/.goversion.toml")
	file, _ := toml.LoadFile(absPath)
	os.Args[0] = "../../../../testdata/"

	assert.NoError(t, updateTomlFileVersion(file, "v0.2.1"))
	assert.NoError(t, updateTomlFileVersion(file, "v0.2.0-SNAPSHOT"))
}
