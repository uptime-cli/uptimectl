package incidents

import (
	"github.com/spf13/cobra"

	"github.com/uptime-cli/uptimectl/pkg/betteruptime"
)

// deleteCmd represents the get command
var deleteCmd = &cobra.Command{
	Use:     "delete",
	Short:   "Delete an incident",
	Aliases: []string{"del", "rm"},
	Args:    cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		client := betteruptime.NewClient()

		for _, incidentID := range args {
			err := client.DeleteIncident(incidentID)
			if err != nil {
				return err
			}
		}
		return nil
	},
}

func init() {
	IncidentsCmd.AddCommand(deleteCmd)
}
