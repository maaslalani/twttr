package main

import (
	"fmt"
	"os/exec"

	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/maaslalani/twttr/style"
	"github.com/maaslalani/twttr/twitter"
	"github.com/maaslalani/twttr/views"
)

type model struct {
	height        int
	keymap        KeyMap
	selectedIndex int
	timeline      twitter.Timeline
	user          twitter.User
	view          views.View
	width         int
}

type initialMsg struct{}

// Init initializes the model by fetching the current user and the current
// user's home timeline.
func (m model) Init() tea.Cmd {
	return func() tea.Msg {
		return initialMsg{}
	}
}

func fetchTimeline() tea.Msg {
	me := twitter.Me()
	tl := twitter.HomeTimeline(me.ID)
	return fetchMsg{tl, me}
}

type fetchMsg struct {
	timeline twitter.Timeline
	user     twitter.User
}

// Update the model
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keymap.Quit):
			return m, tea.Quit
		case key.Matches(msg, m.keymap.Next):
			m.view = views.Home
			if m.selectedIndex < len(m.timeline.Tweets)-1 {
				m.selectedIndex++
			}
		case key.Matches(msg, m.keymap.Previous):
			m.view = views.Home
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
			m.view = views.Loading
			return m, m.Init()
		case key.Matches(msg, m.keymap.Compose):
			m.view = views.Compose
		case key.Matches(msg, m.keymap.Help):
			if m.view == views.Help {
				m.view = views.Home
			} else {
				m.view = views.Help
			}
		}
	case tea.WindowSizeMsg:
		m.height = msg.Height
		m.width = msg.Width
	case initialMsg:
		m.keymap.Next.SetEnabled(false)
		m.keymap.Previous.SetEnabled(false)
		m.view = views.Loading
		return m, fetchTimeline
	case fetchMsg:
		m.timeline = msg.timeline
		m.user = msg.user
		m.view = views.Home

		m.keymap.Next.SetEnabled(true)
		m.keymap.Previous.SetEnabled(true)
	}

	return m, nil
}

func (m model) loadingView() string {
	loadingAuthor := style.AuthorName.Render("Loading") + style.AuthorHandle.Render("@loading")
	loadingTweet := style.LoadingTweet.Render(loadingAuthor + "\n" + "This shouldn't take too long...")
	helpText := style.Help.Render("\n" + "Press ? for help")
	return lipgloss.PlaceVertical(m.height, lipgloss.Center, loadingTweet+helpText)
}

func (m model) tweetsView() string {
	tweet := m.timeline.Tweets[m.selectedIndex]
	author := getAuthor(m.timeline.Includes.Users, tweet.AuthorID)
	authorNameStyled := style.AuthorName.Render(author.Name)
	authorHandleStyled := style.AuthorHandle.Render("@" + author.Username)
	styledTweet := style.Tweet.Render(authorNameStyled + authorHandleStyled + "\n" + tweet.Text)

	return lipgloss.PlaceVertical(m.height, lipgloss.Center, styledTweet)
}

func (m model) composeView() string {
	return lipgloss.PlaceVertical(m.height, lipgloss.Center, "Compose a tweet")
}

func (m model) helpView() string {
	helpText := `
	?  Toggle Help
	q  Quit

	r  Reload Timeline
	c  Compose Tweet

	n  Next Tweet
	p  Previous Tweet

	L  Like Tweet
	R  Retweet Tweet
	o  Open in Browser
	`
	return lipgloss.PlaceVertical(m.height, lipgloss.Center, style.Help.Render(helpText))
}

func (m model) View() string {
	switch m.view {
	case views.Home:
		return m.tweetsView()
	case views.Compose:
		return m.composeView()
	case views.Help:
		return m.helpView()
	case views.Loading:
		return m.loadingView()
	}

	return m.loadingView()
}

func tweetHeight(tweet twitter.Tweet) int {
	tweetHeight := lipgloss.Height(style.Tweet.Render(tweet.Text))
	authorHeight := 1
	return tweetHeight + authorHeight
}

func getAuthor(users []twitter.User, id string) twitter.User {
	for _, user := range users {
		if user.ID == id {
			return user
		}
	}
	return twitter.User{}
}
