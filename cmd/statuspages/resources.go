package statuspages

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/uptime-cli/uptimectl/pkg/betteruptime"
	"github.com/uptime-cli/uptimectl/pkg/table"
)

// getResourcesCmd represents the get command
var getResourcesCmd = &cobra.Command{
	Use:     "resources [id]",
	Short:   "Get a list of resources belong to a status page",
	Aliases: []string{},
	Args:    cobra.MaximumNArgs(10),
	RunE: func(cmd *cobra.Command, args []string) error {
		client := betteruptime.NewClient()

		statusPageResources := []betteruptime.StatusPageResource{}
		for _, id := range args {
			resources, err := client.GetStatusPageResources(id)
			if err != nil {
				return err
			}
			statusPageResources = append(statusPageResources, resources...)
		}

		body := make([][]string, 0, len(statusPageResources))
		for _, item := range statusPageResources {
			body = append(body, []string{
				item.Id,
				item.Attributes.PublicName,
				fmt.Sprint(item.Attributes.History),
				item.Attributes.Explanation,
				item.Attributes.ResourceType,
				item.Attributes.WidgetType,
			})
		}

		if noHeader {
			table.Print(nil, body)
		} else {
			table.Print([]string{"ID", "Name", "Show History", "Explaination", "Type", "Widget"}, body)
		}
		return nil
	},
}

func init() {
	StatusPagesCmd.AddCommand(getResourcesCmd)

	getResourcesCmd.Flags().BoolVar(&noHeader, NoHeaderKey, false, "Do not print the header")
}
