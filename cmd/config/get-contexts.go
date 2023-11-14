package config

import (
	"github.com/spf13/cobra"

	"github.com/uptime-cli/uptimectl/pkg/contextmanager"
	"github.com/uptime-cli/uptimectl/pkg/table"
)

const NoHeaderKey = "no-header"

var noHeader bool

// getContextCmd represents the get command
var getContextCmd = &cobra.Command{
	Use:   "get-contexts",
	Short: "Get the contexts",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		contexts := contextmanager.GlobalContextManager().Config().Contexts

		contextNames := []string{}
		for _, context := range contexts {
			contextNames = append(contextNames, context.Name)
		}
		body := make([][]string, 0, len(contexts))
		for _, c := range contexts {
			body = append(body, []string{c.Name, c.Context.Organisation, c.Context.User, c.Context.API})
		}

		if noHeader {
			table.Print(nil, body)
		} else {
			table.Print([]string{"Name", "Organisation", "User", "API"}, body)
		}
		return nil
	},
}

func init() {
	ConfigCmd.AddCommand(getContextCmd)

	getContextCmd.Flags().BoolVar(&noHeader, NoHeaderKey, false, "Do not print the header")
}
