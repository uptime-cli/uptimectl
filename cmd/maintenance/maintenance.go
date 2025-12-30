package maintenance

import (
	"github.com/spf13/cobra"
)

// MaintenanceCmd represents the incidents command
var MaintenanceCmd = &cobra.Command{
	Use:     "maintenance",
	Aliases: []string{"maint", "maintenances"},
	Short:   "Manage maintenance windows for monitors",
}

func init() {
}
