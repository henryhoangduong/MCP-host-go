package cmd

import (
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
)

var (
	renderer *glamour.TermRenderer

	// Tokyo Night theme colors
	tokyoPurple = lipgloss.Color("99")  // #9d7cd8
	tokyoCyan   = lipgloss.Color("73")  // #7dcfff
	tokyoBlue   = lipgloss.Color("111") // #7aa2f7
	tokyoGreen  = lipgloss.Color("120") // #73daca
	tokyoRed    = lipgloss.Color("203") // #f7768e
	tokyoOrange = lipgloss.Color("215") // #ff9e64
	tokyoFg     = lipgloss.Color("189") // #c0caf5
	tokyoGray   = lipgloss.Color("237") // #3b4261
	tokyoBg     = lipgloss.Color("234") // #1a1b26

	serverCommandStyle = lipgloss.NewStyle().
				Foreground(tokyoOrange).
				Bold(true)

	serverArgumentsStyle = lipgloss.NewStyle().
				Foreground(tokyoFg)

	serverHeaderStyle = lipgloss.NewStyle().
				Foreground(tokyoCyan).
				Bold(true)

	promptStyle = lipgloss.NewStyle().
			Foreground(tokyoBlue).
			PaddingLeft(2)

	responseStyle = lipgloss.NewStyle().
			Foreground(tokyoFg).
			PaddingLeft(2)

	errorStyle = lipgloss.NewStyle().
			Foreground(tokyoRed).
			Bold(true)

	serverBox = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(tokyoPurple).
			Padding(1).
			MarginBottom(1).
			AlignHorizontal(lipgloss.Left) // Force left alignment

	toolNameStyle = lipgloss.NewStyle().
			Foreground(tokyoCyan).
			Bold(true)

	descriptionStyle = lipgloss.NewStyle().
				Foreground(tokyoFg).
				PaddingBottom(1)
	configFile string
)
