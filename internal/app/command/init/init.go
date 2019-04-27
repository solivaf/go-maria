package init

import (
	"errors"
	"fmt"
	"github.com/pelletier/go-toml"
	"github.com/solivaf/go-maria/internal/app/file"
	"gopkg.in/urfave/cli.v2"
	"log"
	"os"
)

func Execute(c *cli.Context) error {
	if c.Args().First() == "" {
		return errors.New("Missing module name")
	}
	absPath := file.GetAbsolutePath()
	initialFile := createInitFile(absPath)

	appName := c.Args().First()
	templateFile := openInitTemplate(absPath, appName)
	writeContent(templateFile, initialFile)

	return nil
}

func writeContent(source *toml.Tree, destination *os.File) {
	if _, err := source.WriteTo(destination); err != nil {
		log.Fatal(err)
	}
}

func openInitTemplate(path, appName string) *toml.Tree {
	tomlFile, err := toml.LoadFile(path + "/templates/init.toml")
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	module := tomlFile.Get("module").(*toml.Tree)
	module.Set("name", appName)
	tomlFile.Set("module", module)

	return tomlFile
}

func createInitFile(path string) *os.File {
	return file.CreateInitialFile(path)
}
