// PicoClaw - Ultra-lightweight personal AI agent
// License: MIT

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/sipeed/picoclaw/pkg/config"
)

var (
	version   = "dev"
	gitCommit string
	buildTime string
	goVersion string
)

const logo = "🦞"

// formatVersion returns the version string with optional git commit
func formatVersion() string {
	v := version
	if gitCommit != "" {
		v += fmt.Sprintf(" (git: %s)", gitCommit)
	}
	return v
}

// formatBuildInfo returns build time and go version info
func formatBuildInfo() (build string, goVer string) {
	if buildTime != "" {
		build = buildTime
	}
	goVer = goVersion
	if goVer == "" {
		goVer = runtime.Version()
	}
	return
}

// getConfigPath returns the path to the config file
func getConfigPath() string {
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".picoclaw", "config.json")
}

// loadConfig loads the configuration from the default path
func loadConfig() (*config.Config, error) {
	return config.LoadConfig(getConfigPath())
}

// exitOnError prints an error message and exits with status 1
func exitOnError(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
	os.Exit(1)
}
