package command_test

import (
	"github.com/solivaf/go-maria/internal/app/command"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateInit(t *testing.T) {
	initCmd := command.CreateInit()

	assert.Equal(t, "init", initCmd.Name)
	assert.Equal(t, "go-maria init <app-name>", initCmd.UsageText)
	assert.NotNil(t, initCmd.Action)
	assert.Zero(t, len(initCmd.Subcommands))
}

func TestCreateRelease(t *testing.T) {
	releaseCmd := command.CreateRelease()

	assert.Equal(t, 3, len(releaseCmd.Subcommands))
	assert.Equal(t, "release", releaseCmd.Name)
	assert.Nil(t, releaseCmd.Action)
	assert.Equal(t, "major", releaseCmd.Subcommands[0].Name)
	assert.NotNil(t, releaseCmd.Subcommands[0].Action)
	assert.Equal(t, "minor", releaseCmd.Subcommands[1].Name)
	assert.NotNil(t, releaseCmd.Subcommands[1].Action)
	assert.Equal(t, "patch", releaseCmd.Subcommands[2].Name)
	assert.NotNil(t, releaseCmd.Subcommands[2].Action)
}
