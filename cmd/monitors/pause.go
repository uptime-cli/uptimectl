package monitors

import (
	_ "embed"

	"github.com/spf13/cobra"
	"github.com/uptime-cli/uptimectl/pkg/betteruptime"
)

//go:embed examples/pause.txt
var pauseExamples string

// pauseCmd represents the pause command
var pauseCmd = &cobra.Command{
	Use:     "pause",
	Short:   "Pause monitors by ID",
	Args:    cobra.MinimumNArgs(1),
	Example: pauseExamples,
	RunE: func(cmd *cobra.Command, args []string) error {
		client := betteruptime.NewClient()

		for _, monitor := range args {
			err := client.PauseMonitor(monitor)
			if err != nil {
				return err
			}
		}
		return nil
	},
}

func init() {
	MonitorsCmd.AddCommand(pauseCmd)
}
