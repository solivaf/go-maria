package command

import (
	_init "github.com/solivaf/go-maria/internal/app/command/init"
	"github.com/solivaf/go-maria/internal/app/command/release"
	"github.com/solivaf/go-maria/internal/app/config"
)

type Command interface {
	Execute() error
}

func CreateCommand(config config.Config) (Command, error) {
	if config.ReleaseConfig != nil {
		return &release.Command{Config: config.ReleaseConfig}, nil
	}
	if config.InitConfig != nil {
		return &_init.Command{Config: config.InitConfig}, nil
	}

	return nil, nil
}
