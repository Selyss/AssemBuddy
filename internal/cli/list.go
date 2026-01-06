package cli

import (
	"fmt"
	"strings"

	"github.com/Selyss/AssemBuddy/internal/query"
	"github.com/Selyss/AssemBuddy/internal/render"
	"github.com/spf13/cobra"
)

func newListCommand(app *App) *cobra.Command {
	var archInput string
	var filter string
	var columnsInput string

	cmd := &cobra.Command{
		Use:   "list",
		Short: "List syscalls for an architecture",
		RunE: func(cmd *cobra.Command, args []string) error {
			if strings.TrimSpace(archInput) == "" {
				return usageError(fmt.Errorf("--arch is required"))
			}
			arch, err := query.NormalizeArch(archInput)
			if err != nil {
				return usageError(err)
			}

			results, err := query.ListArch(app.Store, arch, filter, false)
			if err != nil {
				return usageError(err)
			}
			if len(results) == 0 {
				return notFoundError(fmt.Errorf("no syscalls found"))
			}

			cols, err := render.ParseColumns(columnsInput)
			if err != nil {
				return usageError(err)
			}
			if cols == nil {
				cols = []render.ColumnKey{render.ColumnNR, render.ColumnName}
			}

			format := strings.ToLower(app.Config.Format)
			if format == "json" {
				payload, err := render.RenderJSONRecords(results)
				if err != nil {
					return internalError(err)
				}
				fmt.Fprint(cmd.OutOrStdout(), payload)
				return nil
			}

			output := render.RenderTable(results, cols, app.Config.Width)
			return render.OutputWithPager(output, format, app.Config.NoPager)
		},
	}

	cmd.Flags().StringVarP(&archInput, "arch", "a", "", "Architecture")
	cmd.Flags().StringVar(&filter, "filter", "", "Filter by substring")
	cmd.Flags().StringVar(&columnsInput, "columns", "", "Comma-separated columns")

	return cmd
}
