package incidents

import (
	"github.com/spf13/cobra"
)

// IncidentsCmd represents the incidents command
var IncidentsCmd = &cobra.Command{
	Use:     "incident",
	Aliases: []string{"incidents", "i"},
	Short:   "Manage incidents",
}

func init() {
}
