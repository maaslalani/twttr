package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	m := model{}
	program := tea.NewProgram(m, tea.WithAltScreen())

	err := program.Start()
	if err != nil {
		log.Fatal(err)
	}
}
