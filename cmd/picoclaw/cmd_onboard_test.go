package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewOnboardCommand(t *testing.T) {
	cmd := newOnboardCmd()

	require.NotNil(t, cmd)

	assert.Equal(t, "onboard", cmd.Use)
	assert.Equal(t, "Initialize picoclaw configuration and workspace", cmd.Short)

	assert.NotNil(t, cmd.RunE)
	assert.Nil(t, cmd.Run)

	assert.Nil(t, cmd.PersistentPreRun)
	assert.Nil(t, cmd.PersistentPostRun)

	assert.False(t, cmd.HasFlags())
	assert.False(t, cmd.HasSubCommands())
}
