package monitorgroups

import (
	"github.com/spf13/cobra"
)

// MonitorGroupsCmd represents the incidents command
var MonitorGroupsCmd = &cobra.Command{
	Use:     "monitor-groups",
	Aliases: []string{"monitorgroups", "monitor-group", "monitorgroup"},
	Short:   "Manage monitor groups",
}

func init() {
}
