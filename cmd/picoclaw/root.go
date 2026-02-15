// PicoClaw - Ultra-lightweight personal AI agent
// License: MIT

package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var showVersion bool

var rootCmd = &cobra.Command{
	Use:   "picoclaw",
	Short: "Personal AI Assistant",
	Long: `PicoClaw - Ultra-lightweight personal AI agent.
A simple and powerful AI assistant that runs on your local machine.`,
	Run: func(cmd *cobra.Command, args []string) {
		if showVersion {
			printVersion()
			return
		}
		printHelp()
		os.Exit(1)
	},
}

func printHelp() {
	fmt.Printf("%s picoclaw - Personal AI Assistant v%s\n\n", logo, version)
	fmt.Println("Usage: picoclaw <command>")
	fmt.Println()
	fmt.Println("Commands:")
	fmt.Println("  onboard     Initialize picoclaw configuration and workspace")
	fmt.Println("  agent       Interact with the agent directly")
	fmt.Println("  auth        Manage authentication (login, logout, status)")
	fmt.Println("  gateway     Start picoclaw gateway")
	fmt.Println("  status      Show picoclaw status")
	fmt.Println("  cron        Manage scheduled tasks")
	fmt.Println("  migrate     Migrate from OpenClaw to PicoClaw")
	fmt.Println("  skills      Manage skills (install, list, remove)")
	fmt.Println("  version     Show version information")
}

func init() {
	rootCmd.Flags().BoolVarP(&showVersion, "version", "v", false, "Show version information")

	// Add subcommands
	rootCmd.AddCommand(onboardCmd)
	rootCmd.AddCommand(agentCmd)
	rootCmd.AddCommand(gatewayCmd)
	rootCmd.AddCommand(statusCmd)
	rootCmd.AddCommand(migrateCmd)
	rootCmd.AddCommand(authCmd)
	rootCmd.AddCommand(cronCmd)
	rootCmd.AddCommand(skillsCmd)
	rootCmd.AddCommand(versionCmd)
}
