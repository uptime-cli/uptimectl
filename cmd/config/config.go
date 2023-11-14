package config

import (
	"github.com/spf13/cobra"
)

// ConfigCmd represents the get command
var ConfigCmd = &cobra.Command{
	Use:   "config",
	Short: "Manage config",
}

func init() {
}
