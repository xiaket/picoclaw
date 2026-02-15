// PicoClaw - Ultra-lightweight personal AI agent
// License: MIT

package cronpkg

import (
	"fmt"

	"github.com/sipeed/picoclaw/pkg/cron"
	"github.com/spf13/cobra"
)

var RemoveCmd = &cobra.Command{
	Use:   "remove <job_id>",
	Short: "Remove a job by ID",
	Long:  `Remove a scheduled job by its ID.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		removeImpl(args[0])
	},
}

func removeImpl(jobID string) {
	cs := cron.NewCronService(cronStorePath, nil)
	if cs.RemoveJob(jobID) {
		fmt.Printf("✓ Removed job %s\n", jobID)
	} else {
		fmt.Printf("✗ Job %s not found\n", jobID)
	}
}
