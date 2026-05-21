package ui

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/cronscope/cronscope/internal/cron"
)

// keyMap defines keybindings for the application.
type keyMap struct {
	Submit key.Binding
	Quit   key.Binding
}

var keys = keyMap{
	Submit: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "preview schedule"),
	),
	Quit: key.NewBinding(
		key.WithKeys("ctrl+c", "esc"),
		key.WithHelp("ctrl+c", "quit"),
	),
}

// AppModel is the root Bubble Tea model for cronscope.
type AppModel struct {
	input    InputModel
	schedule *ScheduleView
}

// NewAppModel creates the root application model.
func NewAppModel() AppModel {
	return AppModel{
		input: NewInputModel(),
	}
}

// Init implements tea.Model.
func (m AppModel) Init() tea.Cmd {
	return m.input.Init()
}

// Update implements tea.Model.
func (m AppModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if key.Matches(msg, keys.Quit) {
			return m, tea.Quit
		}
		if key.Matches(msg, keys.Submit) {
			expr := m.input.Value()
			sv, err := NewScheduleView(expr, cron.DefaultPreviewCount)
			if err != nil {
				m.input.SetError(err)
				m.schedule = nil
			} else {
				m.input.SetError(nil)
				m.schedule = sv
			}
			return m, nil
		}
	}

	var cmd tea.Cmd
	m.input, cmd = m.input.Update(msg)
	return m, cmd
}

// View implements tea.Model.
func (m AppModel) View() string {
	out := m.input.View() + "\n\n"
	if m.schedule != nil {
		out += m.schedule.Render()
	}
	out += "\n" + keys.Quit.Help().Desc + " · " + keys.Submit.Help().Desc
	return out
}
