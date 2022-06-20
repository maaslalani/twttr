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
		BorderForeground(darkGrey).
		Width(80).
		MarginLeft(1).
		MarginBottom(1).
		Padding(0, 1)

	// SelectedTweet is the style for a selected tweet.
	SelectedTweet = Tweet.Copy().BorderForeground(hotPink)

	// Author is the style for the author of a tweet.
	Author = lipgloss.NewStyle().
		Foreground(lightGrey)

	// SelectedAuthor is the style for the author of the selected tweet.
	SelectedAuthor = lipgloss.NewStyle().
			Bold(true)
)
