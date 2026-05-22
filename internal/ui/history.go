package ui

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	historyTitleStyle  = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("63"))
	historyItemStyle   = lipgloss.NewStyle().PaddingLeft(2)
	historySelectStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("212")).Bold(true)
)

// HistorySelectedMsg is emitted when the user picks a history entry.
type HistorySelectedMsg struct{ Expression string }

// HistoryModel is a scrollable list of recent cron expressions.
type HistoryModel struct {
	items    []string
	cursor   int
	visible  bool
}

// NewHistoryModel creates a HistoryModel from a slice of expressions.
func NewHistoryModel(items []string) HistoryModel {
	return HistoryModel{items: items}
}

// SetItems replaces the history list.
func (m *HistoryModel) SetItems(items []string) { m.items = items }

// Toggle shows or hides the history panel.
func (m *HistoryModel) Toggle() { m.visible = !m.visible }

// IsVisible reports whether the panel is currently shown.
func (m HistoryModel) IsVisible() bool { return m.visible }

// Update handles keyboard navigation within the history list.
func (m HistoryModel) Update(msg tea.Msg) (HistoryModel, tea.Cmd) {
	if !m.visible {
		return m, nil
	}
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.items)-1 {
				m.cursor++
			}
		case "enter":
			if len(m.items) > 0 {
				return m, func() tea.Msg {
					return HistorySelectedMsg{Expression: m.items[m.cursor]}
				}
			}
		}
	}
	return m, nil
}

// View renders the history panel.
func (m HistoryModel) View() string {
	if !m.visible {
		return ""
	}
	var b strings.Builder
	b.WriteString(historyTitleStyle.Render("Recent expressions") + "\n")
	if len(m.items) == 0 {
		b.WriteString(historyItemStyle.Render("(no history yet)") + "\n")
		return b.String()
	}
	for i, expr := range m.items {
		line := fmt.Sprintf("%2d. %s", i+1, expr)
		if i == m.cursor {
			b.WriteString(historySelectStyle.Render("> "+line) + "\n")
		} else {
			b.WriteString(historyItemStyle.Render("  "+line) + "\n")
		}
	}
	b.WriteString("\n" + lipgloss.NewStyle().Faint(true).Render("↑/↓ navigate · enter select · h hide"))
	return b.String()
}
