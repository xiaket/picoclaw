// PicoClaw - Ultra-lightweight personal AI agent
// License: MIT

package main

import (
	"fmt"

	"github.com/sipeed/picoclaw/cmd/picoclaw/cronpkg"
	"github.com/spf13/cobra"
)

var cronCmd = &cobra.Command{
	Use:   "cron",
	Short: "Manage scheduled tasks",
	Long:  `Manage cron jobs and scheduled tasks.`,
}

func init() {
	// PreRun to load config and store cronStorePath for subcommands
	cronCmd.PersistentPreRun = func(cmd *cobra.Command, args []string) {
		cfg, err := loadConfig()
		if err != nil {
			fmt.Printf("Error loading config: %v\n", err)
			return
		}
		cronpkg.SetCronStorePath(cfg.WorkspacePath())
	}

	cronCmd.AddCommand(cronpkg.ListCmd)
	cronCmd.AddCommand(cronpkg.AddCmd)
	cronCmd.AddCommand(cronpkg.RemoveCmd)
	cronCmd.AddCommand(cronpkg.EnableCmd)
	cronCmd.AddCommand(cronpkg.DisableCmd)
}
