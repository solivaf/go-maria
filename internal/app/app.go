package app

import (
	"github.com/solivaf/go-maria/internal/app/command"
	"github.com/solivaf/go-maria/internal/app/config"
	"github.com/solivaf/go-maria/internal/app/module"
	"github.com/solivaf/go-maria/internal/app/vcs"
)

type App struct {
	module module.Module
	vcs vcs.VCS
}

func ExecuteCommand(config config.Config) error {
	cmd, err := command.CreateInit(config)
	if err != nil {
		return err
	}
	if err := cmd.Execute(); err != nil {
		return err
	}
	return nil
}

func (a App) Module() module.Module {
	return a.module
}

func (a App) VCS() vcs.VCS {
	return a.vcs
}