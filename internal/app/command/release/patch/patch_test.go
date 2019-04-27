package patch_test

import (
	"github.com/solivaf/go-maria/internal/app/command/release/patch"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCommand(t *testing.T) {
	cmd := patch.Command()

	assert.Equal(t, "patch", cmd.Name)
	assert.NotNil(t, cmd.Action)
}
