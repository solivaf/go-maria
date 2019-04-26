package app

import (
	"go-maria/internal/app/command"
	"go-maria/internal/app/config"
	"go-maria/internal/app/module"
	"go-maria/internal/app/vcs"
)

type App struct {
	module module.Module
	vcs vcs.VCS
}

func ExecuteCommand(config config.Config) error {
	cmd, err := command.CreateCommand(config)
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