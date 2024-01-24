package incidents

import (
	"github.com/spf13/cobra"

	"github.com/uptime-cli/uptimectl/pkg/betteruptime"
)

var (
	resolvedBy string
)

// resolveCmd represents the get command
var resolveCmd = &cobra.Command{
	Use:     "resolve",
	Short:   "Resolve an incident",
	Aliases: []string{"res"},
	Args:    cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		client := betteruptime.NewClient()

		for _, possibleID := range args {
			incidentID, err := betteruptime.IncidentIDFromURL(possibleID)
			if err != nil {
				return err
			}

			err = client.ResolveIncident(incidentID, resolvedBy)
			if err != nil && err != betteruptime.ErrIncidentAlreadyResolved {
				return err
			}
		}
		return nil
	},
}

func init() {
	IncidentsCmd.AddCommand(resolveCmd)
	resolveCmd.Flags().StringVar(&resolvedBy, "resolved-by", "uptimectl", "User e-mail or a custom identifier of the entity that resolved the incident")
}
