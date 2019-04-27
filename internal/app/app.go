package app

import (
	"github.com/solivaf/go-maria/internal/app/module"
	"github.com/solivaf/go-maria/internal/app/vcs"
)

type App struct {
	module module.Module
	vcs    vcs.VCS
}

func (a App) Module() module.Module {
	return a.module
}

func (a App) VCS() vcs.VCS {
	return a.vcs
}
