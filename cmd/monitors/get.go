package monitors

import (
	"strings"

	"github.com/spf13/cobra"

	"github.com/uptime-cli/uptimectl/pkg/betteruptime"
	"github.com/uptime-cli/uptimectl/pkg/table"
	"github.com/uptime-cli/uptimectl/pkg/timeformat"
)

const NoHeaderKey = "no-header"

var noHeader bool
var monitorGroups []string

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:     "get",
	Short:   "Get a list of monitors",
	Aliases: []string{"g", "list", "ls"},
	Args:    cobra.NoArgs, // TODO: allow to filter get cmd
	RunE: func(cmd *cobra.Command, args []string) error {
		client := betteruptime.NewClient()

		monitors := []betteruptime.Monitor{}
		if len(monitorGroups) > 0 {
			for _, monitorGroup := range monitorGroups {
				group := strings.Split(monitorGroup, ",") // support comma seperated list
				for _, g := range group {
					retrieved, err := client.ListMonitoringGroupMonitors(cmd.Context(), g)
					if err != nil {
						if err == betteruptime.ErrNotFound {
							continue
						}
						return err
					}
					monitors = append(monitors, retrieved...)
				}
			}
		} else {
			m, err := client.ListMonitors(cmd.Context())
			if err != nil {
				return err
			}
			monitors = m
		}

		body := make([][]string, 0, len(monitors))
		for _, item := range monitors {
			body = append(body, []string{
				item.Id,
				item.Attributes.PronounceableName,
				item.Attributes.Status,
				timeformat.FormatTime(item.Attributes.LastCheckedAt.Local(), false),
			})
		}

		if noHeader {
			table.Print(nil, body)
		} else {
			table.Print([]string{"ID", "Name", "Status", "Checked at"}, body)
		}
		return nil
	},
}

func init() {
	MonitorsCmd.AddCommand(getCmd)

	getCmd.Flags().BoolVar(&noHeader, NoHeaderKey, false, "Do not print the header")
	getCmd.Flags().StringArrayVarP(&monitorGroups, "groups", "g", []string{}, "Filter monitors by group")
}
