package config

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/uptime-cli/uptimectl/pkg/contextmanager"
)

// currentContextCmd represents the get command
var currentContextCmd = &cobra.Command{
	Use:   "current-context",
	Short: "Display the current-context",
	Args:  cobra.NoArgs,

	RunE: func(cmd *cobra.Command, args []string) error {
		currentContext := contextmanager.CurrentContext()
		fmt.Printf("%s\n", currentContext.Name)
		return nil
	},
}

func init() {
	ConfigCmd.AddCommand(currentContextCmd)
}
