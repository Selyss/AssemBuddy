package cli

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

func newVersionCommand(app *App, version string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "version",
		Short: "Print version information",
		RunE: func(cmd *cobra.Command, args []string) error {
			format := strings.ToLower(app.Config.Format)
			if format == "json" {
				payload := fmt.Sprintf("{\"version\":%q,\"dataset_version\":%q}", version, app.Store.Meta.DatasetVersion)
				fmt.Fprint(cmd.OutOrStdout(), payload)
				return nil
			}
			fmt.Fprintf(cmd.OutOrStdout(), "assembuddy %s\n", version)
			fmt.Fprintf(cmd.OutOrStdout(), "dataset %s\n", app.Store.Meta.DatasetVersion)
			return nil
		},
	}

	return cmd
}
