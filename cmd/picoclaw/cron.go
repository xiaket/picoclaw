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
	Run: func(cmd *cobra.Command, args []string) {
		cronHelp()
	},
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

func cronHelp() {
	fmt.Println("\nCron commands:")
	fmt.Println("  list              List all scheduled jobs")
	fmt.Println("  add              Add a new scheduled job")
	fmt.Println("  remove <id>       Remove a job by ID")
	fmt.Println("  enable <id>      Enable a job")
	fmt.Println("  disable <id>     Disable a job")
	fmt.Println()
	fmt.Println("Add options:")
	fmt.Println("  -n, --name       Job name")
	fmt.Println("  -m, --message    Message for agent")
	fmt.Println("  -e, --every      Run every N seconds")
	fmt.Println("  -c, --cron       Cron expression (e.g. '0 9 * * *')")
	fmt.Println("  -d, --deliver     Deliver response to channel")
	fmt.Println("  --to             Recipient for delivery")
	fmt.Println("  --channel        Channel for delivery")
}
