package index

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/Thothica/thothica/internal/opensearch"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

type index struct {
	Settings struct {
		IndexKNN        bool   `json:"index.knn"`
		DefaultPipeline string `json:"default_pipeline"`
	} `json:"settings"`
	Mappings struct {
		Properties map[string]interface{} `json:"properties"`
	} `json:"mappings"`
}

var (
	Pipeline       string
	VectorIndex    string
	Filename       string
	Index          string
	IndexBody      string
	data           []map[string]interface{}
	i              index
	magentaStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	c              = opensearch.NewClient()
	BadFormatError = errors.New("Wrong file format, only json file of format - \n \t[{\"data_key\":\"data_value\", ..}..{}] are accepted")
	createCmd      = &cobra.Command{
		Use:   "create",
		Short: "Creates a index for semantic search.",
		Long:  `create (thothica index create) is a tool which helps you to define mappings and create index in opensearch to perform semantic search.`,
		Run: func(cmd *cobra.Command, args []string) {
			InitialseIndex(&i)

			filetype := strings.Split(Filename, ".")[1]
			if filetype != "json" {
				cobra.CheckErr(BadFormatError)
			}

			file, err := os.Open(Filename)
			if err != nil {
				cobra.CheckErr(err)
			}

			err = json.NewDecoder(file).Decode(&data)
			if err != nil {
				cobra.CheckErr(err)
			}

			i.Mappings.Properties[VectorIndex] = map[string]interface{}{
				"type":      "knn_vector",
				"dimension": 768,
			}

			for k, v := range data[0] {
				switch v.(type) {
				case string:
					i.Mappings.Properties[k] = map[string]string{
						"type": "text",
					}
				case float32:
					i.Mappings.Properties[k] = map[string]string{
						"type": "float",
					}
				case int:
					i.Mappings.Properties[k] = map[string]string{
						"type": "integer",
					}
				}
			}

			IndexJson, err := json.MarshalIndent(i, "", "   ")
			if err != nil {
				cobra.CheckErr(err)
			}

			IndexBody = string(IndexJson)

			if _, err := tea.NewProgram(initialModel()).Run(); err != nil {
				cobra.CheckErr(err)
			}
		},
	}
)

func init() {
	createCmd.Flags().StringVarP(&Pipeline, "pipeline", "p", "", "Name of embedding pipeline for this index.")
	createCmd.MarkFlagRequired("pipeline")
	createCmd.Flags().StringVarP(&VectorIndex, "vector-index", "v", "", "Name of mapping/column which will hold embeddings as defined in the pipeline.")
	createCmd.MarkFlagRequired("vector-index")
	createCmd.Flags().StringVarP(&Filename, "file", "f", "", "Path of file containing data for semantic search.\n Json file should be of format - \n \t[{\"data_key\":\"data_value\", ..}..{}]")
	createCmd.MarkFlagRequired("vector-index")
	createCmd.Flags().StringVarP(&Index, "index", "i", "", "Name of index you are creating.")
	createCmd.MarkFlagRequired("index")
}

func InitialseIndex(i *index) {
	i.Settings.IndexKNN = true
	i.Settings.DefaultPipeline = Pipeline
	i.Mappings.Properties = make(map[string]interface{})
}

type apiResponse string
type apiError error

type model struct {
	textarea    textarea.Model
	spinner     spinner.Model
	loading     bool
	apiResponse string
}

func CreateIndex(Body, Index string) tea.Cmd {
	return func() tea.Msg {
		res, err := c.CreateIndex(Body, Index)
		if err != nil {
			return apiError(err)
		}
		return apiResponse(res)
	}
}

func initialModel() model {
	t := textarea.New()
	t.Focus()
	t.SetHeight(25)
	t.CharLimit = 0
	t.SetValue(IndexBody)
	t.FocusedStyle.CursorLine = magentaStyle

	s := spinner.New()
	s.Spinner = spinner.Dot
	s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))

	return model{
		textarea:    t,
		spinner:     s,
		apiResponse: "",
	}
}

func (m model) Init() tea.Cmd {
	return textarea.Blink
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "ctrl+s":
			m.loading = true
			return m, tea.Batch(m.spinner.Tick, CreateIndex(m.textarea.Value(), Index))
		}
	default:
		if !m.textarea.Focused() {
			cmd = m.textarea.Focus()
			cmds = append(cmds, cmd)
		}

	case apiError:
		m.loading = false
		m.apiResponse = msg.Error()
		return m, tea.Quit

	case apiResponse:
		m.loading = false
		m.apiResponse = string(msg)
		return m, tea.Quit
	}

	m.textarea, cmd = m.textarea.Update(msg)
	cmds = append(cmds, cmd)
	m.spinner, cmd = m.spinner.Update(msg)
	cmds = append(cmds, cmd)
	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	var b strings.Builder

	if m.loading {
		fmt.Fprintf(&b, "\n %s Creating Index...\n\n", m.spinner.View())
		return b.String()
	}

	if m.apiResponse != "" {
		b.WriteString(m.apiResponse)
		return b.String()
	}

	fmt.Fprintf(&b, "\n%s\n\n%s\n\n(press ctrl+c to abort or %s to create Index)", magentaStyle.Copy().Render("Edit Index as you like!"), m.textarea.View(), magentaStyle.Copy().Render("ctrl+s"))
	return b.String()
}
