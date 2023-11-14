package betteruptime

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/go-resty/resty/v2"

	"github.com/uptime-cli/uptimectl/pkg/contextmanager"
	versionpkg "github.com/uptime-cli/uptimectl/pkg/version"
)

type client struct {
	rest *resty.Client
}

func NewClient() client {
	restyclient := resty.New()

	token := contextmanager.BetteruptimeToken()
	if token == "" {
		envToken, found := os.LookupEnv("BETTERUPTIME_TOKEN")
		if !found {
			fmt.Printf("error: no access token configured. Please log in or set the `BETTERUPTIME_TOKEN` environment variable")
			os.Exit(1)
		}
		token = envToken
	}

	// bearer token
	restyclient.SetAuthToken(token)

	restyclient.SetHeaders(map[string]string{
		"User-Agent": fmt.Sprintf("uptime-cli/uptimectl/%s", versionpkg.Version()),
	})
	// restyclient.Debug = true

	// Configure retry mechanism
	restyclient.
		// Set retry count to non zero to enable retries
		SetRetryCount(3).
		// You can override initial retry wait time.
		// Default is 100 milliseconds.
		SetRetryWaitTime(5 * time.Second).
		// MaxWaitTime can be overridden as well.
		// Default is 2 seconds.
		SetRetryMaxWaitTime(20 * time.Second).
		// SetRetryAfter sets callback to calculate wait time between retries.
		// Default (nil) implies exponential backoff with jitter
		SetRetryAfter(func(client *resty.Client, resp *resty.Response) (time.Duration, error) {
			return 0, errors.New("quota exceeded")
		})

	return client{
		rest: restyclient,
	}
}

type Pagination struct {
	First string
	Last  string
	Prev  *string
	Next  *string
}
