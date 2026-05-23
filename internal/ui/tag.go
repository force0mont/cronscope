package ui

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"cronscope/internal/cron"
)

var (
	tagTitleStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("205"))
	tagLabelStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("39"))
	tagExprStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("250"))
)

// TagModel manages the tag panel UI.
type TagModel struct {
	store    *cron.TagStore
	visible  bool
	selected int
	labels   []string
}

// NewTagModel creates a TagModel backed by the given store.
func NewTagModel(store *cron.TagStore) *TagModel {
	return &TagModel{store: store}
}

// Toggle shows or hides the tag panel.
func (m *TagModel) Toggle() {
	m.visible = !m.visible
	if m.visible {
		m.labels = m.store.Labels()
		m.selected = 0
	}
}

// Visible reports whether the panel is currently shown.
func (m *TagModel) Visible() bool { return m.visible }

// Update handles key messages for navigation.
func (m *TagModel) Update(msg tea.Msg) (*TagModel, tea.Cmd) {
	if !m.visible {
		return m, nil
	}
	if key, ok := msg.(tea.KeyMsg); ok {
		switch key.String() {
		case "up", "k":
			if m.selected > 0 {
				m.selected--
			}
		case "down", "j":
			if m.selected < len(m.labels)-1 {
				m.selected++
			}
		case "esc":
			m.visible = false
		}
	}
	return m, nil
}

// SelectedLabel returns the currently highlighted label, or empty string.
func (m *TagModel) SelectedLabel() string {
	if !m.visible || len(m.labels) == 0 {
		return ""
	}
	return m.labels[m.selected]
}

// View renders the tag panel.
func (m *TagModel) View() string {
	if !m.visible {
		return ""
	}
	var sb strings.Builder
	sb.WriteString(tagTitleStyle.Render("Tags") + "\n")
	if len(m.labels) == 0 {
		sb.WriteString("  (no tags)\n")
		return sb.String()
	}
	for i, label := range m.labels {
		cursor := "  "
		if i == m.selected {
			cursor = "> "
		}
		exprs := m.store.ByLabel(label)
		sb.WriteString(fmt.Sprintf("%s%s (%d)\n", cursor, tagLabelStyle.Render(label), len(exprs)))
		if i == m.selected {
			for _, e := range exprs {
				sb.WriteString("    " + tagExprStyle.Render(e) + "\n")
			}
		}
	}
	return sb.String()
}
