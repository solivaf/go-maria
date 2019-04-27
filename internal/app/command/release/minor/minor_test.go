package minor_test

import (
	"github.com/solivaf/go-maria/internal/app/command/release/minor"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCommand(t *testing.T) {
	cmd := minor.Command()

	assert.Equal(t, "minor", cmd.Name)
	assert.NotNil(t, cmd.Action)
}
