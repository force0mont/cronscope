package ui

import (
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

// InputModel handles the cron expression text input field.
type InputModel struct {
	textInput textinput.Model
	Err       error
}

// NewInputModel creates and initializes a new InputModel.
func NewInputModel() InputModel {
	ti := textinput.New()
	ti.Placeholder = "e.g. 0 9 * * 1-5"
	ti.Focus()
	ti.CharLimit = 100
	ti.Width = 40

	return InputModel{
		textInput: ti,
	}
}

// Init implements tea.Model.
func (m InputModel) Init() tea.Cmd {
	return textinput.Blink
}

// Update implements tea.Model.
func (m InputModel) Update(msg tea.Msg) (InputModel, tea.Cmd) {
	var cmd tea.Cmd
	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

// View implements tea.Model.
func (m InputModel) View() string {
	var sb strings.Builder
	sb.WriteString("Enter cron expression:\n")
	sb.WriteString(m.textInput.View())
	if m.Err != nil {
		sb.WriteString("\n⚠  " + m.Err.Error())
	}
	return sb.String()
}

// Value returns the current text input value.
func (m InputModel) Value() string {
	return m.textInput.Value()
}

// SetError sets a validation error to display beneath the input.
func (m *InputModel) SetError(err error) {
	m.Err = err
}

// Reset clears the input field and any error.
func (m *InputModel) Reset() {
	m.textInput.Reset()
	m.Err = nil
}
