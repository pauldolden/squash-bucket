package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	figure "github.com/common-nighthawk/go-figure"
)

type viewTypes int

const (
	formView viewTypes = iota
	loadingView
)

type Model struct {
	state      viewTypes
	focusIndex int
	inputs     []textinput.Model
	cursorMode textinput.CursorMode
}

func NewModel() Model {
	m := Model{
		state:  formView,
		inputs: make([]textinput.Model, 5),
	}

	var t textinput.Model
	for i := range m.inputs {
		t = textinput.New()
		t.CursorStyle = CursorStyle
		t.CharLimit = 32
		t.Prompt = ""

		switch i {
		case 0:
			t.Placeholder = "Image Path"
			t.Focus()
			t.PromptStyle = FocusedStyle
			t.TextStyle = FocusedStyle
		case 1:
			t.Placeholder = "Compression Quality"
			t.CharLimit = 3
		case 2:
			t.Placeholder = "Resize Width"
			t.CharLimit = 4
		case 3:
			t.Placeholder = "Resize Height"
			t.CharLimit = 4
		case 4:
			t.Placeholder = "Bucket Name"
			t.CharLimit = 64
		}

		m.inputs[i] = t
	}

	return m
}

func (m *Model) SetActiveView(view viewTypes) {
	m.state = view
}

func (m *Model) GetActiveView() viewTypes {
	return m.state
}

func (m Model) Init() tea.Cmd {
	return textinput.Blink
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			return m, tea.Quit

		// Change cursor mode
		case "ctrl+r":
			m.cursorMode++
			if m.cursorMode > textinput.CursorHide {
				m.cursorMode = textinput.CursorBlink
			}
			cmds := make([]tea.Cmd, len(m.inputs))
			for i := range m.inputs {
				cmds[i] = m.inputs[i].SetCursorMode(m.cursorMode)
			}
			return m, tea.Batch(cmds...)

		// Set focus to next input
		case "tab", "shift+tab", "enter", "up", "down":
			s := msg.String()

			// Did the user press enter while the submit button was focused?
			// If so, exit.
			if s == "enter" && m.focusIndex == len(m.inputs) {
				m.SetActiveView(loadingView)
			}

			// Cycle indexes
			if s == "up" || s == "shift+tab" {
				m.focusIndex--
			} else {
				m.focusIndex++
			}

			if m.focusIndex > len(m.inputs) {
				m.focusIndex = 0
			} else if m.focusIndex < 0 {
				m.focusIndex = len(m.inputs)
			}

			cmds := make([]tea.Cmd, len(m.inputs))
			for i := 0; i <= len(m.inputs)-1; i++ {
				if i == m.focusIndex {
					// Set focused state
					cmds[i] = m.inputs[i].Focus()
					m.inputs[i].PromptStyle = FocusedStyle
					m.inputs[i].TextStyle = FocusedStyle
					continue
				}
				// Remove focused state
				m.inputs[i].Blur()
				m.inputs[i].PromptStyle = NoStyle
				m.inputs[i].TextStyle = NoStyle
			}

			return m, tea.Batch(cmds...)
		}
	}

	// Handle character input and blinking
	cmd := m.updateInputs(msg)

	return m, cmd
}

func (m Model) View() string {
	s := ""

	s += m.buildForm()

	if m.state == loadingView {
		s += "Loading..."
	}

	return s
}

func (m Model) buildForm() string {
	var b strings.Builder
	logo := figure.NewFigure("Squash->Bucket", "slant", true)
	b.WriteString(TitleStyle.Render(logo.String()))
	b.WriteString("\n")

	for i := range m.inputs {
		b.WriteString(m.inputs[i].View())
		if i < len(m.inputs)-1 {
			b.WriteRune('\n')
		}
	}

	button := &BlurredButton
	if m.focusIndex == len(m.inputs) {
		button = &FocusedButton
	}
	fmt.Fprintf(&b, "\n\n%s\n\n", *button)

	b.WriteString(HelpStyle.Render("cursor mode is "))
	b.WriteString(CursorModeHelpStyle.Render(m.cursorMode.String()))
	b.WriteString(HelpStyle.Render(" (ctrl+r to change style)"))

	return b.String()
}

func (m *Model) updateInputs(msg tea.Msg) tea.Cmd {
	cmds := make([]tea.Cmd, len(m.inputs))

	// Only text inputs with Focus() set will respond, so it's safe to simply
	// update all of them here without any further logic.
	for i := range m.inputs {
		m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
	}

	return tea.Batch(cmds...)
}
