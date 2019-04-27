package major

import (
	"github.com/solivaf/go-maria/internal/app/command/release"
	"gopkg.in/urfave/cli.v2"
)

func Command() *cli.Command {
	return &cli.Command{Name: "major", Action: execute}
}

func execute(c *cli.Context) error {
	if release.SkipPush(c){
		return release.MajorVersion(true)
	}
	return release.MajorVersion(false)
}
