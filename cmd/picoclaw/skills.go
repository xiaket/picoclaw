// PicoClaw - Ultra-lightweight personal AI agent
// License: MIT

package main

import (
	"fmt"
	"path/filepath"

	"github.com/sipeed/picoclaw/cmd/picoclaw/skillspkg"
	"github.com/spf13/cobra"
)

var skillsCmd = &cobra.Command{
	Use:   "skills",
	Short: "Manage skills",
	Long:  `Manage skills installation and listing.`,
}

func init() {
	// PreRun to load config and create installer/loader for subcommands
	skillsCmd.PersistentPreRun = func(cmd *cobra.Command, args []string) {
		cfg, err := loadConfig()
		if err != nil {
			fmt.Printf("Error loading config: %v\n", err)
			return
		}
		workspace := cfg.WorkspacePath()

		// 获取全局配置目录和内置 skills 目录
		globalDir := filepath.Dir(getConfigPath())
		globalSkillsDir := filepath.Join(globalDir, "skills")
		builtinSkillsDir := filepath.Join(globalDir, "picoclaw", "skills")

		skillspkg.SetWorkspace(workspace)
		skillspkg.SetGlobalDirs(globalSkillsDir, builtinSkillsDir)
	}

	skillsCmd.AddCommand(skillspkg.ListCmd)
	skillsCmd.AddCommand(skillspkg.InstallCmd)
	skillsCmd.AddCommand(skillspkg.RemoveCmd)
	skillsCmd.AddCommand(skillspkg.InstallBuiltinCmd)
	skillsCmd.AddCommand(skillspkg.ListBuiltinCmd)
	skillsCmd.AddCommand(skillspkg.SearchCmd)
	skillsCmd.AddCommand(skillspkg.ShowCmd)
}
