package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/uptime-cli/uptimectl/cmd/auth"
	"github.com/uptime-cli/uptimectl/cmd/config"
	"github.com/uptime-cli/uptimectl/cmd/incidents"
	"github.com/uptime-cli/uptimectl/cmd/monitorgroups"
	"github.com/uptime-cli/uptimectl/cmd/monitors"
	"github.com/uptime-cli/uptimectl/cmd/oncall"
	"github.com/uptime-cli/uptimectl/cmd/statuspages"
	"github.com/uptime-cli/uptimectl/cmd/version"
	"github.com/uptime-cli/uptimectl/pkg/authmanager"
	"github.com/uptime-cli/uptimectl/pkg/contextmanager"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "uptimectl",
	Short: "A command-line interface for working with Better Uptime",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		switch err {
		case authmanager.ErrNoLogin:
			fallthrough
		case authmanager.ErrSessionNotActive:
			_, _ = fmt.Fprintf(os.Stderr, "%v, please use: uptimectl auth login\n", err)
		default:
			_, _ = fmt.Fprintf(os.Stderr, "failed: %v\n", err)
		}
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(auth.AuthCmd)
	rootCmd.AddCommand(incidents.IncidentsCmd)
	rootCmd.AddCommand(monitorgroups.MonitorGroupsCmd)
	rootCmd.AddCommand(statuspages.StatusPagesCmd)
	rootCmd.AddCommand(config.ConfigCmd)
	rootCmd.AddCommand(version.VersionCmd)
	rootCmd.AddCommand(oncall.OncallCmd)
	rootCmd.AddCommand(monitors.MonitorsCmd)
	cobra.OnInitialize(contextmanager.Init)
}
