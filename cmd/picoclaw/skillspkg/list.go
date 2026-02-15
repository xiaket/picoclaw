// PicoClaw - Ultra-lightweight personal AI agent
// License: MIT

package skillspkg

import (
	"fmt"

	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "List installed skills",
	Long:  `Display all installed skills with their descriptions.`,
	Run: func(cmd *cobra.Command, args []string) {
		listImpl()
	},
}

func listImpl() {
	loader := getLoader()
	allSkills := loader.ListSkills()

	if len(allSkills) == 0 {
		fmt.Println("No skills installed.")
		return
	}

	fmt.Println("\nInstalled Skills:")
	fmt.Println("------------------")
	for _, skill := range allSkills {
		fmt.Printf("  ✓ %s (%s)\n", skill.Name, skill.Source)
		if skill.Description != "" {
			fmt.Printf("    %s\n", skill.Description)
		}
	}
}
