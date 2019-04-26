package init

import (
	"fmt"
	"github.com/solivaf/go-maria/internal/app/file"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Config struct {
	AppName *string
}

type Command struct {
	*Config
}

func (c *Command) Execute() error {
	localPath := os.Args[0]
	absPath, _ := filepath.Abs(filepath.Dir(localPath))
	file := c.createInitFile(absPath)
	templateFile := c.openInitTemplate(absPath)

	c.WriteContent(templateFile, file)

	return nil
}

func (c *Command) WriteContent(source, destination *os.File) {
	contentBytes, err := ioutil.ReadAll(source)
	if err != nil {
		panic(err)
	}

	destination.Write(contentBytes)
}


func (c *Command) openInitTemplate(path string) *os.File {
	templateFile, err := os.Open(path + "/templates/init.tmpl")
	if err != nil {
		fmt.Println(err.Error())
	}

	return templateFile
}

func (c *Command) createInitFile(path string) *os.File {
	return file.CreateInitialFile(path)
}
