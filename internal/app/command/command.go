package command

import (
	_init "github.com/solivaf/go-maria/internal/app/command/init"
	"github.com/solivaf/go-maria/internal/app/command/release/major"
	"github.com/solivaf/go-maria/internal/app/command/release/minor"
	"github.com/solivaf/go-maria/internal/app/command/release/patch"
	"gopkg.in/urfave/cli.v2"
)

func CreateRelease() *cli.Command {
	return &cli.Command{
		Name: "release", Aliases: []string{"r"},
		Usage: "Releases a new version", UsageText: "release [options] [arguments...]",
		Subcommands: createReleaseSubCommands(),
	}
}

func createReleaseSubCommands() []*cli.Command {
	return []*cli.Command{major.Command(), minor.Command(), patch.Command()}
}

func CreateInit() *cli.Command {
	return &cli.Command{
		Name:      "init",
		Usage:     "Initialize project with go-maria configuration",
		UsageText: "go-maria init <app-name>",
		Aliases:   []string{"i"},
		Action:    _init.Execute}
}
