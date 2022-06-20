package main

import (
	"log"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	view := viewport.New(80, 0)
	view.KeyMap = viewport.KeyMap{}

	model := model{viewport: view}
	program := tea.NewProgram(model, tea.WithAltScreen())

	err := program.Start()
	if err != nil {
		log.Fatal(err)
	}
}
