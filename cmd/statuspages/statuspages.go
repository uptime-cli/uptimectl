package statuspages

import (
	"github.com/spf13/cobra"
)

// StatusPagesCmd represents the incidents command
var StatusPagesCmd = &cobra.Command{
	Use:     "status-pages",
	Aliases: []string{"statuspage", "statuspages"},
	Short:   "Manage status pages",
}

func init() {
}
