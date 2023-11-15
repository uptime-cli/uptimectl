package incidents

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/uptime-cli/uptimectl/pkg/betteruptime"
	"github.com/uptime-cli/uptimectl/pkg/table"
	"github.com/uptime-cli/uptimectl/pkg/timeformat"
)

const NoHeaderKey = "no-header"

var noHeader bool

var (
	showAll       bool
	showDays      int
	showExactTime bool
	showMax       int
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:     "list",
	Short:   "Get a list of incidents",
	Aliases: []string{"g", "get", "ls"},
	Args:    cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		client := betteruptime.NewClient()
		incidents, err := client.ListIncidents(showAll, showDays, showMax)
		if err != nil {
			return err
		}
		body := make([][]string, 0, len(incidents))
		for _, incident := range incidents {
			resolved := incident.Attributes.Resolved_at != nil

			resolvedAt := ""
			if resolved {
				resolvedAt = timeformat.FormatTime(incident.Attributes.Resolved_at.Local(), showExactTime)
			}

			body = append(body, []string{
				incident.Id,
				incident.Attributes.Name,
				fmt.Sprint(resolved),
				timeformat.FormatTime(incident.Attributes.Started_at.Local(), showExactTime),

				resolvedAt,
				strings.Join(incident.Attributes.Regions, ","),
				incident.Attributes.Cause,
			})
		}

		if noHeader {
			table.Print(nil, body)
		} else {
			table.Print([]string{"ID", "Name", "Resolved", "Started at", "Resolved at", "regions", "cause"}, body)
		}
		return nil
	},
}

func init() {
	IncidentsCmd.AddCommand(getCmd)

	getCmd.Flags().BoolVar(&noHeader, NoHeaderKey, false, "Do not print the header")
	getCmd.Flags().BoolVarP(&showAll, "all", "a", false, "show all incidents, including resolved")
	getCmd.Flags().IntVarP(&showDays, "show-days", "d", 7, "show incidents from within the the past amount of days")
	getCmd.Flags().IntVarP(&showMax, "limit", "l", 0, "limit the amount of incidents displayed")
	getCmd.Flags().BoolVar(&showExactTime, "show-exact-time", false, "show dates using the full timestamp")
}
