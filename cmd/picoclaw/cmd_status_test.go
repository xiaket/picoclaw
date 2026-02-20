package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewStatusCommand(t *testing.T) {
	cmd := newStatusCmd()

	require.NotNil(t, cmd)

	assert.Equal(t, "status", cmd.Use)

	assert.Equal(t, "Show picoclaw status", cmd.Short)

	assert.False(t, cmd.HasSubCommands())

	assert.Nil(t, cmd.Run)
	assert.NotNil(t, cmd.RunE)

	assert.Nil(t, cmd.PersistentPreRun)
	assert.Nil(t, cmd.PersistentPostRun)
}
