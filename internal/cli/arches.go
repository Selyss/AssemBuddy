package cli

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/Selyss/AssemBuddy/internal/model"
	"github.com/Selyss/AssemBuddy/internal/query"
	"github.com/Selyss/AssemBuddy/internal/render"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
	"github.com/spf13/cobra"
)

type archInfo struct {
	Arch    string   `json:"arch"`
	Aliases []string `json:"aliases"`
}

func newArchesCommand(app *App) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "arches",
		Short: "List supported architectures and aliases",
		RunE: func(cmd *cobra.Command, args []string) error {
			format := strings.ToLower(app.Config.Format)
			aliases := query.ArchAliasTable()
			info := make([]archInfo, 0, len(model.CanonicalArchOrder))
			for _, arch := range model.CanonicalArchOrder {
				info = append(info, archInfo{Arch: string(arch), Aliases: aliases[arch]})
			}

			if format == "json" {
				payload, err := json.MarshalIndent(info, "", "  ")
				if err != nil {
					return internalError(err)
				}
				fmt.Fprint(cmd.OutOrStdout(), string(payload))
				return nil
			}

			writer := table.NewWriter()
			writer.SetStyle(table.StyleRounded)
			writer.SetColumnConfigs([]table.ColumnConfig{
				{Number: 1, Align: text.AlignLeft},
				{Number: 2, Align: text.AlignLeft},
			})
			writer.AppendHeader(table.Row{"ARCH", "ALIASES"})
			for _, entry := range info {
				writer.AppendRow(table.Row{entry.Arch, strings.Join(entry.Aliases, ", ")})
			}

			output := writer.Render()
			return render.OutputWithPager(output, format, app.Config.NoPager)
		},
	}

	return cmd
}
