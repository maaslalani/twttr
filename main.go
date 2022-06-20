package main

// A simple program that makes a GET request and prints the response status.

import (
	"log"
	"strings"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/maaslalani/twttr/style"
	"github.com/maaslalani/twttr/twitter"
)

type model struct {
	viewport viewport.Model
	timeline twitter.Timeline
	user     twitter.User
}

type fetchMsg struct {
	timeline twitter.Timeline
	user     twitter.User
}

func main() {
	m := model{viewport: viewport.New(80, 24)}
	p := tea.NewProgram(m, tea.WithAltScreen())
	err := p.Start()
	if err != nil {
		log.Fatal(err)
	}
}

func (m model) Init() tea.Cmd {
	return fetchTimeline
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c", "esc":
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
	case fetchMsg:
		m.timeline = msg.timeline
		m.user = msg.user
		m.viewport.SetContent(m.tweetsView())
	}

	var cmd tea.Cmd
	m.viewport, cmd = m.viewport.Update(msg)
	return m, cmd
}

func (m model) tweetsView() string {
	var s strings.Builder

	for _, tweet := range m.timeline.Tweets {
		authorName := getUserName(m.timeline.Includes.Users, tweet.AuthorID)
		s.WriteString(style.Tweet.Render(style.Author.Render(authorName) + "\n" + tweet.Text))
		s.WriteRune('\n')
	}

	return s.String()
}

func (m model) View() string {
	return m.viewport.View()
}

func fetchTimeline() tea.Msg {
	user := twitter.Me()
	timeline := twitter.HomeTimeline(user.ID)
	return fetchMsg{timeline, user}
}

func getUserName(users []twitter.User, id string) string {
	for _, user := range users {
		if user.ID == id {
			return user.Name
		}
	}
	return ""
}
