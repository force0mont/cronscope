package ui

import "github.com/charmbracelet/bubbles/key"

// SearchKeyMap defines the keybindings for the search panel.
type SearchKeyMap struct {
	Toggle key.Binding
	Close  key.Binding
}

// DefaultSearchKeys returns the default search keybindings.
func DefaultSearchKeys() SearchKeyMap {
	return SearchKeyMap{
		Toggle: key.NewBinding(
			key.WithKeys("/"),
			key.WithHelp("/", "search"),
		),
		Close: key.NewBinding(
			key.WithKeys("esc"),
			key.WithHelp("esc", "close search"),
		),
	}
}

// ShortHelp returns a compact list of key bindings for the search panel.
func (k SearchKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Toggle, k.Close}
}

// FullHelp returns the full list of key bindings for the search panel.
func (k SearchKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{{k.Toggle, k.Close}}
}
