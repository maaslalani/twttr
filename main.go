package main

import (
	"log"

	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/maaslalani/twttr/style"
	"github.com/maaslalani/twttr/views"
)

func main() {
	textarea := textarea.New()
	textarea.CharLimit = 280
	textarea.Height = 3
	textarea.Width = 50
	textarea.Placeholder = "What's happening?"
	textarea.Prompt = lipgloss.ThickBorder().Left
	textarea.PromptStyle = style.Prompt

	m := model{keymap: DefaultKeyMap, view: views.Loading, textarea: textarea}
	err := tea.NewProgram(m).Start()
	if err != nil {
		log.Fatal(err)
	}
}
