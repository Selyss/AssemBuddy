package cli

import (
	"fmt"
	"strings"

	"github.com/Selyss/AssemBuddy/internal/model"
	"github.com/Selyss/AssemBuddy/internal/query"
	"github.com/Selyss/AssemBuddy/internal/render"
	"github.com/spf13/cobra"
)

func newQueryCommand(app *App) *cobra.Command {
	var name string
	var archInput string
	var allArch bool
	var caseSensitive bool
	var exact bool
	var columnsInput string

	cmd := &cobra.Command{
		Use:   "query",
		Short: "Query syscall information by name",
		RunE: func(cmd *cobra.Command, args []string) error {
			if strings.TrimSpace(name) == "" {
				return usageError(fmt.Errorf("--name is required"))
			}
			if allArch && archInput != "" {
				return usageError(fmt.Errorf("--arch and --all-arch are mutually exclusive"))
			}
			if !allArch && archInput == "" {
				return usageError(fmt.Errorf("--arch or --all-arch is required"))
			}

			var arch model.Arch
			if !allArch {
				parsed, err := query.NormalizeArch(archInput)
				if err != nil {
					return usageError(err)
				}
				arch = parsed
			}

			opts := query.QueryOptions{
				Name:          name,
				Arch:          arch,
				AllArch:       allArch,
				Exact:         exact,
				CaseSensitive: caseSensitive,
			}
			results, err := query.QueryByName(app.Store, opts)
			if err != nil {
				return usageError(err)
			}
			if len(results) == 0 {
				return notFoundError(fmt.Errorf("syscall not found"))
			}

			cols, err := render.ParseColumns(columnsInput)
			if err != nil {
				return usageError(err)
			}
			if cols == nil {
				if allArch {
					cols = []render.ColumnKey{render.ColumnArch, render.ColumnNR, render.ColumnName}
				} else {
					cols = []render.ColumnKey{
						render.ColumnArch,
						render.ColumnName,
						render.ColumnNR,
						render.ColumnReturn,
						render.ColumnArg0,
						render.ColumnArg1,
						render.ColumnArg2,
						render.ColumnArg3,
						render.ColumnArg4,
						render.ColumnArg5,
						render.ColumnReferences,
					}
				}
			}

			format := strings.ToLower(app.Config.Format)
			if format == "json" {
				if len(results) == 1 && !allArch {
					payload, err := render.RenderJSONRecord(results[0])
					if err != nil {
						return internalError(err)
					}
					fmt.Fprint(cmd.OutOrStdout(), payload)
					return nil
				}
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

	cmd.Flags().StringVarP(&name, "name", "n", "", "Syscall name")
	cmd.Flags().StringVarP(&archInput, "arch", "a", "", "Architecture")
	cmd.Flags().BoolVar(&allArch, "all-arch", false, "Query all architectures")
	cmd.Flags().BoolVar(&caseSensitive, "case-sensitive", false, "Case sensitive match")
	cmd.Flags().BoolVar(&exact, "exact", true, "Exact match")
	cmd.Flags().StringVar(&columnsInput, "columns", "", "Comma-separated columns")

	return cmd
}
