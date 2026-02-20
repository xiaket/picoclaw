package main

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewAuthCommand(t *testing.T) {
	cmd := newAuthCmd()

	require.NotNil(t, cmd)

	assert.Equal(t, "auth", cmd.Use)
	assert.Equal(t, "Manage authentication (login, logout, status)", cmd.Short)

	assert.Len(t, cmd.Aliases, 0)

	assert.Nil(t, cmd.Run)
	assert.Nil(t, cmd.RunE)

	assert.Nil(t, cmd.PersistentPreRun)
	assert.Nil(t, cmd.PersistentPostRun)

	assert.False(t, cmd.HasFlags())
	assert.True(t, cmd.HasSubCommands())
	assert.True(t, cmd.HasExample())

	allowedCommands := map[string]struct{}{
		"login":  {},
		"logout": {},
		"status": {},
		"models": {},
	}

	subcommands := cmd.Commands()
	assert.Len(t, subcommands, len(allowedCommands))

	for _, subcmd := range subcommands {
		_, found := allowedCommands[subcmd.Name()]
		assert.True(t, found, "unexpected subcommand %q", subcmd.Name())

		assert.Len(t, subcmd.Aliases, 0)
		assert.False(t, subcmd.Hidden)

		assert.False(t, subcmd.HasSubCommands())

		assert.Nil(t, subcmd.Run)
		assert.NotNil(t, subcmd.RunE)

		assert.Nil(t, subcmd.PersistentPreRun)
		assert.Nil(t, subcmd.PersistentPostRun)
	}
}

func TestNewLoginSubCommand(t *testing.T) {
	cmd := newAuthLoginCmd()

	require.NotNil(t, cmd)

	assert.Equal(t, "Login via OAuth or paste token", cmd.Short)

	assert.True(t, cmd.HasFlags())

	assert.NotNil(t, cmd.Flags().Lookup("device-code"))

	providerFlag := cmd.Flags().Lookup("provider")
	require.NotNil(t, providerFlag)

	val, found := providerFlag.Annotations[cobra.BashCompOneRequiredFlag]
	require.True(t, found)
	require.NotEmpty(t, val)
	assert.Equal(t, "true", val[0])
}

func TestNewLogoutSubcommand(t *testing.T) {
	cmd := newAuthLogoutCmd()

	require.NotNil(t, cmd)

	assert.Equal(t, "Remove stored credentials", cmd.Short)

	assert.True(t, cmd.HasFlags())

	assert.NotNil(t, cmd.Flags().Lookup("provider"))
}

func TestNewModelsCommand(t *testing.T) {
	cmd := newAuthModelsCmd()

	require.NotNil(t, cmd)

	assert.Equal(t, "models", cmd.Use)
	assert.Equal(t, "List available Antigravity models", cmd.Short)

	assert.False(t, cmd.HasFlags())
}

func TestNewStatusSubcommand(t *testing.T) {
	cmd := newAuthStatusCmd()

	require.NotNil(t, cmd)

	assert.Equal(t, "Show current auth status", cmd.Short)

	assert.False(t, cmd.HasFlags())
}
