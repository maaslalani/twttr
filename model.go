package main

import (
	"fmt"
	"os/exec"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/maaslalani/twttr/keymap"
	"github.com/maaslalani/twttr/twitter"
)

type model struct {
	height        int
	keymap        keymap.Keymap
	selectedIndex int
	timeline      twitter.Timeline
	user          twitter.User
	view          View
	width         int
	textarea      textarea.Model
	err           error
}

type initialMsg struct{}

// Init initializes the model by fetching the current user and the current
// user's home timeline.
func (m model) Init() tea.Cmd {
	return func() tea.Msg {
		return initialMsg{}
	}
}

type fetchMsg struct {
	timeline twitter.Timeline
	user     twitter.User
}

type errorMsg struct {
	err error
}

func fetchTimeline() tea.Msg {
	me, err := twitter.Me()
	if err != nil {
		return errorMsg{err: err}
	}
	tl, err := twitter.HomeTimeline(me.ID)
	if err != nil {
		return errorMsg{err: err}
	}
	return fetchMsg{tl, me}
}

type sentTweetMsg struct{}

func (m model) sendTweet() tea.Msg {
	text := m.textarea.Value()
	err := twitter.CreateTweet(text)
	if err != nil {
		return errorMsg{err: err}
	}
	return sentTweetMsg{}
}

// Update the model
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keymap.Quit):
			if m.view != HomeView {
				m.view = HomeView
				m.textarea.Reset()
				m.textarea.Blur()
				m.keymap = keymap.Default
				break
			}
			return m, tea.Quit
		case key.Matches(msg, m.keymap.Next):
			m.view = HomeView
			if m.selectedIndex < len(m.timeline.Tweets)-1 {
				m.selectedIndex++
			}
		case key.Matches(msg, m.keymap.Previous):
			m.view = HomeView
			if m.selectedIndex > 0 {
				m.selectedIndex--
			}
		case key.Matches(msg, m.keymap.Open):
			tweet := m.timeline.Tweets[m.selectedIndex]
			author := getAuthor(m.timeline.Includes.Users, tweet.AuthorID)
			url := fmt.Sprintf("https://twitter.com/%s/status/%s", author.Username, tweet.ID)
			exec.Command("open", url).Run()
		case key.Matches(msg, m.keymap.Reload):
			m.selectedIndex = 0
			m.view = LoadingView
			return m, m.Init()
		case key.Matches(msg, m.keymap.Compose):
			m.keymap = keymap.Composing
			m.view = ComposeView
			m.textarea.Reset()
			m.textarea.Focus()
			// We return here to avoid the text area capturing the key binding
			// that led to the compose view.
			var cmd tea.Cmd
			// Note that we don't forward the message (i.e. the "c" key press).
			// Rather, send a blink to the cursor.
			m.textarea, cmd = m.textarea.Update(textarea.Blink())
			return m, cmd
		case key.Matches(msg, m.keymap.Tweet):
			m.view = TweetingView
			return m, m.sendTweet
		case key.Matches(msg, m.keymap.Help):
			if m.view == HelpView {
				m.view = HomeView
			} else {
				m.view = HelpView
			}
		}
	case tea.WindowSizeMsg:
		m.height = msg.Height
		m.width = min(msg.Width, maxTweetWidth)
	case initialMsg:
		m.keymap = keymap.Loading
		m.view = LoadingView
		return m, fetchTimeline
	case errorMsg:
		m.keymap = keymap.Error
		m.err = msg.err
		m.view = ErrorView
	case fetchMsg:
		m.timeline = msg.timeline
		m.user = msg.user
		if m.view == LoadingView {
			m.view = HomeView
		}

		m.keymap = keymap.Default
	case sentTweetMsg:
		m.view = HomeView
		m.keymap = keymap.Default
		return m, fetchTimeline
	}

	var cmd tea.Cmd
	m.textarea, cmd = m.textarea.Update(msg)
	return m, cmd
}

func getAuthor(users []twitter.User, id string) twitter.User {
	for _, user := range users {
		if user.ID == id {
			return user
		}
	}
	return twitter.User{}
}
