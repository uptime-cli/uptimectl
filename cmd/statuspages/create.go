package statuspages

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/uptime-cli/uptimectl/pkg/betteruptime"
	"github.com/uptime-cli/uptimectl/pkg/table"
)

// getCmd represents the get command
var createCmd = &cobra.Command{
	Use:   "create [company name]",
	Short: "create a status page",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		// organisation := contextmanager.Organisation()
		client := betteruptime.NewClient()
		statusPage, err := client.CreateStatusPage(betteruptime.CreateStatusPageRequest{
			CompanyName: &args[0],
			CompanyURL:  nil,
			SubDomain:   nil,
			Timezone:    nil,
		})
		if err != nil {
			return err
		}
		body := make([][]string, 0, 1)
		body = append(body, []string{
			statusPage.Id,
			statusPage.Attributes.CompanyName,
			fmt.Sprint(statusPage.Attributes.Subscribable),
			statusPage.Attributes.CreatedAt.Local().String(),
		})

		if noHeader {
			table.Print(nil, body)
		} else {
			table.Print([]string{"ID", "Name", "Subscribable", "Created at"}, body)
		}
		return nil
	},
}

func init() {
	StatusPagesCmd.AddCommand(createCmd)
}
