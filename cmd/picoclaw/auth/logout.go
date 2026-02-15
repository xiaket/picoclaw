// PicoClaw - Ultra-lightweight personal AI agent
// License: MIT

package auth

import (
	"fmt"
	"os"

	"github.com/sipeed/picoclaw/pkg/auth"
	"github.com/sipeed/picoclaw/pkg/config"
	"github.com/spf13/cobra"
)

var LogoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Remove stored credentials",
	Long:  `Remove stored authentication credentials for a specific provider or all providers.`,
	Run: func(cmd *cobra.Command, args []string) {
		logoutImpl()
	},
}

var logoutProvider string

func init() {
	LogoutCmd.Flags().StringVarP(&logoutProvider, "provider", "p", "", "Provider to logout from (openai, anthropic)")
}

func logoutImpl() {
	if logoutProvider != "" {
		if err := auth.DeleteCredential(logoutProvider); err != nil {
			fmt.Printf("Failed to remove credentials: %v\n", err)
			os.Exit(1)
		}

		appCfg, err := config.LoadConfig(configPath)
		if err == nil {
			switch logoutProvider {
			case "openai":
				appCfg.Providers.OpenAI.AuthMethod = ""
			case "anthropic":
				appCfg.Providers.Anthropic.AuthMethod = ""
			}
			config.SaveConfig(configPath, appCfg)
		}

		fmt.Printf("Logged out from %s\n", logoutProvider)
	} else {
		if err := auth.DeleteAllCredentials(); err != nil {
			fmt.Printf("Failed to remove credentials: %v\n", err)
			os.Exit(1)
		}

		appCfg, err := config.LoadConfig(configPath)
		if err == nil {
			appCfg.Providers.OpenAI.AuthMethod = ""
			appCfg.Providers.Anthropic.AuthMethod = ""
			config.SaveConfig(configPath, appCfg)
		}

		fmt.Println("Logged out from all providers")
	}
}
