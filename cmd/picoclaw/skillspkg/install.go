// PicoClaw - Ultra-lightweight personal AI agent
// License: MIT

package skillspkg

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/cobra"
)

var InstallCmd = &cobra.Command{
	Use:   "install <repo>",
	Short: "Install skill from GitHub",
	Long:  `Install a skill from a GitHub repository.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		installImpl(args[0])
	},
}

func installImpl(repo string) {
	fmt.Printf("Installing skill from %s...\n", repo)

	installer := getInstaller()
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	if err := installer.InstallFromGitHub(ctx, repo); err != nil {
		fmt.Printf("✗ Failed to install skill: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("✓ Skill '%s' installed successfully!\n", filepath.Base(repo))
}
