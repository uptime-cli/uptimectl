package main

import (
	"math/rand"
	"time"

	"github.com/uptime-cli/uptimectl/cmd"
	versionpkg "github.com/uptime-cli/uptimectl/pkg/version"
)

// these must be set by the compiler using LDFLAGS
// -X main.version= -X main.commit= -X main.date= -X main.builtBy=
var (
	version string
	commit  string
	date    string
	builtBy string
)

func main() {
	// make sure we have seed the rand package
	rand.Seed(time.Now().UnixNano())

	// execute Cobra root cmd
	cmd.Execute()
}

func init() {
	versionpkg.Init(version, commit, date, builtBy)
}
