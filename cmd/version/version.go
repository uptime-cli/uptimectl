package version

import (
	"github.com/spf13/cobra"

	versionpkg "github.com/uptime-cli/uptimectl/pkg/version"
)

var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version information",
	Long:  `version information`,

	Run: func(cmd *cobra.Command, args []string) {
		versionpkg.Print()
	},
}
