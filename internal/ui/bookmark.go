package ui

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"

	"cronscope/internal/cron"
)

var (
	bookmarkTitleStyle  = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("33"))
	bookmarkItemStyle   = lipgloss.NewStyle().PaddingLeft(2)
	bookmarkCursorStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("205")).Bold(true)
)

// BookmarkModel is a Bubble Tea component for browsing saved bookmarks.
type BookmarkModel struct {
	store   *cron.BookmarkStore
	cursor  int
	visible bool
}

// NewBookmarkModel creates a BookmarkModel backed by the given store.
func NewBookmarkModel(store *cron.BookmarkStore) BookmarkModel {
	return BookmarkModel{store: store}
}

func (m BookmarkModel) Init() tea.Cmd { return nil }

func (m BookmarkModel) Update(msg tea.Msg) (BookmarkModel, tea.Cmd) {
	if !m.visible {
		return m, nil
	}
	switch msg := msg.(type) {
	case tea.KeyMsg:
		items := m.store.All()
		switch msg.String() {
		case "j", "down":
			if m.cursor < len(items)-1 {
				m.cursor++
			}
		case "k", "up":
			if m.cursor > 0 {
				m.cursor--
			}
		}
	}
	return m, nil
}

// Toggle shows or hides the bookmark panel.
func (m *BookmarkModel) Toggle() {
	m.visible = !m.visible
	m.cursor = 0
}

// IsVisible reports whether the bookmark panel is currently shown.
func (m BookmarkModel) IsVisible() bool {
	return m.visible
}

// Selected returns the currently highlighted bookmark, or nil.
func (m BookmarkModel) Selected() *cron.Bookmark {
	items := m.store.All()
	if len(items) == 0 || m.cursor >= len(items) {
		return nil
	}
	b := items[m.cursor]
	return &b
}

func (m BookmarkModel) View() string {
	if !m.visible {
		return ""
	}
	items := m.store.All()
	var sb strings.Builder
	sb.WriteString(bookmarkTitleStyle.Render("Bookmarks") + "\n")
	if len(items) == 0 {
		sb.WriteString(bookmarkItemStyle.Render("(no bookmarks saved)") + "\n")
		return sb.String()
	}
	for i, b := range items {
		cursor := "  "
		if i == m.cursor {
			cursor = bookmarkCursorStyle.Render("> ")
		}
		line := fmt.Sprintf("%s%s  %s", cursor, b.Label, b.Expr)
		if b.Timezone != "" {
			line += fmt.Sprintf(" (%s)", b.Timezone)
		}
		sb.WriteString(bookmarkItemStyle.Render(line) + "\n")
	}
	return sb.String()
}
