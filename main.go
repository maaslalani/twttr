package main

import (
	"log"

	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/maaslalani/twttr/keymap"
	"github.com/maaslalani/twttr/style"
)

func main() {
	textarea := textarea.New()
	textarea.CharLimit = 280
	textarea.SetHeight(3)
	textarea.SetWidth(50)
	textarea.Placeholder = "What's happening?"
	textarea.FocusedStyle.Prompt = style.Prompt
	textarea.FocusedStyle.CursorLine = lipgloss.NewStyle()
	textarea.ShowLineNumbers = false

	m := model{keymap: keymap.Default, view: LoadingView, textarea: textarea}
	err := tea.NewProgram(m).Start()
	if err != nil {
		log.Fatal(err)
	}
}
