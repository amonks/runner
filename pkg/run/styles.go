package run

import (
	"github.com/amonks/run/internal/color"
	"github.com/charmbracelet/lipgloss"
)

var (
	logStyle = lipgloss.NewStyle().
			Foreground(color.XXXLight).
			Italic(true)
	errorStyle = lipgloss.NewStyle().
			Foreground(color.Red).
			Italic(true)
)
