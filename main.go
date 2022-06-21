package main

import (
	"log"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

// KeyMap is a map of key bindings.
// It defines the actions that a user can take.
type KeyMap struct {
	Quit, Next, Previous, Reload key.Binding
}

// DefaultKeyMap is the default key map that controls navigation and user
// actions for the application.
var DefaultKeyMap = KeyMap{
	Quit: key.NewBinding(
		key.WithKeys("q", "esc", "ctrl+c"),
		key.WithHelp("q", "quit"),
	),
	Next: key.NewBinding(
		key.WithKeys("right", "down", "l", "j", "n"),
		key.WithHelp("n", "next tweet"),
	),
	Previous: key.NewBinding(
		key.WithKeys("left", "up", "h", "k", "p"),
		key.WithHelp("p", "previous tweet"),
	),
	Reload: key.NewBinding(
		key.WithKeys("r"),
		key.WithHelp("r", "refresh"),
	),
}

func main() {
	m := model{keymap: DefaultKeyMap}
	program := tea.NewProgram(m, tea.WithAltScreen())

	err := program.Start()
	if err != nil {
		log.Fatal(err)
	}
}
