package contextmanager

import (
	"errors"
	"fmt"
	"io/fs"
	"os"

	"github.com/mitchellh/go-homedir"
)

const (
	DefaultConfigFilename = ".uptimectl.yaml"
)

var (
	globalContextManager ContextManager

	OrganisationFlag string
)

func Init() {
	configFilename := os.Getenv("UPTIMECONFIG")
	if configFilename == "" {
		configFilename = os.Getenv("UPTIME_CONFIG")
	}
	if configFilename == "" {
		home, err := homedir.Dir()
		if err != nil {
			fmt.Printf("Failed to initialize context: %v\n", err)
			os.Exit(1)
		}
		configFilename = fmt.Sprintf("%s/%s", home, DefaultConfigFilename)
	}

	// TODO: add support for merging multiple config files

	globalContextManager = NewConfigFileContextManager(configFilename)
	err := globalContextManager.Load()
	if errors.Is(err, fs.ErrNotExist) {
		// ignore
	} else if err != nil {
		fmt.Printf("Failed to initialize context: %v\n", err)
		os.Exit(1)
	}
	return
}

func GlobalContextManager() ContextManager {
	return globalContextManager
}

func CurrentContext() Context {
	return globalContextManager.CurrentContext()
}

func SetCurrentContext(name string) error {
	return globalContextManager.SetCurrentContext(name)
}

func AddOrMergeContext(context Context) error {
	return globalContextManager.AddOrMergeContext(context)
}

func Load() error {
	return globalContextManager.Load()
}

func Save() error {
	return globalContextManager.Save()
}

// --------------------

func Organisation() string {
	if OrganisationFlag != "" {
		return OrganisationFlag
	}

	organisation := globalContextManager.CurrentContext().Organisation
	if organisation == "" {
		_, _ = fmt.Fprintf(os.Stderr, "no organisation set\n")
		os.Exit(1)
	}
	return organisation
}

func APIEndpoint() string {
	return globalContextManager.CurrentContext().API.API.URL
}

func BetteruptimeToken() string {
	return globalContextManager.CurrentContext().User.User.BetterUptimeToken
}
