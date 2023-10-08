package main

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	fuchsia   = lipgloss.AdaptiveColor{Light: "#EE6FF8", Dark: "#EE6FF8"}
	docStyle  = lipgloss.NewStyle().Margin(1, 2)
	helpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("241")).Render
)
