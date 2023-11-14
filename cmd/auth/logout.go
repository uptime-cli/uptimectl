package auth

import (
	"github.com/spf13/cobra"

	"github.com/uptime-cli/uptimectl/pkg/authmanager"
)

// logout represents the logout command
var logoutCmd = &cobra.Command{
	Use:   "logout",
	Short: "Remove credentials from config",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		return authmanager.NewAuthManager().Logout()
	},
}

func init() {
	AuthCmd.AddCommand(logoutCmd)
}
