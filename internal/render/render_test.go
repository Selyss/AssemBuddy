package render

import (
	"strings"
	"testing"

	"github.com/Selyss/AssemBuddy/internal/model"
)

func TestParseColumnsInvalid(t *testing.T) {
	_, err := ParseColumns("arch,unknown")
	if err == nil {
		t.Fatalf("expected error for invalid column")
	}
}

func TestTableColumnOrder(t *testing.T) {
	records := []model.SyscallRecord{{Name: "read", Number: 0, Arch: "x64"}}
	cols := []ColumnKey{ColumnNR, ColumnName, ColumnArch}
	output := RenderTable(records, cols, 0)

	idxNumber := strings.Index(output, "NR")
	idxName := strings.Index(output, "NAME")
	idxArch := strings.Index(output, "ARCH")
	if idxNumber == -1 || idxName == -1 || idxArch == -1 {
		t.Fatalf("missing headers in output: %s", output)
	}
	if !(idxNumber < idxName && idxName < idxArch) {
		t.Fatalf("unexpected header order: %d %d %d", idxNumber, idxName, idxArch)
	}
}

func TestJSONOutput(t *testing.T) {
	records := []model.SyscallRecord{{Name: "read", Number: 0, Arch: "x64"}}
	payload, err := RenderJSONRecords(records)
	if err != nil {
		t.Fatalf("json render failed: %v", err)
	}
	if !strings.Contains(payload, "\"nr\"") {
		t.Fatalf("expected json output")
	}
}

func TestPagerSuppression(t *testing.T) {
	if shouldUsePager(10, 20, "json", false, true) {
		t.Fatalf("expected pager off for json")
	}
	if shouldUsePager(10, 20, "table", true, true) {
		t.Fatalf("expected pager off for no-pager")
	}
	if shouldUsePager(10, 0, "table", false, true) {
		t.Fatalf("expected pager off for unknown height")
	}
	if shouldUsePager(5, 10, "table", false, true) {
		t.Fatalf("expected pager off when output fits")
	}
	if !shouldUsePager(20, 10, "table", false, true) {
		t.Fatalf("expected pager on when output exceeds")
	}
}
