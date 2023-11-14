package monitorgroups

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/uptime-cli/uptimectl/pkg/betteruptime"
	"github.com/uptime-cli/uptimectl/pkg/table"
)

// getCmd represents the get command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create a monitor group",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		// organisation := contextmanager.Organisation()
		client := betteruptime.NewClient()
		monitorGroup, err := client.CreateMonitorGroup(args[0])
		if err != nil {
			return err
		}
		body := make([][]string, 0, 1)
		body = append(body, []string{
			monitorGroup.Id,
			monitorGroup.Attributes.Name,
			fmt.Sprint(monitorGroup.Attributes.Paused),
			monitorGroup.Attributes.CreatedAt.Local().String(),
		})

		if noHeader {
			table.Print(nil, body)
		} else {
			table.Print([]string{"ID", "Name", "Paused", "Created at"}, body)
		}
		return nil
	},
}

func init() {
	MonitorGroupsCmd.AddCommand(createCmd)
}
