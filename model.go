package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/maaslalani/twttr/style"
	"github.com/maaslalani/twttr/twitter"
)

type model struct {
	timeline      twitter.Timeline
	user          twitter.User
	selectedIndex int
	height        int
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
		case "down", "j", "n":
			if m.selectedIndex < len(m.timeline.Tweets)-1 {
				m.selectedIndex++
			}
		case "up", "k", "p":
			if m.selectedIndex > 0 {
				m.selectedIndex--
			}
		}
	case tea.WindowSizeMsg:
		m.height = msg.Height
	case fetchMsg:
		m.timeline = msg.timeline
		m.user = msg.user
	}

	return m, nil
}

func (m model) navigationView() string {
	return lipgloss.PlaceHorizontal(15, lipgloss.Center, "Navigation")
}

func (m model) loadingTweetsView() string {
	loadingAuthor := style.AuthorName.Render("Loading") + style.AuthorHandle.Render("@loading")
	loadingTweet := style.LoadingTweet.Render(loadingAuthor+"\n"+"This shouldn't take too long, only a few seconds....") + "\n"
	return lipgloss.PlaceVertical(m.height, lipgloss.Center, loadingTweet)
}

func (m model) tweetsView() string {
	if len(m.timeline.Tweets) == 0 {
		return m.loadingTweetsView()
	}

	tweet := m.timeline.Tweets[m.selectedIndex]
	author := getAuthor(m.timeline.Includes.Users, tweet.AuthorID)
	authorNameStyled := style.AuthorName.Render(author.Name)
	authorHandleStyled := style.AuthorHandle.Render("@" + author.Username)
	styledTweet := style.Tweet.Render(authorNameStyled + authorHandleStyled + "\n" + tweet.Text)
	return lipgloss.PlaceVertical(m.height, lipgloss.Center, styledTweet)
}

func (m model) View() string {
	return lipgloss.JoinHorizontal(
		lipgloss.Center,
		m.navigationView(),
		m.tweetsView(),
	)
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
