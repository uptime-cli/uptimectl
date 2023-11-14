package auth

import (
	"context"
	"net/url"

	"github.com/spf13/cobra"

	"github.com/uptime-cli/uptimectl/pkg/authmanager"
	"github.com/uptime-cli/uptimectl/pkg/contextmanager"
)

var (
	// idp settings
	token string
	// betteruptime api
	apiURL string

	// context creation options
	createContext bool
	contextName   string
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "Login to betteruptime and store credentials in config file",
	Args:  cobra.NoArgs,
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := context.Background()
		manager := authmanager.NewAuthManager()

		// TODO: Should context creation be removed from login cmd?
		// Maybe a new dedicated cmd
		if createContext {
			err := createNewCurrentContext()
			if err != nil {
				return err
			}
		}
		return manager.Login(ctx, token)
	},
}

func createNewCurrentContext() error {
	u, err := url.Parse(apiURL)
	if err != nil {
		return err
	}

	defaultContext := newDefaultContext(contextName, contextmanager.OrganisationFlag, u.Host)
	err = contextmanager.GlobalContextManager().AddOrMergeContext(defaultContext)
	if err != nil {
		return err
	}
	err = contextmanager.SetCurrentContext(defaultContext.Name)
	if err != nil {
		return err
	}
	return nil
}

func init() {
	AuthCmd.AddCommand(loginCmd)

	loginCmd.Flags().StringVar(&token, "token", "", "api access token for betteruptime")
	loginCmd.Flags().StringVar(&apiURL, "api-url", "https://betteruptime.com", "specify a custom api url")

	loginCmd.Flags().BoolVar(&createContext, "create-context", false, "creates a context")
	loginCmd.Flags().StringVar(&contextName, "context-name", "default", "name of the context")
}

func newDefaultContext(contextName, organisation, apiName string) contextmanager.Context {
	return contextmanager.Context{
		Name:         contextName,
		Organisation: organisation,
		API: contextmanager.APIs{
			Name: apiName,
			API: contextmanager.API{
				URL: apiURL,
			},
		},
		User: contextmanager.Users{
			Name: "accesstoken",
			User: contextmanager.User{
				BetterUptimeToken: token,
			},
		},
	}
}
