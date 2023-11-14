package auth

import (
	"github.com/spf13/cobra"
)

// AuthCmd represents the login command
var AuthCmd = &cobra.Command{
	Use:   "auth",
	Short: "Manage auth",
}
