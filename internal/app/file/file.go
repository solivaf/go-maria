package file

import (
	"errors"
	"fmt"
	"github.com/pelletier/go-toml"
	"os"
	"path/filepath"
)

const fileName = ".goversion.toml"

func OpenFile(path string) *os.File {
	f, err := os.OpenFile(getFilePath(path), os.O_WRONLY | os.O_TRUNC, os.ModeAppend)
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

func CreateInitialFile(path string) *os.File {
	_, err := os.Open(fileName)
	if err == nil {
		fmt.Println("file go.version.toml already exists in path")
		os.Exit(1)
	}
	f, err := os.Create(fileName)
	if err != nil {
		panic(errors.New(fmt.Sprintf("Error %s creating file go.version.toml on path %s", err.Error(), path)))
	}

	return f
}

func GetAbsolutePath() string {
	p := os.Args[0]
	absPath, _ := filepath.Abs(filepath.Dir(p))
	return absPath
}

func CreateFile(path, content string) error {
	f, err := os.Create(getFilePath(path))
	if err != nil {
		panic(err)
	}
	if _, err := f.WriteString(content); err != nil {
		panic(err)
	}

	return nil
}

func getFilePath(path string) string {
	return path + "/" + fileName
}
