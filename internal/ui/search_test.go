package ui

import (
	"strings"
	"testing"

	tea "github.com/charmbracelet/bubbletea"
)

func TestSearchModel_DefaultNotVisible(t *testing.T) {
	m := NewSearchModel([]string{"* * * * *"})
	if m.Visible() {
		t.Error("expected search model to be hidden by default")
	}
}

func TestSearchModel_Toggle_ShowsAndHides(t *testing.T) {
	m := NewSearchModel([]string{"* * * * *"})
	m.Toggle()
	if !m.Visible() {
		t.Error("expected search model to be visible after toggle")
	}
	m.Toggle()
	if m.Visible() {
		t.Error("expected search model to be hidden after second toggle")
	}
}

func TestSearchModel_View_HiddenReturnsEmpty(t *testing.T) {
	m := NewSearchModel([]string{"* * * * *"})
	if m.View() != "" {
		t.Error("expected empty view when hidden")
	}
}

func TestSearchModel_View_VisibleContainsInput(t *testing.T) {
	m := NewSearchModel([]string{"* * * * *"})
	m.Toggle()
	view := m.View()
	if !strings.Contains(view, "Search") {
		t.Error("expected view to contain 'Search'")
	}
}

func TestSearchModel_Update_FiltersResults(t *testing.T) {
	pool := []string{"0 9 * * 1", "* * * * *", "30 6 * * *"}
	m := NewSearchModel(pool)
	m.Toggle()

	// Simulate typing "0 9"
	for _, ch := range "0 9" {
		m, _ = m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{ch}})
	}
	view := m.View()
	if !strings.Contains(view, "0 9 * * 1") {
		t.Error("expected view to contain matching expression")
	}
}

func TestSearchModel_Toggle_ClearsResultsOnHide(t *testing.T) {
	pool := []string{"* * * * *"}
	m := NewSearchModel(pool)
	m.Toggle()
	m, _ = m.Update(tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune{'*'}})
	m.Toggle() // hide
	if len(m.results) != 0 {
		t.Error("expected results cleared after hiding")
	}
}
