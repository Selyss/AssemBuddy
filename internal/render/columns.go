package render

import (
	"fmt"
	"strings"
)

type ColumnKey string

const (
	ColumnArch       ColumnKey = "arch"
	ColumnNR         ColumnKey = "nr"
	ColumnName       ColumnKey = "name"
	ColumnReturn     ColumnKey = "return"
	ColumnReferences ColumnKey = "references"
	ColumnArg0       ColumnKey = "arg0"
	ColumnArg1       ColumnKey = "arg1"
	ColumnArg2       ColumnKey = "arg2"
	ColumnArg3       ColumnKey = "arg3"
	ColumnArg4       ColumnKey = "arg4"
	ColumnArg5       ColumnKey = "arg5"
)

var ValidColumns = map[ColumnKey]struct{}{
	ColumnArch:       {},
	ColumnNR:         {},
	ColumnName:       {},
	ColumnReturn:     {},
	ColumnReferences: {},
	ColumnArg0:       {},
	ColumnArg1:       {},
	ColumnArg2:       {},
	ColumnArg3:       {},
	ColumnArg4:       {},
	ColumnArg5:       {},
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
