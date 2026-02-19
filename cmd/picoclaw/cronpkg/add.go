// PicoClaw - Ultra-lightweight personal AI agent
// License: MIT

package cronpkg

import (
	"fmt"
	"os"

	"github.com/sipeed/picoclaw/pkg/cron"
	"github.com/spf13/cobra"
)

var AddCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new scheduled job",
	Long:  `Add a new cron job with specified schedule and message.`,
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if cronName == "" {
			return fmt.Errorf("--name is required")
		}
		if cronMessage == "" {
			return fmt.Errorf("--message is required")
		}
		if cronEvery == 0 && cronCronExpr == "" {
			return fmt.Errorf("Either --every or --cron must be specified")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		addImpl()
	},
}

var (
	cronName     string
	cronMessage  string
	cronEvery    int64
	cronCronExpr string
	cronDeliver  bool
	cronTo       string
	cronChannel  string
)

func init() {
	AddCmd.Flags().StringVarP(&cronName, "name", "n", "", "Job name (required)")
	AddCmd.Flags().StringVarP(&cronMessage, "message", "m", "", "Message for agent (required)")
	AddCmd.Flags().StringVarP(&cronCronExpr, "cron", "c", "", "Cron expression (e.g. '0 9 * * *')")
	AddCmd.Flags().StringVar(&cronTo, "to", "", "Recipient for delivery")
	AddCmd.Flags().StringVar(&cronChannel, "channel", "", "Channel for delivery")
	AddCmd.Flags().Int64VarP(&cronEvery, "every", "e", 0, "Run every N seconds")
	AddCmd.Flags().BoolVarP(&cronDeliver, "deliver", "d", false, "Deliver response to channel")
}

func addImpl() {
	var schedule cron.CronSchedule
	if cronEvery > 0 {
		everyMS := cronEvery * 1000
		schedule = cron.CronSchedule{
			Kind:    "every",
			EveryMS: &everyMS,
		}
	} else {
		schedule = cron.CronSchedule{
			Kind: "cron",
			Expr: cronCronExpr,
		}
	}

	cs := cron.NewCronService(cronStorePath, nil)
	job, err := cs.AddJob(cronName, schedule, cronMessage, cronDeliver, cronChannel, cronTo)
	if err != nil {
		fmt.Printf("Error adding job: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("✓ Added job '%s' (%s)\n", job.Name, job.ID)
}
