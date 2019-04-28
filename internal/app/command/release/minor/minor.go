package minor

import (
	"github.com/solivaf/go-maria/internal/app/command/release"
	"github.com/solivaf/go-maria/internal/pkg/file"
	"gopkg.in/urfave/cli.v2"
)

func Command() *cli.Command {
	return &cli.Command{Name: "minor", Action: execute}
}

func execute(c *cli.Context) error {
	tomlFile := file.LoadTomlFile(file.GetAbsolutePath())
	r := release.CreateRelease(tomlFile)
	if r.SkipPush(c) {
		return r.ReleaseMinor(false)
	}

	return r.ReleaseMinor(true)
}
