package init

import (
	"errors"
	"fmt"
	"github.com/pelletier/go-toml"
	"github.com/solivaf/go-maria/internal/pkg/file"
	"gopkg.in/urfave/cli.v2"
	"log"
	"os"
)

func Execute(c *cli.Context) error {
	if c.Args().First() == "" {
		return errors.New("Missing module name")
	}
	absPath := file.GetAbsolutePath()
	initialFile, err := createInitFile(absPath)
	if err != nil {
		return err
	}

	appName := c.Args().First()
	templateFile, err := openInitTemplate(appName)
	if err != nil {
		return err
	}
	writeContent(templateFile, initialFile)

	return nil
}

func writeContent(source *toml.Tree, destination *os.File) {
	if _, err := source.WriteTo(destination); err != nil {
		log.Fatal(err)
	}
}

func openInitTemplate(appName string) (*toml.Tree, error) {
	templatePath := file.GetTemplatesPath()
	tomlFile, err := toml.LoadFile(templatePath + "init.toml")
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	module := tomlFile.Get("module").(*toml.Tree)
	module.Set("name", appName)
	tomlFile.Set("module", module)

	return tomlFile, nil
}

func createInitFile(path string) (*os.File, error) {
	return file.CreateInitialFile(path)
}
