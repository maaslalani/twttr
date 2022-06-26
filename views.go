package main

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/maaslalani/twttr/style"
	"github.com/maaslalani/twttr/twitter"
	"github.com/muesli/reflow/wordwrap"
)

// View is a type that represents a view in the application.
// See below for the different views a user can navigate to.
type View int

const (
	// LoadingView is the view when the model is fetching data from the
	// server, in this case the Twitter API.
	LoadingView View = iota
	// TweetingView is the view when the user has sent a tweet while it is still
	// loading.
	TweetingView
	// HelpView is the view to show when the user requests to view all of the
	// keybindings.
	HelpView
	// ComposeView is the view to show when the user is composing a new tweet.
	ComposeView
	// HomeView is the view to show when the user is viewing their home
	// timeline.
	HomeView
)

func (m model) loadingView() string {
	loadingAuthor := style.AuthorName.Render("Loading") + style.AuthorHandle.Render("@loading")
	loadingTweet := style.LoadingTweet.Render(loadingAuthor + "\n" + "This shouldn't take too long...")
	helpText := style.Help.Render("\n" + "Press ? for help")
	return loadingTweet + helpText
}

func (m model) tweetingView() string {
	author := style.AuthorName.Render(m.user.Name) + style.AuthorHandle.Render("@"+m.user.Username)
	sentTweet := style.LoadingTweet.Render(author + "\n" + m.textarea.Value())
	return sentTweet
}

func (m model) tweetsView() string {
	// tweet := m.timeline.Tweets[m.selectedIndex]
	// author := getAuthor(m.timeline.Includes.Users, tweet.AuthorID)
	authorNameStyled := style.AuthorName.Render("jack")
	authorHandleStyled := style.AuthorHandle.Render("@" + "jack")
	styledTweet := style.Tweet.Render(authorNameStyled + authorHandleStyled + "\n" + wordwrap.String("just setting up my twttr", m.width))
	return styledTweet
}

func (m model) composeView() string {
	tweet := m.textarea.View()
	return "\n" + tweet
}

func (m model) helpView() string {
	navigationHelp := style.Help.Render(`
?  Toggle Help
r  Reload Timeline
c  Compose Tweet`)

	tweetHelp := style.Help.Render(`
n  Next Tweet
p  Previous Tweet
o  Open Tweet in Browser`)

	return lipgloss.JoinHorizontal(lipgloss.Top, navigationHelp, tweetHelp)
}

func (m model) View() string {
	switch m.view {
	case HomeView:
		return m.tweetsView()
	case ComposeView:
		return m.composeView()
	case HelpView:
		return m.helpView()
	case LoadingView:
		return m.loadingView()
	case TweetingView:
		return m.tweetingView()
	}

	return m.loadingView()
}

func tweetHeight(tweet twitter.Tweet) int {
	tweetHeight := lipgloss.Height(style.Tweet.Render(tweet.Text))
	authorHeight := 1
	return tweetHeight + authorHeight
}

const maxTweetWidth = 80

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
