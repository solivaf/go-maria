package minor

import (
	"github.com/solivaf/go-maria/internal/app/command/release"
	"gopkg.in/urfave/cli.v2"
)

func Command() *cli.Command {
	return &cli.Command{Name: "minor", Action: execute}
}

func execute(c *cli.Context) error {
	if release.SkipPush(c) {
		return release.MinorVersion(true)
	}
	return release.MinorVersion(false)
}