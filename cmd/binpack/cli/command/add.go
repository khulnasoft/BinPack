package command

import (
	"github.com/spf13/cobra"

	"github.com/khulnasoft/gob"
)

func Add(app gob.Application) *cobra.Command {
	cmd := app.SetupCommand(&cobra.Command{
		Use:   "add",
		Short: "Add a new tool to the configuration",
	})

	cmd.AddCommand(
		AddGoInstall(app),
		AddGithubRelease(app),
	)

	return cmd
}
