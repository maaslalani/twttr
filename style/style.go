package style

import (
	"github.com/charmbracelet/lipgloss"
)

const hotPink = lipgloss.Color("#FF00FF")
const lightGrey = lipgloss.Color("#5F5F5F")
const darkGrey = lipgloss.Color("#444444")

var (
	// Tweet is the style for a tweet.
	Tweet = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder(), true).
		BorderForeground(darkGrey).
		Width(80).
		Margin(0, 5).
		Padding(0, 2)

	// ActiveTweet is the style for an active tweet.
	ActiveTweet = Tweet.Copy().BorderForeground(hotPink)

	// Author is the style for the author of a tweet.
	Author = lipgloss.NewStyle().
		Bold(true).
		Foreground(lightGrey)
)
