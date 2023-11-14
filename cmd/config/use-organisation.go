package config

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/uptime-cli/uptimectl/cmd/incidents"
	"github.com/uptime-cli/uptimectl/pkg/contextmanager"
	"github.com/uptime-cli/uptimectl/pkg/fzf"
)

// useOrganisationCmd represents the use-organisation command
var useOrganisationCmd = &cobra.Command{
	Use:     "use-organisation <organisation>",
	Aliases: []string{"set-organisation"},
	Args:    cobra.RangeArgs(0, 1),
	Short:   "Set the organisation in the current-context",
	RunE: func(cmd *cobra.Command, args []string) error {
		currentContext := contextmanager.CurrentContext()
		selectedOrganisation, err := getSelectedOrganisation(args)
		if err != nil {
			return err
		}
		currentContext.Organisation = selectedOrganisation
		err = contextmanager.AddOrMergeContext(currentContext)
		if err != nil {
			return err
		}
		fmt.Println(currentContext.Organisation)
		return contextmanager.Save()
	},
}

func getSelectedOrganisation(args []string) (string, error) {
	if contextmanager.OrganisationFlag != "" {
		return contextmanager.OrganisationFlag, nil
	}

	if len(args) == 0 && fzf.IsInteractiveMode(os.Stdout) {
		command := fmt.Sprintf("%s config get-organisations --%s", os.Args[0], incidents.NoHeaderKey)
		return fzf.InteractiveChoice(command)
	} else if len(args) == 1 {
		return args[0], nil
	} else {
		return "", errors.New("invalid organisation")
	}
}

func init() {
	ConfigCmd.AddCommand(useOrganisationCmd)
}
