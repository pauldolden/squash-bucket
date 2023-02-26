package ui

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

var (
	TitleStyle          = lipgloss.NewStyle().Foreground(lipgloss.Color("#ff69b4")).Bold(true)
	FocusedStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("#ff69b4"))
	BlurredStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	CursorStyle         = FocusedStyle.Copy()
	NoStyle             = lipgloss.NewStyle()
	HelpStyle           = BlurredStyle.Copy()
	CursorModeHelpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("244"))

	FocusedButton = FocusedStyle.Copy().Render("[ Squash! ]")
	BlurredButton = fmt.Sprintf("[ %s ]", BlurredStyle.Render("Squash!"))
)
