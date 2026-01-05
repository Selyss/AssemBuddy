package store

import "testing"

func TestLoad(t *testing.T) {
	st, err := Load()
	if err != nil {
		t.Fatalf("load failed: %v", err)
	}
	if st.Meta.DatasetVersion == "" {
		t.Fatalf("expected dataset version")
	}
	if len(st.ByArch) != 4 {
		t.Fatalf("expected 4 arch datasets, got %d", len(st.ByArch))
	}
	if len(st.ByArchName) != 4 {
		t.Fatalf("expected 4 arch name maps, got %d", len(st.ByArchName))
	}
	if len(st.AllByName) == 0 {
		t.Fatalf("expected name index")
	}
}
