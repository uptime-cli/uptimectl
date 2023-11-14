package statuspages

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/uptime-cli/uptimectl/pkg/betteruptime"
	"github.com/uptime-cli/uptimectl/pkg/table"
	"github.com/uptime-cli/uptimectl/pkg/timeformat"
)

const NoHeaderKey = "no-header"

var noHeader bool

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:     "get [id]",
	Short:   "Get a list of status pages",
	Aliases: []string{"g", "list", "ls"},
	Args:    cobra.MaximumNArgs(10),
	RunE: func(cmd *cobra.Command, args []string) error {
		// organisation := contextmanager.Organisation()
		// clusters, err := platformapi.Client().GetClustersByOrg(organisation)
		client := betteruptime.NewClient()

		statusPages := []betteruptime.StatusPage{}
		if len(args) > 0 {
			for _, id := range args {
				statuspage, err := client.GetStatusPage(id)
				if err != nil {
					return err
				}
				statusPages = append(statusPages, *statuspage)
			}
		} else {
			var err error
			statusPages, err = client.ListStatusPages()
			if err != nil {
				return err
			}
		}

		body := make([][]string, 0, len(statusPages))
		for _, item := range statusPages {
			body = append(body, []string{
				item.Id,
				item.Attributes.CompanyName,
				item.Attributes.CustomDomain,
				fmt.Sprint(item.Attributes.PasswordEnabled),
				fmt.Sprint(item.Attributes.Subscribable),
				timeformat.FormatTime(item.Attributes.CreatedAt.Local(), false),
			})
		}

		if noHeader {
			table.Print(nil, body)
		} else {
			table.Print([]string{"ID", "Name", "Custom domain", "password protected", "Subscribable", "Created at"}, body)
		}
		return nil
	},
}

func init() {
	StatusPagesCmd.AddCommand(getCmd)

	getCmd.Flags().BoolVar(&noHeader, NoHeaderKey, false, "Do not print the header")
}
