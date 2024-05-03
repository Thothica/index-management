package search

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

var (
	SearchCmd = &cobra.Command{
		Use:   "search",
		Short: "perform semantic search on index.",
		Long:  "perfoms semantic search on the provided index.",
		Run: func(cmd *cobra.Command, args []string) {
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
			t.Focus()
		case 1:
			t.Placeholder = "Search"
			t.CharLimit = 100
		}

		m.inputs[i] = t

	}
	return m
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}
