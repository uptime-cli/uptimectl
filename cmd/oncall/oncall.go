package oncall

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	"github.com/uptime-cli/uptimectl/pkg/betteruptime"
	"github.com/uptime-cli/uptimectl/pkg/table"
)

const NoHeaderKey = "no-header"

var noHeader bool

// OncallCmd represents the get command
var OncallCmd = &cobra.Command{
	Use:     "on-call",
	Short:   "Get all current on calls",
	Aliases: []string{},
	Args:    cobra.MaximumNArgs(0),
	RunE: func(cmd *cobra.Command, args []string) error {
		client := betteruptime.NewClient()

		onCalls, err := client.GetOnCall()
		if err != nil {
			return fmt.Errorf("failed to fetch on call engineers")
		}

		body := make([][]string, 0, len(onCalls.Data))

		for _, oncallSchedules := range onCalls.Data {
			onCallUsersInSchedule := []betteruptime.UserReferences{}

			for _, ref := range oncallSchedules.Relationships.OnCallUsers.Data {
				for _, userReference := range onCalls.Included {
					if ref.Id == userReference.Id {
						onCallUsersInSchedule = append(onCallUsersInSchedule, userReference)
					}
				}
			}

			for _, userReference := range onCallUsersInSchedule {
				body = append(body, []string{
					oncallSchedules.Attributes.Name,
					fmt.Sprint(oncallSchedules.Attributes.DefaultCalendar),
					userReference.Attributes.FirstName + " " + userReference.Attributes.LastName,
					userReference.Attributes.Email,
					strings.Join(userReference.Attributes.PhoneNumbers, ","),
				})
			}
		}

		if noHeader {
			table.Print(nil, body)
		} else {
			table.Print([]string{"Schedule", "Default", "Name", "Email", "Phone"}, body)
		}
		return nil
	},
}

func init() {
	OncallCmd.Flags().BoolVar(&noHeader, NoHeaderKey, false, "Do not print the header")
}
