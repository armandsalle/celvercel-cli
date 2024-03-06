package styles

import "github.com/charmbracelet/lipgloss"

var Wrapper = lipgloss.NewStyle().
	Border(lipgloss.RoundedBorder()).
	BorderForeground(lipgloss.Color("#7D56F4")).
	Width(22).
	Align(lipgloss.Center)

var OddWordStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#FAFAFA")).
	Background(lipgloss.Color("#04B575")).
	Padding(0, 1)

var EvenWordStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#3C3C3C")).
	Background(lipgloss.Color("#FAFAFA")).
	Padding(0, 1)

var ErrorStyle = OddWordStyle.Copy().
	Background(lipgloss.Color("#f87171"))
