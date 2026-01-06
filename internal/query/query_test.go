package query

import (
	"testing"

	"github.com/Selyss/AssemBuddy/internal/model"
	"github.com/Selyss/AssemBuddy/internal/store"
)

func TestQueryAllArchOrder(t *testing.T) {
	st, err := store.Load()
	if err != nil {
		t.Fatalf("load failed: %v", err)
	}
	results, err := QueryByName(st, QueryOptions{
		Name:    "read",
		AllArch: true,
		Exact:   true,
	})
	if err != nil {
		t.Fatalf("query failed: %v", err)
	}
	if len(results) != 4 {
		t.Fatalf("expected 4 results, got %d", len(results))
	}
	if results[0].Arch != string(model.ArchArm) || results[1].Arch != string(model.ArchArm64) || results[2].Arch != string(model.ArchX64) || results[3].Arch != string(model.ArchX86) {
		t.Fatalf("unexpected arch order: %v", []string{results[0].Arch, results[1].Arch, results[2].Arch, results[3].Arch})
	}
}

func TestQuerySubstring(t *testing.T) {
	st, err := store.Load()
	if err != nil {
		t.Fatalf("load failed: %v", err)
	}
	results, err := QueryByName(st, QueryOptions{
		Name:    "ea",
		Arch:    model.ArchX64,
		Exact:   false,
		AllArch: false,
	})
	if err != nil {
		t.Fatalf("query failed: %v", err)
	}
	if len(results) == 0 {
		t.Fatalf("expected substring match")
	}
}

func TestQueryCaseSensitive(t *testing.T) {
	st, err := store.Load()
	if err != nil {
		t.Fatalf("load failed: %v", err)
	}
	results, err := QueryByName(st, QueryOptions{
		Name:          "Read",
		Arch:          model.ArchX64,
		Exact:         true,
		CaseSensitive: true,
	})
	if err != nil {
		t.Fatalf("query failed: %v", err)
	}
	if len(results) != 0 {
		t.Fatalf("expected no results, got %d", len(results))
	}
}
