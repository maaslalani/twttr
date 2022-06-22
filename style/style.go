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
		Margin(1, 1, 2, 1).
		Padding(0, 2)

	// LoadingTweet is the style for a tweet that is loading.
	LoadingTweet = Tweet.Copy().BorderForeground(darkGrey).Margin(1)

	// AuthorName is the style for the author of a tweet.
	AuthorName = lipgloss.NewStyle().Bold(true)

	// AuthorHandle is the style for the handle of a tweet author.
	AuthorHandle = lipgloss.NewStyle().Foreground(darkGrey).MarginLeft(1)

	// Help is the style for the help text.
	Help = lipgloss.NewStyle().Foreground(lightGrey).Margin(0, 2)

	// Compose is the style for the compose a tweet text.
	Compose = lipgloss.NewStyle().Border(lipgloss.ThickBorder(), true).Margin(0, 1)

	// Prompt is the style for the compose tweet multi-line prompt.
	Prompt = lipgloss.NewStyle().Foreground(darkGrey).Margin(0, 1)
)
