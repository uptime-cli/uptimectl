package config

import (
	"fmt"

	"gopkg.in/yaml.v3"

	"github.com/spf13/cobra"

	"github.com/uptime-cli/uptimectl/pkg/contextmanager"
)

// viewCmd represents the get command
var viewCmd = &cobra.Command{
	Use:   "view",
	Short: "Shows current config",

	RunE: func(cmd *cobra.Command, args []string) error {
		manager := contextmanager.GlobalContextManager()
		config := manager.Config()
		data, err := yaml.Marshal(config)
		if err != nil {
			return err
		}
		fmt.Printf("%s\n", data)
		return nil
	},
}

func init() {
	ConfigCmd.AddCommand(viewCmd)
}
