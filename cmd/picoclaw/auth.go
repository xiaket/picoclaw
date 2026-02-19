// PicoClaw - Ultra-lightweight personal AI agent
// License: MIT

package main

import (
	authpkg "github.com/sipeed/picoclaw/cmd/picoclaw/auth"
	"github.com/spf13/cobra"
)

var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Manage authentication",
	Long:  `Manage authentication for different providers (login, logout, status).`,
}

func init() {
	authCmd.PersistentPreRun = func(cmd *cobra.Command, args []string) {
		authpkg.SetConfigPath(getConfigPath())
	}

	authCmd.AddCommand(authpkg.LoginCmd)
	authCmd.AddCommand(authpkg.LogoutCmd)
	authCmd.AddCommand(authpkg.StatusCmd)
}
