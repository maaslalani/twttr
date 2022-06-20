package main

import (
	"fmt"
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
	scrollOffset  int
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
				m.scrollOffset += tweetHeight(m.timeline.Tweets[m.selectedIndex])
			}
		case "up", "k", "p":
			if m.selectedIndex > 0 {
				m.scrollOffset -= tweetHeight(m.timeline.Tweets[m.selectedIndex])
				m.selectedIndex--
			}
		}
	case tea.WindowSizeMsg:
		m.viewport.Height = msg.Height
	case fetchMsg:
		m.timeline = msg.timeline
		m.user = msg.user
	}

	min := m.viewport.YOffset
	max := min + m.viewport.Height
	if m.scrollOffset >= max {
		m.viewport.LineDown(tweetHeight(m.timeline.Tweets[m.selectedIndex]))
	} else if m.scrollOffset < min {
		m.viewport.LineUp(tweetHeight(m.timeline.Tweets[m.selectedIndex]))
	}
	m.viewport.SetContent(m.tweetsView())
	var cmd tea.Cmd
	m.viewport, cmd = m.viewport.Update(msg)
	return m, cmd
}

func (m model) navigationView() string {
	var s strings.Builder
	for _, tab := range []string{"Home", "Explore", "Notification", "Messages", "Bookmarks", "Lists", "Profile"} {
		s.WriteString(style.Tab.Render(tab))
	}
	return s.String()
}

func (m model) trendingView() string {
	return fmt.Sprintf("Y Offset %d, Tweet Offset %d", m.viewport.YOffset, m.scrollOffset)
}

func loadingTweetsView() string {
	loadingAuthor := style.AuthorName.Render("Loading") + style.AuthorHandle.Render("@loading")
	loadingTweet := style.Tweet.Render(loadingAuthor+"\n"+"This shouldn't take too long, only a few seconds....") + "\n"
	return strings.Repeat(loadingTweet, 10)
}

func (m model) tweetsView() string {
	if len(m.timeline.Tweets) == 0 {
		return loadingTweetsView()
	}

	var s strings.Builder

	for i, tweet := range m.timeline.Tweets {
		var tweetStyle, authorNameStyle, authorHandleStyle lipgloss.Style
		if m.selectedIndex == i {
			tweetStyle = style.SelectedTweet
			authorNameStyle = style.SelectedAuthorName
			authorHandleStyle = style.SelectedAuthorHandle
		} else {
			tweetStyle = style.Tweet
			authorNameStyle = style.AuthorName
			authorHandleStyle = style.AuthorHandle
		}
		author := getAuthor(m.timeline.Includes.Users, tweet.AuthorID)
		authorNameStyled := authorNameStyle.Render(author.Name)
		authorHandleStyled := authorHandleStyle.Render("@" + author.Username)
		s.WriteString(tweetStyle.Render(authorNameStyled + authorHandleStyled + "\n" + tweet.Text))
		s.WriteRune('\n')
	}

	return s.String()
}

func (m model) View() string {
	return lipgloss.JoinHorizontal(
		lipgloss.Center,
		m.navigationView(),
		m.viewport.View(),
		m.trendingView(),
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
