package search

import (
	"fmt"
	"log"
	"strings"

	"github.com/Thothica/thothica/internal/opensearch"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

var (
	c             = opensearch.NewClient()
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
	spinner    spinner.Model
	loading    bool
	complete   bool
}

type SemanticSearchResponse string
type SemanticSearchError error

func SearchRequest(query, index string) tea.Cmd {
	return func() tea.Msg {
		res, err := c.SemanticSearch(query, index, 5)
		if err != nil {
			return SemanticSearchError(err)
		}
		return SemanticSearchResponse(res)
	}
}

func initialModel() model {
	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))

	m := model{
		inputs:  make([]textinput.Model, 2),
		spinner: s,
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
				m.loading = true
				return m, tea.Batch(m.spinner.Tick, SearchRequest(m.inputs[1].Value(), m.inputs[0].Value()))
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
	case SemanticSearchResponse:
		fmt.Println(msg)
		m.loading = false
		m.complete = true
		return m, tea.Quit

	case SemanticSearchError:
		fmt.Println(msg)
		m.loading = false
		m.complete = true
		return m, tea.Quit

	default:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
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

	if !m.complete {
		for i := range m.inputs {
			b.WriteString(m.inputs[i].View())
			if i < len(m.inputs)-1 {
				b.WriteRune('\n')
			}
		}
	}

	button := &blurredButton
	if m.focusIndex == len(m.inputs) {
		button = &focusedButton
	}

	if m.loading {
		fmt.Fprintf(&b, "\n\n   %s Searching your query...\n\n", m.spinner.View())
	}
	if !m.complete {
		fmt.Fprintf(&b, "\n\n%s\n\n", *button)
	}

	return b.String()
}
