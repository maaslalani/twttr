package style

import (
	"github.com/charmbracelet/lipgloss"
)

const (
	hotPink   = lipgloss.Color("#FF00FF")
	lightGrey = lipgloss.Color("#5F5F5F")
	darkGrey  = lipgloss.Color("#444444")
)

var (
	// Tweet is the style for a tweet.
	Tweet = lipgloss.NewStyle().
		Border(lipgloss.ThickBorder(), false).
		BorderLeft(true).
		BorderForeground(hotPink).
		Width(80).
		Margin(0, 5).
		Padding(0, 3)

	// LoadingTweet is the style for a tweet that is loading.
	LoadingTweet = Tweet.Copy().BorderForeground(darkGrey)

	// AuthorName is the style for the author of a tweet.
	AuthorName = lipgloss.NewStyle().Bold(true)

	// AuthorHandle is the style for the handle of a tweet author.
	AuthorHandle = lipgloss.NewStyle().Foreground(darkGrey).MarginLeft(1)
)
