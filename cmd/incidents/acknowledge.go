package incidents

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/uptime-cli/uptimectl/pkg/betteruptime"
)

var (
	acknowledgeAll bool
	acknowledgedBy string
)

// acknowledgeCmd represents the get command
var acknowledgeCmd = &cobra.Command{
	Use:     "acknowledge",
	Short:   "acknowledge an incident",
	Long:    "This will acknowledge an ongoing incident, preventing further escalations",
	Aliases: []string{"ack"},
	Args:    cobra.MinimumNArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		client := betteruptime.NewClient()

		if acknowledgeAll {
			activeIncidents, err := client.ListIncidents(false, 14, 0)
			if err != nil {
				return err
			}
			for _, incident := range activeIncidents {
				fmt.Printf("acknowledging incident %s (%s)\n", incident.Attributes.Name, incident.Id)
				err := client.AcknowledgeIncident(cmd.Context(), incident.Id, acknowledgedBy)
				if err != nil {
					return err
				}
			}
			return nil
		}

		for _, possibleID := range args {
			incidentID, err := betteruptime.IncidentIDFromURL(possibleID)
			if err != nil {
				return err
			}

			err = client.AcknowledgeIncident(cmd.Context(), incidentID, acknowledgedBy)
			if err != nil {
				return err
			}
		}
		return nil
	},
}

func init() {
	IncidentsCmd.AddCommand(acknowledgeCmd)

	acknowledgeCmd.Flags().BoolVarP(&acknowledgeAll, "all", "a", false, "acknowledge all current active incidents")
	acknowledgeCmd.Flags().StringVar(&acknowledgedBy, "acknowledged-by", "uptimectl", "User e-mail or a custom identifier of the entity that acknowledged the incident")
}
