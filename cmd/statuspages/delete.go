package statuspages

import (
	"github.com/spf13/cobra"

	"github.com/uptime-cli/uptimectl/pkg/betteruptime"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:     "delete [id]",
	Aliases: []string{"rm"},
	Short:   "delete a status page",
	Args:    cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		// organisation := contextmanager.Organisation()
		client := betteruptime.NewClient()
		err := client.DeleteStatusPage(args[0])
		if err != nil {
			return err
		}
		return nil
	},
}

func init() {
	StatusPagesCmd.AddCommand(deleteCmd)
}
