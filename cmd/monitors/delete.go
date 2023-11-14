package monitors

import (
	"github.com/spf13/cobra"

	"github.com/uptime-cli/uptimectl/pkg/betteruptime"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"rm"},
	Short:   "delete a monitor",
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		client := betteruptime.NewClient()
		err := client.DeleteMonitor(args[0])
		if err != nil {
			return err
		}
		return nil
	},
}

func init() {
	MonitorsCmd.AddCommand(deleteCmd)
}
