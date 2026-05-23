package ui

import (
	"strings"
	"testing"

	tea "github.com/charmbracelet/bubbletea"

	"cronscope/internal/cron"
)

func newTestTagStore() *cron.TagStore {
	ts := cron.NewTagStore("")
	_ = ts.Add("daily", "0 9 * * *")
	_ = ts.Add("weekly", "0 9 * * 1")
	return ts
}

func TestTagModel_DefaultNotVisible(t *testing.T) {
	m := NewTagModel(newTestTagStore())
	if m.Visible() {
		t.Error("expected not visible by default")
	}
}

func TestTagModel_Toggle_ShowsView(t *testing.T) {
	m := NewTagModel(newTestTagStore())
	m.Toggle()
	if !m.Visible() {
		t.Error("expected visible after toggle")
	}
}

func TestTagModel_View_HiddenReturnsEmpty(t *testing.T) {
	m := NewTagModel(newTestTagStore())
	if got := m.View(); got != "" {
		t.Errorf("expected empty view, got %q", got)
	}
}

func TestTagModel_View_ShowsLabels(t *testing.T) {
	m := NewTagModel(newTestTagStore())
	m.Toggle()
	view := m.View()
	if !strings.Contains(view, "daily") {
		t.Error("expected 'daily' in view")
	}
	if !strings.Contains(view, "weekly") {
		t.Error("expected 'weekly' in view")
	}
}

func TestTagModel_Navigation_DownUp(t *testing.T) {
	m := NewTagModel(newTestTagStore())
	m.Toggle()
	down := tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("j")}
	m, _ = m.Update(down)
	if m.SelectedLabel() == "" {
		t.Error("expected a selected label after navigation")
	}
	up := tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune("k")}
	m, _ = m.Update(up)
	if m.selected != 0 {
		t.Errorf("expected selected=0 after up, got %d", m.selected)
	}
}

func TestTagModel_EscHidesPanel(t *testing.T) {
	m := NewTagModel(newTestTagStore())
	m.Toggle()
	esc := tea.KeyMsg{Type: tea.KeyEsc}
	m, _ = m.Update(esc)
	if m.Visible() {
		t.Error("expected panel hidden after esc")
	}
}

func TestTagModel_View_EmptyStore(t *testing.T) {
	m := NewTagModel(cron.NewTagStore(""))
	m.Toggle()
	if !strings.Contains(m.View(), "no tags") {
		t.Error("expected 'no tags' message for empty store")
	}
}
