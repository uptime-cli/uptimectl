package monitors

import (
	"github.com/spf13/cobra"
)

// MonitorsCmd represents the monitors command
var MonitorsCmd = &cobra.Command{
	Use:     "monitors",
	Aliases: []string{"monitor"},
	Short:   "Manage monitors",
}

func init() {
}
