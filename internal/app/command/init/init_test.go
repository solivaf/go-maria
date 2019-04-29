package init

import (
	"fmt"
	"github.com/pelletier/go-toml"
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"
)

const testData = "/tmp/"

func TestCreateInitFile(t *testing.T) {
	os.Remove("./.goversion.toml")
	file, err := createInitFile(testData)
	defer file.Close()

	assert.NotNil(t, file)
	assert.NoError(t, err)
}

func TestCreateInitFileError(t *testing.T) {
	file, err := createInitFile("/invalidpath/")
	defer file.Close()

	fmt.Println(err.Error())
	assert.Error(t, err)
	assert.Nil(t, file)
}

func TestOpenInitTemplateError(t *testing.T) {
	tomlFile, err := openInitTemplate("/invalidpath/", ".goversion.toml")

	assert.Error(t, err)
	assert.Nil(t, tomlFile)
}

func TestOpenInitTemplate(t *testing.T) {
	absPath, _ := filepath.Abs("../../../../")
	tomlFile, err := openInitTemplate(absPath, "init.toml")

	assert.NotNil(t, tomlFile)

	module := tomlFile.Get("module").(*toml.Tree)
	assert.NotNil(t, module)

	assert.Equal(t, "init.toml", module.Get("name").(string))
	assert.Equal(t, "v0.0.1-SNAPSHOT", module.Get("version").(string))
	assert.NoError(t, err)
}
