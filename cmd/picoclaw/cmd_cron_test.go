package main

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewCronCommand(t *testing.T) {
	cmd := newCronCmd()

	require.NotNil(t, cmd)

	assert.Equal(t, "Manage scheduled tasks", cmd.Short)

	assert.False(t, cmd.HasFlags())

	assert.Nil(t, cmd.Run)
	assert.Nil(t, cmd.RunE)

	assert.NotNil(t, cmd.PersistentPreRunE)
	assert.Nil(t, cmd.PersistentPreRun)
	assert.Nil(t, cmd.PersistentPostRun)

	assert.True(t, cmd.HasSubCommands())

	allowedCommands := map[string]struct{}{
		"list":    {},
		"add":     {},
		"remove":  {},
		"enable":  {},
		"disable": {},
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

func TestNewAddSubcommand(t *testing.T) {
	cmd := newCronAddCmd("")

	require.NotNil(t, cmd)

	assert.Equal(t, "add", cmd.Use)
	assert.Equal(t, "Add a new scheduled job", cmd.Short)

	assert.True(t, cmd.HasFlags())

	assert.NotNil(t, cmd.Flags().Lookup("every"))
	assert.NotNil(t, cmd.Flags().Lookup("cron"))
	assert.NotNil(t, cmd.Flags().Lookup("deliver"))
	assert.NotNil(t, cmd.Flags().Lookup("to"))
	assert.NotNil(t, cmd.Flags().Lookup("channel"))

	nameFlag := cmd.Flags().Lookup("name")
	require.NotNil(t, nameFlag)

	messageFlag := cmd.Flags().Lookup("message")
	require.NotNil(t, messageFlag)

	val, found := nameFlag.Annotations[cobra.BashCompOneRequiredFlag]
	require.True(t, found)
	require.NotEmpty(t, val)
	assert.Equal(t, "true", val[0])

	val, found = messageFlag.Annotations[cobra.BashCompOneRequiredFlag]
	require.True(t, found)
	require.NotEmpty(t, val)
	assert.Equal(t, "true", val[0])
}

func TestNewCronDisableSubcommand(t *testing.T) {
	cmd := newCronDisableCmd("")

	require.NotNil(t, cmd)

	assert.Equal(t, "disable", cmd.Use)
	assert.Equal(t, "Disable a job", cmd.Short)

	assert.True(t, cmd.HasExample())
}

func TestNewCronEnableSubcommand(t *testing.T) {
	cmd := newCronEnableCmd("")

	require.NotNil(t, cmd)

	assert.Equal(t, "enable", cmd.Use)
	assert.Equal(t, "Enable a job", cmd.Short)

	assert.True(t, cmd.HasExample())
}

func TestNewCronListSubcommand(t *testing.T) {
	cmd := newCronListCmd("")

	require.NotNil(t, cmd)

	assert.Equal(t, "List all scheduled jobs", cmd.Short)
}

func TestNewCronRemoveSubcommand(t *testing.T) {
	cmd := newCronRemoveCmd("")

	require.NotNil(t, cmd)

	assert.Equal(t, "Remove a job by ID", cmd.Short)

	assert.True(t, cmd.HasExample())
}
