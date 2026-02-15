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

var LoginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login via OAuth or paste token",
	Long:  `Login to a provider using OAuth browser flow or paste a token.`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if loginProvider == "" {
			return fmt.Errorf("--provider is required")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		loginImpl()
	},
}

var (
	loginProvider   string
	loginDeviceCode bool
)

func init() {
	LoginCmd.Flags().StringVarP(&loginProvider, "provider", "p", "", "Provider to login with (openai, anthropic)")
	LoginCmd.Flags().BoolVar(&loginDeviceCode, "device-code", false, "Use device code flow (for headless environments)")
}

func loginImpl() {
	switch loginProvider {
	case "openai":
		loginOpenAI(loginDeviceCode)
	case "anthropic":
		loginPasteToken(loginProvider)
	default:
		fmt.Printf("Unsupported provider: %s\n", loginProvider)
		fmt.Println("Supported providers: openai, anthropic")
		os.Exit(1)
	}
}

func loginOpenAI(useDeviceCode bool) {
	cfg := auth.OpenAIOAuthConfig()

	var cred *auth.AuthCredential
	var err error

	if useDeviceCode {
		cred, err = auth.LoginDeviceCode(cfg)
	} else {
		cred, err = auth.LoginBrowser(cfg)
	}

	if err != nil {
		fmt.Printf("Login failed: %v\n", err)
		os.Exit(1)
	}

	if err := auth.SetCredential("openai", cred); err != nil {
		fmt.Printf("Failed to save credentials: %v\n", err)
		os.Exit(1)
	}

	appCfg, err := config.LoadConfig(configPath)
	if err == nil {
		appCfg.Providers.OpenAI.AuthMethod = "oauth"
		if err := config.SaveConfig(configPath, appCfg); err != nil {
			fmt.Printf("Warning: could not update config: %v\n", err)
		}
	}

	fmt.Println("Login successful!")
	if cred.AccountID != "" {
		fmt.Printf("Account: %s\n", cred.AccountID)
	}
}

func loginPasteToken(provider string) {
	cred, err := auth.LoginPasteToken(provider, os.Stdin)
	if err != nil {
		fmt.Printf("Login failed: %v\n", err)
		os.Exit(1)
	}

	if err := auth.SetCredential(provider, cred); err != nil {
		fmt.Printf("Failed to save credentials: %v\n", err)
		os.Exit(1)
	}

	appCfg, err := config.LoadConfig(configPath)
	if err == nil {
		switch provider {
		case "anthropic":
			appCfg.Providers.Anthropic.AuthMethod = "token"
		case "openai":
			appCfg.Providers.OpenAI.AuthMethod = "token"
		}
		if err := config.SaveConfig(configPath, appCfg); err != nil {
			fmt.Printf("Warning: could not update config: %v\n", err)
		}
	}

	fmt.Printf("Token saved for %s!\n", provider)
}
