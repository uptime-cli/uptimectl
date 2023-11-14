package monitorgroups

import (
	"github.com/spf13/cobra"

	"github.com/uptime-cli/uptimectl/pkg/betteruptime"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"rm"},
	Short:   "delete a monitor group",
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		// organisation := contextmanager.Organisation()
		client := betteruptime.NewClient()
		err := client.DeleteMonitorGroup(args[0])
		if err != nil {
			return err
		}
		return nil
	},
}

func init() {
	MonitorGroupsCmd.AddCommand(deleteCmd)
}
