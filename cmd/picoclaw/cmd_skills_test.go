package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewSkillsCommand(t *testing.T) {
	cmd := newSkillsCmd()

	require.NotNil(t, cmd)

	assert.Equal(t, "skills", cmd.Use)
	assert.Equal(t, "Manage skills (install, list, remove)", cmd.Short)

	assert.Len(t, cmd.Aliases, 0)

	assert.False(t, cmd.HasFlags())

	assert.Nil(t, cmd.Run)
	assert.Nil(t, cmd.RunE)

	assert.Nil(t, cmd.PersistentPreRunE)
	assert.Nil(t, cmd.PersistentPreRun)
	assert.Nil(t, cmd.PersistentPostRun)

	assert.True(t, cmd.HasSubCommands())
	assert.True(t, cmd.HasExample())

	allowedCommands := map[string]struct{}{
		"list":            {},
		"install":         {},
		"remove":          {},
		"search":          {},
		"show":            {},
		"list-builtin":    {},
		"install-builtin": {},
	}

	subcommands := cmd.Commands()
	assert.Len(t, subcommands, len(allowedCommands))

	for _, subcmd := range subcommands {
		_, found := allowedCommands[subcmd.Name()]
		assert.True(t, found, "unexpected subcommand %q", subcmd.Name())

		assert.False(t, subcmd.Hidden)
	}
}

func TestNewInstallSubcommand(t *testing.T) {
	cmd := newSkillsInstallCmd()

	require.NotNil(t, cmd)

	assert.Equal(t, "install", cmd.Use)
	assert.Equal(t, "Install skill from GitHub", cmd.Short)

	assert.Nil(t, cmd.Run)
	assert.NotNil(t, cmd.RunE)

	assert.True(t, cmd.HasExample())
	assert.False(t, cmd.HasSubCommands())

	assert.True(t, cmd.HasFlags())
	assert.NotNil(t, cmd.Flags().Lookup("registry"))

	assert.Len(t, cmd.Aliases, 0)
}

func TestNewInstallbuiltinSubcommand(t *testing.T) {
	cmd := newSkillsInstallBuiltinCmd()

	require.NotNil(t, cmd)

	assert.Equal(t, "install-builtin", cmd.Use)
	assert.Equal(t, "Install all builtin skills to workspace", cmd.Short)

	assert.NotNil(t, cmd.RunE)

	assert.True(t, cmd.HasExample())
	assert.False(t, cmd.HasSubCommands())

	assert.False(t, cmd.HasFlags())

	assert.Len(t, cmd.Aliases, 0)
}

func TestNewSkillsListSubcommand(t *testing.T) {
	cmd := newSkillsListCmd()

	require.NotNil(t, cmd)

	assert.Equal(t, "list", cmd.Use)
	assert.Equal(t, "List installed skills", cmd.Short)

	assert.NotNil(t, cmd.RunE)

	assert.True(t, cmd.HasExample())
	assert.False(t, cmd.HasSubCommands())

	assert.False(t, cmd.HasFlags())

	assert.Len(t, cmd.Aliases, 0)
}

func TestNewListbuiltinSubcommand(t *testing.T) {
	cmd := newSkillsListBuiltinCmd()

	require.NotNil(t, cmd)

	assert.Equal(t, "list-builtin", cmd.Use)
	assert.Equal(t, "List available builtin skills", cmd.Short)

	assert.NotNil(t, cmd.RunE)

	assert.True(t, cmd.HasExample())
	assert.False(t, cmd.HasSubCommands())

	assert.False(t, cmd.HasFlags())

	assert.Len(t, cmd.Aliases, 0)
}

func TestNewRemoveSubcommand(t *testing.T) {
	cmd := newSkillsRemoveCmd()

	require.NotNil(t, cmd)

	assert.Equal(t, "remove", cmd.Use)
	assert.Equal(t, "Remove installed skill", cmd.Short)

	assert.NotNil(t, cmd.RunE)

	assert.True(t, cmd.HasExample())
	assert.False(t, cmd.HasSubCommands())

	assert.False(t, cmd.HasFlags())

	assert.Len(t, cmd.Aliases, 2)
	assert.True(t, cmd.HasAlias("rm"))
	assert.True(t, cmd.HasAlias("uninstall"))
}

func TestNewSearchSubcommand(t *testing.T) {
	cmd := newSkillsSearchCmd()

	require.NotNil(t, cmd)

	assert.Equal(t, "search", cmd.Use)
	assert.Equal(t, "Search available skills", cmd.Short)

	assert.NotNil(t, cmd.RunE)

	assert.False(t, cmd.HasSubCommands())
	assert.False(t, cmd.HasFlags())

	assert.Len(t, cmd.Aliases, 0)
}

func TestNewShowSubcommand(t *testing.T) {
	cmd := newSkillsShowCmd()

	require.NotNil(t, cmd)

	assert.Equal(t, "show", cmd.Use)
	assert.Equal(t, "Show skill details", cmd.Short)

	assert.NotNil(t, cmd.RunE)

	assert.True(t, cmd.HasExample())
	assert.False(t, cmd.HasSubCommands())

	assert.False(t, cmd.HasFlags())

	assert.Len(t, cmd.Aliases, 0)
}
