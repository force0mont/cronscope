package ui

import "github.com/charmbracelet/bubbles/key"

// TagKeyMap defines keybindings for the tag panel.
type TagKeyMap struct {
	Toggle key.Binding
	Up     key.Binding
	Down   key.Binding
	Close  key.Binding
}

// DefaultTagKeys returns the default key bindings for the tag panel.
func DefaultTagKeys() TagKeyMap {
	return TagKeyMap{
		Toggle: key.NewBinding(
			key.WithKeys("t"),
			key.WithHelp("t", "toggle tags"),
		),
		Up: key.NewBinding(
			key.WithKeys("up", "k"),
			key.WithHelp("↑/k", "move up"),
		),
		Down: key.NewBinding(
			key.WithKeys("down", "j"),
			key.WithHelp("↓/j", "move down"),
		),
		Close: key.NewBinding(
			key.WithKeys("esc"),
			key.WithHelp("esc", "close tags"),
		),
	}
}

// ShortHelp returns a compact help listing.
func (k TagKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Toggle, k.Up, k.Down, k.Close}
}

// FullHelp returns the complete help listing.
func (k TagKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{{k.Toggle, k.Up, k.Down, k.Close}}
}
