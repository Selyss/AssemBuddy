package render

import (
	"fmt"
	"strings"

	"github.com/Selyss/AssemBuddy/internal/model"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

const truncationMarker = "..."

func RenderTable(records []model.SyscallRecord, columns []ColumnKey, width int) string {
	writer := table.NewWriter()
	writer.SetStyle(table.StyleRounded)
	writer.SetColumnConfigs(columnConfigs(columns))

	header := make(table.Row, 0, len(columns))
	for _, col := range columns {
		header = append(header, strings.ToUpper(string(col)))
	}
	writer.AppendHeader(header)

	for _, rec := range records {
		row := make(table.Row, 0, len(columns))
		for _, col := range columns {
			cell := columnValue(rec, col)
			if width > 0 {
				cell = truncate(cell, width)
			}
			row = append(row, cell)
		}
		writer.AppendRow(row)
	}

	return writer.Render()
}

func columnConfigs(columns []ColumnKey) []table.ColumnConfig {
	configs := make([]table.ColumnConfig, 0, len(columns))
	for idx, col := range columns {
		config := table.ColumnConfig{Number: idx + 1, Align: text.AlignLeft}
		if col == ColumnNR {
			config.Align = text.AlignRight
		}
		configs = append(configs, config)
	}
	return configs
}

func columnValue(rec model.SyscallRecord, col ColumnKey) string {
	switch col {
	case ColumnArch:
		return rec.Arch
	case ColumnName:
		return rec.Name
	case ColumnNR:
		return fmt.Sprintf("%d", rec.Number)
	case ColumnReturn:
		return rec.Return
	case ColumnReferences:
		return rec.References
	case ColumnArg0:
		return rec.Arg0
	case ColumnArg1:
		return rec.Arg1
	case ColumnArg2:
		return rec.Arg2
	case ColumnArg3:
		return rec.Arg3
	case ColumnArg4:
		return rec.Arg4
	case ColumnArg5:
		return rec.Arg5
	default:
		return ""
	}
}

func truncate(value string, width int) string {
	runes := []rune(value)
	if len(runes) <= width || width <= 0 {
		return value
	}
	if width <= len(truncationMarker) {
		return string(runes[:width])
	}
	return string(runes[:width-len(truncationMarker)]) + truncationMarker
}
