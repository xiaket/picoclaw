// PicoClaw - Ultra-lightweight personal AI agent
// License: MIT

package cronpkg

import (
	"fmt"

	"github.com/sipeed/picoclaw/pkg/cron"
	"github.com/spf13/cobra"
)

var EnableCmd = &cobra.Command{
	Use:   "enable <job_id>",
	Short: "Enable a job",
	Long:  `Enable a scheduled job by its ID.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		enableImpl(args[0], false)
	},
}

var DisableCmd = &cobra.Command{
	Use:   "disable <job_id>",
	Short: "Disable a job",
	Long:  `Disable a scheduled job by its ID.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		enableImpl(args[0], true)
	},
}

func enableImpl(jobID string, disable bool) {
	cs := cron.NewCronService(cronStorePath, nil)
	enabled := !disable

	job := cs.EnableJob(jobID, enabled)
	if job != nil {
		status := "enabled"
		if disable {
			status = "disabled"
		}
		fmt.Printf("✓ Job '%s' %s\n", job.Name, status)
	} else {
		fmt.Printf("✗ Job %s not found\n", jobID)
	}
}
