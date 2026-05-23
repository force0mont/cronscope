package ui

import (
	"fmt"
	"strings"
	"time"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/user/cronscope/internal/cron"
)

// SearchModel provides a live search UI over a set of cron expressions.
type SearchModel struct {
	input   textinput.Model
	pool    []string
	results []cron.SearchResult
	visible bool
}

// NewSearchModel creates a SearchModel backed by the given expression pool.
func NewSearchModel(pool []string) SearchModel {
	ti := textinput.New()
	ti.Placeholder = "Search expressions…"
	ti.CharLimit = 64
	return SearchModel{
		input: tti,
		pool:  pool,
	}
}

// Toggle shows or hides the search panel.
func (m *SearchModel) Toggle() {
	m.visible = !m.visible
	if m.visible {
		m.input.Focus()
	} else {
		m.input.Blur()
		m.input.Reset()
		m.results = nil
	}
}

// Visible returns whether the search panel is shown.
func (m SearchModel) Visible() bool { return m.visible }

// Update handles key messages when the search panel is active.
func (m SearchModel) Update(msg tea.Msg) (SearchModel, tea.Cmd) {
	if !m.visible {
		return m, nil
	}
	var cmd tea.Cmd
	m.input, cmd = m.input.Update(msg)
	m.results = cron.Search(m.pool, m.input.Value(), 8, time.Now())
	return m, cmd
}

// View renders the search panel.
func (m SearchModel) View() string {
	if !m.visible {
		return ""
	}
	var sb strings.Builder
	sb.WriteString("┌─ Search ──────────────────────────────────┐\n")
	sb.WriteString("  " + m.input.View() + "\n")
	if len(m.results) == 0 {
		sb.WriteString("  (no results)\n")
	} else {
		for _, r := range m.results {
			sb.WriteString(fmt.Sprintf("  %-24s  next: %s\n",
				r.Expression,
				r.NextRun.Format("2006-01-02 15:04:05"),
			))
		}
	}
	sb.WriteString("└───────────────────────────────────────────┘")
	return sb.String()
}
