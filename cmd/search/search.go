package search

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

var (
	focusedStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	blurredStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	noStyle       = lipgloss.NewStyle()
	focusedButton = focusedStyle.Copy().Render("[ Submit ]")
	blurredButton = fmt.Sprintf("[ %s ]", blurredStyle.Render("Submit"))
	SearchCmd     = &cobra.Command{
		Use:   "search",
		Short: "perform semantic search on index.",
		Long:  "performs semantic search on the provided index.",
		Run: func(cmd *cobra.Command, args []string) {
			if _, err := tea.NewProgram(initialModel()).Run(); err != nil {
				cobra.CheckErr(err)
			}
		},
	}
)

type model struct {
	focusIndex int
	inputs     []textinput.Model
}

func initialModel() model {
	m := model{
		inputs:     make([]textinput.Model, 2),
		focusIndex: 0,
	}

	for i := range m.inputs {
		t := textinput.New()
		t.CharLimit = 32

		switch i {
		case 0:
			t.Placeholder = "Index Name"
			t.CharLimit = 50
			t.PromptStyle = focusedStyle
			t.TextStyle = focusedStyle
			t.Focus()
		case 1:
			t.Placeholder = "Search query"
			t.CharLimit = 100
		}

		m.inputs[i] = t

	}
	return m
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			return m, tea.Quit

		case "tab", "shift+tab", "enter", "up", "down":
			s := msg.String()

			if s == "enter" && m.focusIndex == len(m.inputs) {
				return m, tea.Quit
			}

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
			for idx := range m.inputs {
				if idx == m.focusIndex {
					cmds[idx] = m.inputs[idx].Focus()
					m.inputs[idx].PromptStyle = focusedStyle
					m.inputs[idx].TextStyle = focusedStyle
					continue
				}
				m.inputs[idx].Blur()
				m.inputs[idx].PromptStyle = noStyle
				m.inputs[idx].TextStyle = noStyle
			}

			return m, tea.Batch(cmds...)
		}
	}

	updateInputs := func(msg tea.Msg) tea.Cmd {
		cmds := make([]tea.Cmd, len(m.inputs))

		for i := range m.inputs {
			m.inputs[i], cmds[i] = m.inputs[i].Update(msg)
		}

		return tea.Batch(cmds...)
	}(msg)

	return m, updateInputs
}

func (m model) View() string {
	var b strings.Builder

	for i := range m.inputs {
		b.WriteString(m.inputs[i].View())
		if i < len(m.inputs)-1 {
			b.WriteRune('\n')
		}
	}

	button := &blurredButton
	if m.focusIndex == len(m.inputs) {
		button = &focusedButton
	}
	fmt.Fprintf(&b, "\n\n%s\n\n", *button)

	return b.String()
}
