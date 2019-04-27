package major_test

import (
	"github.com/solivaf/go-maria/internal/app/command/release/major"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCommand(t *testing.T) {
	cmd := major.Command()

	assert.Equal(t, "major", cmd.Name)
	assert.NotNil(t, cmd.Action)
}
