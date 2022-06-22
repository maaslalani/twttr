package main

import "github.com/charmbracelet/bubbles/key"

// KeyMap is a map of key bindings.
// It defines the actions that a user can take.
type KeyMap struct {
	Compose  key.Binding
	Help     key.Binding
	Like     key.Binding
	Next     key.Binding
	Open     key.Binding
	Previous key.Binding
	Quit     key.Binding
	Reload   key.Binding
	Retweet  key.Binding
}

// DefaultKeyMap is the default key map that controls navigation and user
// actions for the application.
var DefaultKeyMap = KeyMap{
	Compose: key.NewBinding(
		key.WithKeys("c"),
		key.WithHelp("c", "compose tweet"),
	),
	Help: key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp("?", "help"),
	),
	Like: key.NewBinding(
		key.WithKeys("L"),
		key.WithHelp("L", "like"),
	),
	Next: key.NewBinding(
		key.WithKeys("right", "down", "l", "j", "n"),
		key.WithHelp("n", "next tweet"),
	),
	Open: key.NewBinding(
		key.WithKeys("o"),
		key.WithHelp("o", "open in browser"),
	),
	Previous: key.NewBinding(
		key.WithKeys("left", "up", "h", "k", "p"),
		key.WithHelp("p", "previous tweet"),
	),
	Quit: key.NewBinding(
		key.WithKeys("q", "esc", "ctrl+c"),
		key.WithHelp("q", "quit"),
	),
	Reload: key.NewBinding(
		key.WithKeys("r"),
		key.WithHelp("r", "refresh"),
	),
	Retweet: key.NewBinding(
		key.WithKeys("R"),
		key.WithHelp("R", "retweet"),
	),
}
