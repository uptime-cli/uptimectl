package monitors

import (
	_ "embed"

	"github.com/spf13/cobra"
	"github.com/uptime-cli/uptimectl/pkg/betteruptime"
)

//go:embed examples/unpause.txt
var unPauseExamples string

// unPauseCmd represents the unpause command
var unPauseCmd = &cobra.Command{
	Use:     "unpause",
	Short:   "Unpause monitors by ID",
	Args:    cobra.MinimumNArgs(1),
	Example: unPauseExamples,
	RunE: func(cmd *cobra.Command, args []string) error {
		client := betteruptime.NewClient()

		for _, monitor := range args {
			err := client.UnpauseMonitor(monitor)
			if err != nil {
				return err
			}
		}
		return nil
	},
}

func init() {
	MonitorsCmd.AddCommand(unPauseCmd)
}
