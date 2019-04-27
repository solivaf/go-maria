package patch

import (
	"github.com/solivaf/go-maria/internal/app/command/release"
	"gopkg.in/urfave/cli.v2"
)

func Command() *cli.Command {
	return &cli.Command{Name: "patch", Action: execute}
}

func execute(c *cli.Context) error {
	if release.SkipPush(c) {
		return release.PatchVersion(true)
	}
	return release.PatchVersion(false)
}
