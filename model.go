package main

import (
	"strings"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/maaslalani/twttr/style"
	"github.com/maaslalani/twttr/twitter"
)

type model struct {
	viewport      viewport.Model
	timeline      twitter.Timeline
	user          twitter.User
	selectedIndex int
}

type fetchMsg struct {
	timeline twitter.Timeline
	user     twitter.User
}

// Init initializes the model by fetching the current user and the current
// user's home timeline.
func (m model) Init() tea.Cmd {
	return func() tea.Msg {
		me := twitter.Me()
		tl := twitter.HomeTimeline(me.ID)
		return fetchMsg{tl, me}
	}
}

// Update the model
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c", "esc":
			return m, tea.Quit
		case "j":
			if m.selectedIndex < len(m.timeline.Tweets)-1 {
				m.selectedIndex++
			}
		case "k":
			if m.selectedIndex > 0 {
				m.selectedIndex--
			}
		}
	case tea.WindowSizeMsg:
		m.viewport.Height = msg.Height
	case fetchMsg:
		m.timeline = msg.timeline
		m.user = msg.user
	}

	m.viewport.SetContent(m.tweetsView())

	var cmd tea.Cmd
	m.viewport, cmd = m.viewport.Update(msg)
	return m, cmd
}

func (m model) tweetsView() string {
	var s strings.Builder

	for i, tweet := range m.timeline.Tweets {
		var tweetStyle, authorStyle lipgloss.Style
		if m.selectedIndex == i {
			tweetStyle = style.SelectedTweet
			authorStyle = style.SelectedAuthor
		} else {
			tweetStyle = style.Tweet
			authorStyle = style.Author
		}
		authorName := getUserName(m.timeline.Includes.Users, tweet.AuthorID)
		authorStyled := authorStyle.Render(authorName)
		s.WriteString(tweetStyle.Render(authorStyled + "\n" + tweet.Text))
		s.WriteRune('\n')
	}

	return s.String()
}

func (m model) View() string {
	return m.viewport.View()
}

func getUserName(users []twitter.User, id string) string {
	for _, user := range users {
		if user.ID == id {
			return user.Name
		}
	}
	return ""
}
