package main

import (
	"github.com/solivaf/go-maria/internal/app/command"
	"gopkg.in/urfave/cli.v2"
	"log"
	"os"
)

func main() {
	gomaria := &cli.App{}

	gomaria.Name = "go-maria"
	gomaria.Usage = "Made easy releasing versions in Go"
	gomaria.UsageText = "go-maria [command] [options] [arguments..]"
	gomaria.HideVersion = true
	gomaria.Commands = []*cli.Command{command.CreateInit(), command.CreateRelease()}

	if err := gomaria.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
