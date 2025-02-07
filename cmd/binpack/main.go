package main

import (
	"github.com/khulnasoft/binpack/cmd/binpack/cli"
	"github.com/khulnasoft/gob"
)

// applicationName is the non-capitalized name of the application (do not change this)
const (
	applicationName = "binpack"
	notProvided     = "[not provided]"
)

// all variables here are provided as build-time arguments, with clear default values
var (
	version        = notProvided
	buildDate      = notProvided
	gitCommit      = notProvided
	gitDescription = notProvided
)

func main() {
	app := cli.New(
		gob.Identification{
			Name:           applicationName,
			Version:        version,
			BuildDate:      buildDate,
			GitCommit:      gitCommit,
			GitDescription: gitDescription,
		},
	)

	app.Run()
}
