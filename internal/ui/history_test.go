package ui

import (
	"strings"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
)

func TestHistoryModel_DefaultNotVisible(t *testing.T) {
	m := NewHistoryModel([]string{"0 * * * *"})
	if m.IsVisible() {
		t.Error("expected history to be hidden by default")
	}
}

func TestHistoryModel_Toggle(t *testing.T) {
	m := NewHistoryModel(nil)
	m.Toggle()
	if !m.IsVisible() {
		t.Error("expected history to be visible after toggle")
	}
	m.Toggle()
	if m.IsVisible() {
		t.Error("expected history to be hidden after second toggle")
	}
}

func TestHistoryModel_View_HiddenReturnsEmpty(t *testing.T) {
	m := NewHistoryModel([]string{"0 * * * *"})
	if m.View() != "" {
		t.Error("expected empty view when hidden")
	}
}

func TestHistoryModel_View_ShowsItems(t *testing.T) {
	m := NewHistoryModel([]string{"0 * * * *", "@daily"})
	m.Toggle()
	view := m.View()
	if !strings.Contains(view, "0 * * * *") {
		t.Error("expected view to contain first expression")
	}
	if !strings.Contains(view, "@daily") {
		t.Error("expected view to contain second expression")
	}
}

func TestHistoryModel_Navigation_DownUp(t *testing.T) {
	m := NewHistoryModel([]string{"a", "b", "c"})
	m.Toggle()
	down := tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'j'}}
	m, _ = m.Update(down)
	if m.cursor != 1 {
		t.Errorf("expected cursor 1, got %d", m.cursor)
	}
	up := tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'k'}}
	m, _ = m.Update(up)
	if m.cursor != 0 {
		t.Errorf("expected cursor 0, got %d", m.cursor)
	}
}

func TestHistoryModel_Enter_EmitsSelected(t *testing.T) {
	m := NewHistoryModel([]string{"0 * * * *"})
	m.Toggle()
	enter := tea.KeyMsg{Type: tea.KeyEnter}
	_, cmd := m.Update(enter)
	if cmd == nil {
		t.Fatal("expected a command on enter")
	}
	msg := cmd()
	sel, ok := msg.(HistorySelectedMsg)
	if !ok {
		t.Fatalf("expected HistorySelectedMsg, got %T", msg)
	}
	if sel.Expression != "0 * * * *" {
		t.Errorf("unexpected expression: %q", sel.Expression)
	}
}
