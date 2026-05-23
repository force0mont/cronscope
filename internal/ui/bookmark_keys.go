package ui

import "github.com/charmbracelet/bubbles/key"

// BookmarkKeyMap defines keybindings for the bookmark panel.
type BookmarkKeyMap struct {
	Toggle key.Binding
	Save   key.Binding
	Delete key.Binding
	Select key.Binding
	Up     key.Binding
	Down   key.Binding
}

// DefaultBookmarkKeys returns the default bookmark keybindings.
func DefaultBookmarkKeys() BookmarkKeyMap {
	return BookmarkKeyMap{
		Toggle: key.NewBinding(
			key.WithKeys("ctrl+b"),
			key.WithHelp("ctrl+b", "toggle bookmarks"),
		),
		Save: key.NewBinding(
			key.WithKeys("ctrl+s"),
			key.WithHelp("ctrl+s", "save bookmark"),
		),
		Delete: key.NewBinding(
			key.WithKeys("ctrl+d"),
			key.WithHelp("ctrl+d", "delete bookmark"),
		),
		Select: key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("enter", "load bookmark"),
		),
		Up: key.NewBinding(
			key.WithKeys("up", "k"),
			key.WithHelp("↑/k", "move up"),
		),
		Down: key.NewBinding(
			key.WithKeys("down", "j"),
			key.WithHelp("↓/j", "move down"),
		),
	}
}

// ShortHelp returns a compact keybinding summary for the bookmark panel.
func (k BookmarkKeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Toggle, k.Save, k.Delete, k.Select}
}

// FullHelp returns all keybindings for the bookmark panel.
func (k BookmarkKeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Toggle, k.Save},
		{k.Delete, k.Select},
		{k.Up, k.Down},
	}
}
