package incidents

import (
	"github.com/spf13/cobra"

	"github.com/uptime-cli/uptimectl/pkg/betteruptime"
)

// deleteCmd represents the get command
var resolveCmd = &cobra.Command{
	Use:     "resolve",
	Short:   "Resolve an incident",
	Aliases: []string{"res"},
	Args:    cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		client := betteruptime.NewClient()

		for _, incidentID := range args {
			err := client.ResolveIncident(incidentID, acknowledgedBy)
			if err != nil {
				return err
			}
		}
		return nil
	},
}

func init() {
	IncidentsCmd.AddCommand(resolveCmd)
	resolveCmd.Flags().StringVar(&acknowledgedBy, "acknowledged-by", "uptimectl", "User e-mail or a custom identifier of the entity that acknowledged the incident")
}
