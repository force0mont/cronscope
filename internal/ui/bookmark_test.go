package ui

import (
	"path/filepath"
	"strings"
	"testing"
	"time"

	tea "github.com/charmbracelet/bubbletea"

	"cronscope/internal/cron"
)

func newTestBookmarkStore(t *testing.T) *cron.BookmarkStore {
	t.Helper()
	path := filepath.Join(t.TempDir(), "bookmarks.json")
	s, err := cron.NewBookmarkStore(path)
	if err != nil {
		t.Fatalf("failed to create bookmark store: %v", err)
	}
	return s
}

func TestBookmarkModel_DefaultNotVisible(t *testing.T) {
	s := newTestBookmarkStore(t)
	m := NewBookmarkModel(s)
	if m.View() != "" {
		t.Error("expected empty view when not visible")
	}
}

func TestBookmarkModel_Toggle_ShowsView(t *testing.T) {
	s := newTestBookmarkStore(t)
	m := NewBookmarkModel(s)
	m.Toggle()
	if m.View() == "" {
		t.Error("expected non-empty view after toggle")
	}
}

func TestBookmarkModel_View_ShowsItems(t *testing.T) {
	s := newTestBookmarkStore(t)
	s.Add(cron.Bookmark{Label: "daily", Expr: "0 9 * * *", CreatedAt: time.Now()})
	m := NewBookmarkModel(s)
	m.Toggle()
	view := m.View()
	if !strings.Contains(view, "daily") || !strings.Contains(view, "0 9 * * *") {
		t.Errorf("view missing bookmark data: %s", view)
	}
}

func TestBookmarkModel_View_EmptyMessage(t *testing.T) {
	s := newTestBookmarkStore(t)
	m := NewBookmarkModel(s)
	m.Toggle()
	if !strings.Contains(m.View(), "no bookmarks") {
		t.Error("expected empty-state message")
	}
}

func TestBookmarkModel_Navigation_DownUp(t *testing.T) {
	s := newTestBookmarkStore(t)
	s.Add(cron.Bookmark{Label: "a", Expr: "0 1 * * *"})
	s.Add(cron.Bookmark{Label: "b", Expr: "0 2 * * *"})
	m := NewBookmarkModel(s)
	m.Toggle()
	m, _ = m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("j")})
	if m.cursor != 1 {
		t.Errorf("expected cursor 1, got %d", m.cursor)
	}
	m, _ = m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("k")})
	if m.cursor != 0 {
		t.Errorf("expected cursor 0, got %d", m.cursor)
	}
}

func TestBookmarkModel_Selected_ReturnsHighlighted(t *testing.T) {
	s := newTestBookmarkStore(t)
	s.Add(cron.Bookmark{Label: "first", Expr: "* * * * *"})
	m := NewBookmarkModel(s)
	m.Toggle()
	got := m.Selected()
	if got == nil || got.Label != "first" {
		t.Errorf("expected 'first' bookmark, got %+v", got)
	}
}

func TestBookmarkModel_Selected_NilWhenEmpty(t *testing.T) {
	s := newTestBookmarkStore(t)
	m := NewBookmarkModel(s)
	m.Toggle()
	if m.Selected() != nil {
		t.Error("expected nil selected for empty store")
	}
}
