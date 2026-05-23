package cron

import (
	"testing"
)

func TestTagStore_AddAndByLabel(t *testing.T) {
	ts := NewTagStore("")
	if err := ts.Add("daily", "0 9 * * *"); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	exprs := ts.ByLabel("daily")
	if len(exprs) != 1 || exprs[0] != "0 9 * * *" {
		t.Errorf("expected [0 9 * * *], got %v", exprs)
	}
}

func TestTagStore_Add_Deduplicates(t *testing.T) {
	ts := NewTagStore("")
	_ = ts.Add("weekly", "0 9 * * 1")
	_ = ts.Add("weekly", "0 9 * * 1")
	if got := len(ts.ByLabel("weekly")); got != 1 {
		t.Errorf("expected 1 entry, got %d", got)
	}
}

func TestTagStore_Add_EmptyLabel(t *testing.T) {
	ts := NewTagStore("")
	if err := ts.Add("  ", "* * * * *"); err == nil {
		t.Error("expected error for empty label")
	}
}

func TestTagStore_Remove(t *testing.T) {
	ts := NewTagStore("")
	_ = ts.Add("daily", "0 9 * * *")
	if err := ts.Remove("daily", "0 9 * * *"); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(ts.ByLabel("daily")) != 0 {
		t.Error("expected empty after remove")
	}
}

func TestTagStore_Remove_NotFound(t *testing.T) {
	ts := NewTagStore("")
	if err := ts.Remove("missing", "* * * * *"); err == nil {
		t.Error("expected error for missing label")
	}
}

func TestTagStore_Labels_Sorted(t *testing.T) {
	ts := NewTagStore("")
	_ = ts.Add("zebra", "* * * * *")
	_ = ts.Add("apple", "* * * * *")
	_ = ts.Add("mango", "* * * * *")
	labels := ts.Labels()
	expected := []string{"apple", "mango", "zebra"}
	for i, l := range labels {
		if l != expected[i] {
			t.Errorf("position %d: want %q, got %q", i, expected[i], l)
		}
	}
}

func TestTagStore_MultipleExpressionsPerLabel(t *testing.T) {
	ts := NewTagStore("")
	_ = ts.Add("work", "0 9 * * 1-5")
	_ = ts.Add("work", "0 17 * * 1-5")
	if got := len(ts.ByLabel("work")); got != 2 {
		t.Errorf("expected 2 expressions, got %d", got)
	}
}
