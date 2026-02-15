// PicoClaw - Ultra-lightweight personal AI agent
// License: MIT

package main

import (
	"fmt"

	authpkg "github.com/sipeed/picoclaw/cmd/picoclaw/auth"
	"github.com/spf13/cobra"
)

var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Manage authentication",
	Long:  `Manage authentication for different providers (login, logout, status).`,
	Run: func(cmd *cobra.Command, args []string) {
		authHelp()
	},
}

func init() {
	authCmd.PersistentPreRun = func(cmd *cobra.Command, args []string) {
		authpkg.SetConfigPath(getConfigPath())
	}

	authCmd.AddCommand(authpkg.LoginCmd)
	authCmd.AddCommand(authpkg.LogoutCmd)
	authCmd.AddCommand(authpkg.StatusCmd)
}

func authHelp() {
	fmt.Println("\nAuth commands:")
	fmt.Println("  login       Login via OAuth or paste token")
	fmt.Println("  logout      Remove stored credentials")
	fmt.Println("  status      Show current auth status")
	fmt.Println()
	fmt.Println("Login options:")
	fmt.Println("  --provider <name>    Provider to login with (openai, anthropic)")
	fmt.Println("  --device-code        Use device code flow (for headless environments)")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  picoclaw auth login --provider openai")
	fmt.Println("  picoclaw auth login --provider openai --device-code")
	fmt.Println("  picoclaw auth login --provider anthropic")
	fmt.Println("  picoclaw auth logout --provider openai")
	fmt.Println("  picoclaw auth status")
}
