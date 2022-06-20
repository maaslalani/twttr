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
		MarginBottom(1).
		Padding(0, 1)

	// SelectedTweet is the style for a selected tweet.
	SelectedTweet = Tweet.Copy().BorderForeground(hotPink)

	// AuthorName is the style for the author of a tweet.
	AuthorName = lipgloss.NewStyle().Bold(true)

	// AuthorHandle is the style for the handle of a tweet author.
	AuthorHandle = lipgloss.NewStyle().Foreground(darkGrey).MarginLeft(1)

	// SelectedAuthorHandle is the style for the handle of an author of the
	// current selected tweet.
	SelectedAuthorHandle = AuthorHandle.Copy().Foreground(darkGrey)

	// SelectedAuthorName is the style for the author of the selected tweet.
	SelectedAuthorName = lipgloss.NewStyle().
				Bold(true)

	// Tab is the style for a navigation tab.
	Tab = lipgloss.NewStyle().
		Padding(1).
		Align(lipgloss.Right).
		Width(20)
)
