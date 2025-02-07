package command

import (
	"github.com/spf13/cobra"

	"github.com/khulnasoft/gob"
)

func Root(app gob.Application) *cobra.Command {
	return app.SetupRootCommand(&cobra.Command{})
}
