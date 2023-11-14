package config

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/uptime-cli/uptimectl/cmd/incidents"
	"github.com/uptime-cli/uptimectl/pkg/completion"
	"github.com/uptime-cli/uptimectl/pkg/contextmanager"
	"github.com/uptime-cli/uptimectl/pkg/fzf"
)

// useContextCmd represents the get command
var useContextCmd = &cobra.Command{
	Use:               "use-context <context>",
	Aliases:           []string{"set-context"},
	Args:              cobra.RangeArgs(0, 1),
	ValidArgsFunction: completion.ContextCompletionFunc(false),
	Short:             "Set the current-context",
	RunE: func(cmd *cobra.Command, args []string) error {
		context, err := getSelectedContext(args)
		if err != nil {
			return err
		}

		fmt.Println(context)
		err = contextmanager.SetCurrentContext(context)
		if err != nil {
			return err
		}
		return contextmanager.Save()
	},
}

func getSelectedContext(args []string) (string, error) {
	if len(args) == 0 && fzf.IsInteractiveMode(os.Stdout) {
		command := fmt.Sprintf("%s config get-contexts --%s", os.Args[0], incidents.NoHeaderKey)
		return fzf.InteractiveChoice(command)
	} else if len(args) == 1 {
		return args[0], nil
	} else {
		return "", errors.New("invalid context")
	}
}

func init() {
	ConfigCmd.AddCommand(useContextCmd)
}
