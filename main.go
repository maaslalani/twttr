package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/maaslalani/twttr/views"
)

func main() {
	m := model{keymap: DefaultKeyMap, view: views.Loading}
	err := tea.NewProgram(m, tea.WithAltScreen()).Start()
	if err != nil {
		log.Fatal(err)
	}
}
