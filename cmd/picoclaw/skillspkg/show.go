// PicoClaw - Ultra-lightweight personal AI agent
// License: MIT

package skillspkg

import (
	"fmt"

	"github.com/spf13/cobra"
)

var ShowCmd = &cobra.Command{
	Use:   "show <name>",
	Short: "Show skill details",
	Long:  `Display the full content and details of an installed skill.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		showImpl(args[0])
	},
}

func showImpl(skillName string) {
	loader := getLoader()
	content, ok := loader.LoadSkill(skillName)
	if !ok {
		fmt.Printf("✗ Skill '%s' not found\n", skillName)
		return
	}

	fmt.Printf("\n📦 Skill: %s\n", skillName)
	fmt.Println("----------------------")
	fmt.Println(content)
}
