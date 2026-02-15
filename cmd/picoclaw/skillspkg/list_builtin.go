// PicoClaw - Ultra-lightweight personal AI agent
// License: MIT

package skillspkg

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

var ListBuiltinCmd = &cobra.Command{
	Use:   "list-builtin",
	Short: "List available builtin skills",
	Long:  `Display all builtin skills that can be installed to the workspace.`,
	Run: func(cmd *cobra.Command, args []string) {
		listBuiltinImpl()
	},
}

func listBuiltinImpl() {
	builtinSkillsDir := getBuiltinSkillsDir()

	fmt.Println("\nAvailable Builtin Skills:")
	fmt.Println("-----------------------")

	entries, err := os.ReadDir(builtinSkillsDir)
	if err != nil {
		fmt.Printf("Error reading builtin skills: %v\n", err)
		return
	}

	if len(entries) == 0 {
		fmt.Println("No builtin skills available.")
		return
	}

	for _, entry := range entries {
		if entry.IsDir() {
			skillName := entry.Name()
			skillFile := filepath.Join(builtinSkillsDir, skillName, "SKILL.md")

			description := "No description"
			if data, err := os.ReadFile(skillFile); err == nil {
				for _, line := range strings.Split(string(data), "\n") {
					if strings.HasPrefix(line, "description:") {
						description = strings.TrimSpace(strings.TrimPrefix(line, "description:"))
						break
					}
				}
			}
			status := "✓"
			fmt.Printf("  %s  %s\n", status, entry.Name())
			if description != "" {
				fmt.Printf("     %s\n", description)
			}
		}
	}
}
