package cron

import (
	"os"
	"path/filepath"
	"testing"
)

func tempHistoryPath(t *testing.T) string {
	t.Helper()
	dir := t.TempDir()
	return filepath.Join(dir, "history.json")
}

func TestHistory_Add_DeduplicatesEntries(t *testing.T) {
	h := NewHistory(tempHistoryPath(t), 10)
	h.Add("0 * * * *")
	h.Add("0 * * * *")
	if len(h.Entries) != 1 {
		t.Fatalf("expected 1 entry, got %d", len(h.Entries))
	}
}

func TestHistory_Add_RespectsMaxSize(t *testing.T) {
	h := NewHistory(tempHistoryPath(t), 3)
	h.Add("0 * * * *")
	h.Add("5 * * * *")
	h.Add("10 * * * *")
	h.Add("15 * * * *")
	if len(h.Entries) != 3 {
		t.Fatalf("expected 3 entries, got %d", len(h.Entries))
	}
}

func TestHistory_Add_MostRecentFirst(t *testing.T) {
	h := NewHistory(tempHistoryPath(t), 10)
	h.Add("0 * * * *")
	h.Add("5 * * * *")
	if h.Entries[0].Expression != "5 * * * *" {
		t.Errorf("expected newest entry first, got %q", h.Entries[0].Expression)
	}
}

func TestHistory_SaveAndLoad_RoundTrip(t *testing.T) {
	path := tempHistoryPath(t)
	h := NewHistory(path, 10)
	h.Add("0 * * * *")
	h.Add("@daily")
	if err := h.Save(); err != nil {
		t.Fatalf("Save: %v", err)
	}

	h2 := NewHistory(path, 10)
	if err := h2.Load(); err != nil {
		t.Fatalf("Load: %v", err)
	}
	if len(h2.Entries) != 2 {
		t.Fatalf("expected 2 entries after reload, got %d", len(h2.Entries))
	}
	if h2.Entries[0].Expression != "@daily" {
		t.Errorf("expected @daily first, got %q", h2.Entries[0].Expression)
	}
}

func TestHistory_Load_MissingFileIsEmpty(t *testing.T) {
	h := NewHistory(filepath.Join(t.TempDir(), "no_such.json"), 10)
	if err := h.Load(); err != nil {
		t.Fatalf("unexpected error for missing file: %v", err)
	}
	if len(h.Entries) != 0 {
		t.Errorf("expected empty history, got %d entries", len(h.Entries))
	}
}

func TestHistory_Expressions_ReturnsStrings(t *testing.T) {
	h := NewHistory(tempHistoryPath(t), 10)
	h.Add("0 * * * *")
	h.Add("@weekly")
	exprs := h.Expressions()
	if len(exprs) != 2 || exprs[0] != "@weekly" {
		t.Errorf("unexpected expressions: %v", exprs)
	}
	_ = os.Remove(h.path)
}
