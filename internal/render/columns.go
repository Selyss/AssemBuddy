package render

import (
	"fmt"
	"strings"
)

type ColumnKey string

const (
	ColumnArch      ColumnKey = "arch"
	ColumnName      ColumnKey = "name"
	ColumnNumber    ColumnKey = "number"
	ColumnABI       ColumnKey = "abi"
	ColumnInstr     ColumnKey = "instr"
	ColumnSignature ColumnKey = "signature"
	ColumnNotes     ColumnKey = "notes"
	ColumnSince     ColumnKey = "since"
)

var ValidColumns = map[ColumnKey]struct{}{
	ColumnArch:      {},
	ColumnName:      {},
	ColumnNumber:    {},
	ColumnABI:       {},
	ColumnInstr:     {},
	ColumnSignature: {},
	ColumnNotes:     {},
	ColumnSince:     {},
}

func ParseColumns(input string) ([]ColumnKey, error) {
	if strings.TrimSpace(input) == "" {
		return nil, nil
	}
	parts := strings.Split(input, ",")
	cols := make([]ColumnKey, 0, len(parts))
	for _, part := range parts {
		key := ColumnKey(strings.ToLower(strings.TrimSpace(part)))
		if key == "" {
			continue
		}
		if _, ok := ValidColumns[key]; !ok {
			return nil, fmt.Errorf("invalid column: %s", part)
		}
		cols = append(cols, key)
	}
	if len(cols) == 0 {
		return nil, fmt.Errorf("no columns selected")
	}
	return cols, nil
}
