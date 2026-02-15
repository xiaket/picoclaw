// PicoClaw - Ultra-lightweight personal AI agent
// License: MIT

package main

import (
	"fmt"
	"os"

	"github.com/sipeed/picoclaw/pkg/migrate"
	"github.com/spf13/cobra"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate from OpenClaw to PicoClaw",
	Long:  `Migrate configuration and workspace from OpenClaw to PicoClaw.`,
	Run: func(cmd *cobra.Command, args []string) {
		migrateImpl()
	},
}

var (
	migrateDryRun        bool
	migrateRefresh       bool
	migrateConfigOnly    bool
	migrateWorkspaceOnly bool
	migrateForce         bool
	migrateOpenClawHome  string
	migratePicoClawHome  string
)

func init() {
	migrateCmd.Flags().BoolVar(&migrateDryRun, "dry-run", false, "Show what would be migrated without making changes")
	migrateCmd.Flags().BoolVar(&migrateRefresh, "refresh", false, "Re-sync workspace files from OpenClaw (repeatable)")
	migrateCmd.Flags().BoolVar(&migrateConfigOnly, "config-only", false, "Only migrate config, skip workspace files")
	migrateCmd.Flags().BoolVar(&migrateWorkspaceOnly, "workspace-only", false, "Only migrate workspace files, skip config")
	migrateCmd.Flags().BoolVar(&migrateForce, "force", false, "Skip confirmation prompts")
	migrateCmd.Flags().StringVar(&migrateOpenClawHome, "openclaw-home", "", "Override OpenClaw home directory (default: ~/.openclaw)")
	migrateCmd.Flags().StringVar(&migratePicoClawHome, "picoclaw-home", "", "Override PicoClaw home directory (default: ~/.picoclaw)")
}

func migrateImpl() {
	opts := migrate.Options{
		DryRun:        migrateDryRun,
		Refresh:       migrateRefresh,
		ConfigOnly:    migrateConfigOnly,
		WorkspaceOnly: migrateWorkspaceOnly,
		Force:         migrateForce,
		OpenClawHome:  migrateOpenClawHome,
		PicoClawHome:  migratePicoClawHome,
	}

	result, err := migrate.Run(opts)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	if !opts.DryRun {
		migrate.PrintSummary(result)
	}
}
