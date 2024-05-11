package index

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

var (
	Pipeline  string
	createCmd = &cobra.Command{
		Use:   "create",
		Short: "Creates a index for semantic search.",
		Long:  `create (thothica index create) is a tool which helps you to define mappings and create index in opensearch to perform semantic search.`,
		Run: func(cmd *cobra.Command, args []string) {
		},
	}
)

func init() {
	createCmd.Flags().StringVarP(&Pipeline, "pipeline", "p", "", "Name of embedding pipeline for this index.")
	createCmd.MarkFlagRequired("pipeline")
}

type index struct {
	Settings struct {
		IndexKNN        bool   `json:"index.knn"`
		DefaultPipeline string `json:"default_pipeline"`
	} `json:"settings"`
	Mappings struct {
		Properties []map[string]interface{} `json:"properties"`
	} `json:"mappings"`
}

type model struct {
	mappingFieldTypes []string
	neuralSearchIndex index
	cursor            int
	selectionComplete bool
	fieldInput        textinput.Model
}

func initialModel() model {
	input := textinput.New()
	input.Placeholder = "Input your field name."
	input.Focus()

	return model{
		mappingFieldTypes: []string{"binary", "boolean", "ip", "object", "nested", "keyword", "text", "date"},
		cursor:            0,
		neuralSearchIndex: index{
			Settings: struct {
				IndexKNN        bool   "json:\"index.knn\""
				DefaultPipeline string "json:\"default_pipeline\""
			}{IndexKNN: true, DefaultPipeline: Pipeline},
		},
		selectionComplete: false,
		fieldInput:        input,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return nil, nil
}

func (m model) View() string {
	return ""
}
