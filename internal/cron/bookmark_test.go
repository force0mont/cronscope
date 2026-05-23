package cron

import (
	"os"
	"path/filepath"
	"testing"
	"time"
)

func tempBookmarkPath(t *testing.T) string {
	t.Helper()
	return filepath.Join(t.TempDir(), "bookmarks.json")
}

func TestBookmarkStore_AddAndAll(t *testing.T) {
	s, err := NewBookmarkStore(tempBookmarkPath(t))
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	s.Add(Bookmark{Label: "daily", Expr: "0 9 * * *", CreatedAt: time.Now()})
	s.Add(Bookmark{Label: "weekly", Expr: "0 9 * * 1", CreatedAt: time.Now()})
	if got := len(s.All()); got != 2 {
		t.Errorf("expected 2 bookmarks, got %d", got)
	}
}

func TestBookmarkStore_Add_UpdatesExisting(t *testing.T) {
	s, _ := NewBookmarkStore(tempBookmarkPath(t))
	s.Add(Bookmark{Label: "daily", Expr: "0 9 * * *"})
	s.Add(Bookmark{Label: "daily", Expr: "0 10 * * *"})
	all := s.All()
	if len(all) != 1 {
		t.Fatalf("expected 1 bookmark, got %d", len(all))
	}
	if all[0].Expr != "0 10 * * *" {
		t.Errorf("expected updated expr, got %s", all[0].Expr)
	}
}

func TestBookmarkStore_Remove(t *testing.T) {
	s, _ := NewBookmarkStore(tempBookmarkPath(t))
	s.Add(Bookmark{Label: "daily", Expr: "0 9 * * *"})
	if !s.Remove("daily") {
		t.Error("expected Remove to return true")
	}
	if len(s.All()) != 0 {
		t.Error("expected empty store after remove")
	}
}

func TestBookmarkStore_Remove_NotFound(t *testing.T) {
	s, _ := NewBookmarkStore(tempBookmarkPath(t))
	if s.Remove("ghost") {
		t.Error("expected Remove to return false for missing label")
	}
}

func TestBookmarkStore_SaveAndLoad_RoundTrip(t *testing.T) {
	path := tempBookmarkPath(t)
	s, _ := NewBookmarkStore(path)
	s.Add(Bookmark{Label: "midnight", Expr: "0 0 * * *", Timezone: "UTC", CreatedAt: time.Now()})
	if err := s.Save(); err != nil {
		t.Fatalf("save failed: %v", err)
	}
	s2, err := NewBookmarkStore(path)
	if err != nil {
		t.Fatalf("load failed: %v", err)
	}
	all := s2.All()
	if len(all) != 1 || all[0].Label != "midnight" {
		t.Errorf("round-trip mismatch: %+v", all)
	}
}

func TestNewBookmarkStore_MissingFile_OK(t *testing.T) {
	path := filepath.Join(t.TempDir(), "nonexistent.json")
	s, err := NewBookmarkStore(path)
	if err != nil {
		t.Fatalf("expected no error for missing file, got %v", err)
	}
	if len(s.All()) != 0 {
		t.Error("expected empty store")
	}
}

func TestNewBookmarkStore_CorruptFile_Error(t *testing.T) {
	path := tempBookmarkPath(t)
	_ = os.WriteFile(path, []byte("not json"), 0o644)
	_, err := NewBookmarkStore(path)
	if err == nil {
		t.Error("expected error for corrupt file")
	}
}
