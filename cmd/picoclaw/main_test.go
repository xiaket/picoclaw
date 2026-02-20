package main

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRootCommand(t *testing.T) {
	require.NotNil(t, rootCmd)

	assert.Equal(t, "picoclaw", rootCmd.Use)
	assert.Equal(t, fmt.Sprintf("%s picoclaw - Personal AI Assistant", logo), rootCmd.Short)

	assert.True(t, rootCmd.HasSubCommands())
	assert.True(t, rootCmd.HasAvailableSubCommands())

	assert.True(t, rootCmd.HasFlags())

	assert.NotNil(t, rootCmd.Flags().Lookup("version"))

	allowedCommands := map[string]struct{}{
		"agent":   {},
		"auth":    {},
		"cron":    {},
		"gateway": {},
		"migrate": {},
		"onboard": {},
		"skills":  {},
		"status":  {},
		"version": {},
	}

	subcommands := rootCmd.Commands()
	assert.Len(t, subcommands, len(allowedCommands))

	for _, subcmd := range subcommands {
		_, found := allowedCommands[subcmd.Name()]
		assert.True(t, found, "unexpected subcommand %q", subcmd.Name())

		assert.False(t, subcmd.Hidden)
	}
}

func TestNewVersionCommand(t *testing.T) {
	cmd := newVersionCmd()

	require.NotNil(t, cmd)

	assert.Equal(t, "version", cmd.Use)

	assert.Len(t, cmd.Aliases, 1)
	assert.True(t, cmd.HasAlias("v"))

	assert.False(t, cmd.HasFlags())

	assert.Equal(t, "Show version information", cmd.Short)

	assert.False(t, cmd.HasSubCommands())

	assert.NotNil(t, cmd.Run)
	assert.Nil(t, cmd.RunE)

	assert.Nil(t, cmd.PersistentPreRun)
	assert.Nil(t, cmd.PersistentPostRun)
}

func TestGetConfigPath(t *testing.T) {
	t.Setenv("HOME", "/tmp/home")

	got := getConfigPath()
	want := filepath.Join("/tmp/home", ".picoclaw", "config.json")

	assert.Equal(t, want, got)
}

func TestFormatBuildInfo_WithBuildTimeAndGoVersion(t *testing.T) {
	oldBuildTime, oldGoVersion := buildTime, goVersion
	t.Cleanup(func() {
		buildTime, goVersion = oldBuildTime, oldGoVersion
	})

	buildTime = "2026-02-20T12:00:00Z"
	goVersion = "go1.22.0"

	build, goVer := formatBuildInfo()

	assert.Equal(t, "2026-02-20T12:00:00Z", build)
	assert.Equal(t, "go1.22.0", goVer)
}

func TestFormatBuildInfo_FallbackToRuntimeVersion(t *testing.T) {
	oldBuildTime, oldGoVersion := buildTime, goVersion
	t.Cleanup(func() {
		buildTime, goVersion = oldBuildTime, oldGoVersion
	})

	buildTime = ""
	goVersion = ""

	build, goVer := formatBuildInfo()

	assert.Equal(t, "", build)
	assert.Equal(t, runtime.Version(), goVer)
}

func TestFormatBuildInfo_UsesBuildTimeAndGoVersion_WhenSet(t *testing.T) {
	oldBuildTime, oldGoVersion := buildTime, goVersion
	t.Cleanup(func() { buildTime, goVersion = oldBuildTime, oldGoVersion })

	buildTime = "2026-02-20T00:00:00Z"
	goVersion = "go1.23.0"

	build, goVer := formatBuildInfo()

	assert.Equal(t, buildTime, build)
	assert.Equal(t, goVersion, goVer)
}

func TestFormatBuildInfo_EmptyBuildTime_ReturnsEmptyBuild(t *testing.T) {
	oldBuildTime, oldGoVersion := buildTime, goVersion
	t.Cleanup(func() { buildTime, goVersion = oldBuildTime, oldGoVersion })

	buildTime = ""
	goVersion = "go1.23.0"

	build, goVer := formatBuildInfo()

	assert.Empty(t, build)
	assert.Equal(t, goVersion, goVer)
}

func TestFormatBuildInfo_EmptyGoVersion_FallsBackToRuntimeVersion(t *testing.T) {
	oldBuildTime, oldGoVersion := buildTime, goVersion
	t.Cleanup(func() { buildTime, goVersion = oldBuildTime, oldGoVersion })

	buildTime = "x"
	goVersion = ""

	build, goVer := formatBuildInfo()

	assert.Equal(t, "x", build)
	assert.Equal(t, runtime.Version(), goVer)
}

func TestGetConfigPath_Windows(t *testing.T) {
	if runtime.GOOS != "windows" {
		t.Skip("windows-specific HOME behavior varies; run on windows")
	}

	t.Setenv("USERPROFILE", `C:\Users\Test`)

	got := getConfigPath()
	want := filepath.Join(`C:\Users\Test`, ".picoclaw", "config.json")

	require.True(t, strings.EqualFold(got, want), "GetConfigPath() = %q, want %q", got, want)
}
