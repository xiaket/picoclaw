// PicoClaw - Ultra-lightweight personal AI agent
// License: MIT

package main

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/sipeed/picoclaw/pkg/cron"
	"github.com/spf13/cobra"
)

func newCronCmd() *cobra.Command {
	var storePath string

	cmd := &cobra.Command{
		Use:   "cron",
		Short: "Manage scheduled tasks",
		PersistentPreRunE: func(_ *cobra.Command, _ []string) error {
			s, err := getCronStorePath()
			if err != nil {
				return err
			}
			storePath = s
			return nil
		},
	}
	cmd.AddCommand(
		newCronListCmd(storePath),
		newCronAddCmd(storePath),
		newCronRemoveCmd(storePath),
		newCronEnableCmd(storePath),
		newCronDisableCmd(storePath),
	)
	return cmd
}

func newCronListCmd(storePath string) *cobra.Command {
	return &cobra.Command{
		Use:   "list",
		Short: "List all scheduled jobs",
		RunE: func(_ *cobra.Command, _ []string) error {
			return runCronList(storePath)
		},
	}
}

func newCronAddCmd(storePath string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add",
		Short: "Add a new scheduled job",
		RunE: func(cmd *cobra.Command, _ []string) error {
			return runCronAdd(cmd, storePath)
		},
	}
	cmd.Flags().StringP("name", "n", "", "Job name")
	cmd.MarkFlagRequired("name")
	cmd.Flags().StringP("message", "m", "", "Message for agent")
	cmd.MarkFlagRequired("message")
	cmd.Flags().Int64P("every", "e", 0, "Run every N seconds")
	cmd.Flags().StringP("cron", "c", "", "Cron expression (e.g. '0 9 * * *')")
	cmd.Flags().BoolP("deliver", "d", false, "Deliver response to channel")
	cmd.Flags().String("to", "", "Recipient for delivery")
	cmd.Flags().String("channel", "", "Channel for delivery")
	return cmd
}

func newCronRemoveCmd(storePath string) *cobra.Command {
	return &cobra.Command{
		Use:     "remove <job_id>",
		Short:   "Remove a job by ID",
		Example: `picoclaw cron remove 1`,
		Args:    cobra.ExactArgs(1),
		RunE: func(_ *cobra.Command, args []string) error {
			return runCronRemove(args[0], storePath)
		},
	}
}

func newCronEnableCmd(storePath string) *cobra.Command {
	return &cobra.Command{
		Use:     "enable",
		Short:   "Enable a job",
		Example: `picoclaw cron enable 1`,
		Args:    cobra.ExactArgs(1),
		RunE: func(_ *cobra.Command, args []string) error {
			return runCronEnable(args[0], storePath)
		},
	}
}

func newCronDisableCmd(storePath string) *cobra.Command {
	return &cobra.Command{
		Use:     "disable",
		Short:   "Disable a job",
		Example: `picoclaw cron disable 1`,
		Args:    cobra.ExactArgs(1),
		RunE: func(_ *cobra.Command, args []string) error {
			return runCronDisable(args[0], storePath)
		},
	}
}

func getCronStorePath() (string, error) {
	cfg, err := loadConfig()
	if err != nil {
		return "", fmt.Errorf("Error loading config: %w", err)
	}
	return filepath.Join(cfg.WorkspacePath(), "cron", "jobs.json"), nil
}

func runCronList(storePath string) error {
	cs := cron.NewCronService(storePath, nil)
	jobs := cs.ListJobs(true)

	if len(jobs) == 0 {
		fmt.Println("No scheduled jobs.")
		return nil
	}

	fmt.Println("\nScheduled Jobs:")
	fmt.Println("----------------")
	for _, job := range jobs {
		var schedule string
		if job.Schedule.Kind == "every" && job.Schedule.EveryMS != nil {
			schedule = fmt.Sprintf("every %ds", *job.Schedule.EveryMS/1000)
		} else if job.Schedule.Kind == "cron" {
			schedule = job.Schedule.Expr
		} else {
			schedule = "one-time"
		}

		nextRun := "scheduled"
		if job.State.NextRunAtMS != nil {
			nextTime := time.UnixMilli(*job.State.NextRunAtMS)
			nextRun = nextTime.Format("2006-01-02 15:04")
		}

		status := "enabled"
		if !job.Enabled {
			status = "disabled"
		}

		fmt.Printf("  %s (%s)\n", job.Name, job.ID)
		fmt.Printf("    Schedule: %s\n", schedule)
		fmt.Printf("    Status: %s\n", status)
		fmt.Printf("    Next run: %s\n", nextRun)
	}

	return nil
}

func runCronAdd(cmd *cobra.Command, storePath string) error {
	name, _ := cmd.Flags().GetString("name")
	message, _ := cmd.Flags().GetString("message")
	everySec, _ := cmd.Flags().GetInt64("every")
	cronExpr, _ := cmd.Flags().GetString("cron")
	deliver, _ := cmd.Flags().GetBool("deliver")
	to, _ := cmd.Flags().GetString("to")
	channel, _ := cmd.Flags().GetString("channel")

	if everySec == 0 && cronExpr == "" {
		fmt.Println("Error: Either --every or --cron must be specified")
		return nil
	}

	var schedule cron.CronSchedule
	if everySec != 0 {
		everyMS := everySec * 1000
		schedule = cron.CronSchedule{
			Kind:    "every",
			EveryMS: &everyMS,
		}
	} else {
		schedule = cron.CronSchedule{
			Kind: "cron",
			Expr: cronExpr,
		}
	}

	cs := cron.NewCronService(storePath, nil)
	job, err := cs.AddJob(name, schedule, message, deliver, channel, to)
	if err != nil {
		fmt.Printf("Error adding job: %v\n", err)
		return nil
	}

	fmt.Printf("\u2713 Added job '%s' (%s)\n", job.Name, job.ID)
	return nil
}

func runCronRemove(jobID, storePath string) error {
	cs := cron.NewCronService(storePath, nil)
	if cs.RemoveJob(jobID) {
		fmt.Printf("\u2713 Removed job %s\n", jobID)
	} else {
		fmt.Printf("\u2717 Job %s not found\n", jobID)
	}
	return nil
}

func runCronEnable(jobID, storePath string) error {
	cs := cron.NewCronService(storePath, nil)
	job := cs.EnableJob(jobID, true)
	if job != nil {
		fmt.Printf("\u2713 Job '%s' enabled\n", job.Name)
	} else {
		fmt.Printf("\u2717 Job %s not found\n", jobID)
	}
	return nil
}

func runCronDisable(jobID, storePath string) error {
	cs := cron.NewCronService(storePath, nil)
	job := cs.EnableJob(jobID, false)
	if job != nil {
		fmt.Printf("\u2713 Job '%s' disabled\n", job.Name)
	} else {
		fmt.Printf("\u2717 Job %s not found\n", jobID)
	}
	return nil
}
