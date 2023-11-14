package completion

import (
	"strings"

	"github.com/spf13/cobra"

	"github.com/uptime-cli/uptimectl/pkg/contextmanager"
)

func ContextCompletionFunc(allowMultiple bool) func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
	return func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		if len(args) != 0 && !allowMultiple {
			return nil, cobra.ShellCompDirectiveNoFileComp
		}
		contextNames := []string{}
		for _, context := range contextmanager.GlobalContextManager().Config().Contexts {
			if strings.HasPrefix(context.Name, toComplete) {
				contextNames = append(contextNames, context.Name)
			}
		}
		return contextNames, cobra.ShellCompDirectiveNoFileComp
	}
}
