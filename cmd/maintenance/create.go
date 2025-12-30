package maintenance

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
	"github.com/uptime-cli/uptimectl/pkg/betteruptime"
	"k8s.io/utils/ptr"
)

var (
	maintenanceFrom     string
	maintenanceDuration time.Duration
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "create a maintenance window for a monitor",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		client := betteruptime.NewClient()

		monitor, err := client.GetMonitor(cmd.Context(), args[0])
		if err != nil {
			return err
		}

		// err = client.CreateMaintenanceWindow(cmd.Context(), monitor.Id, maintenanceFrom, maintenanceDuration)
		// if err != nil {
		// 	return err
		// }

		currentDate := time.Now().Format("2006-01-02")
		from, err := time.Parse("2006-01-02 15:04:05", fmt.Sprintf("%s %s", currentDate, maintenanceFrom))
		if err != nil {
			return err
		}
		monitor.Attributes.MaintenanceFrom = &from
		monitor.Attributes.MaintenanceTo = ptr.To(from.Add(maintenanceDuration))
		monitor.Attributes.MaintenanceTimezone = from.Location().String()
		monitor.Attributes.MaintenanceDays = []string{"monday", "tuesday", "wednesday", "thursday", "friday", "saturday", "sunday"}

		if err := client.UpdateMonitor(cmd.Context(), monitor); err != nil {
			return err
		}
		fmt.Printf("Maintenance window created for monitor %s\n", monitor.Attributes.PronounceableName)
		return nil
	},
}

func init() {
	MaintenanceCmd.AddCommand(createCmd)

	createCmd.Flags().StringVarP(&maintenanceFrom, "from", "f", time.Now().Format("15:04:05"), "start time of the maintenance window")
	createCmd.Flags().DurationVarP(&maintenanceDuration, "duration", "d", time.Hour, "duration of the maintenance window")
}
