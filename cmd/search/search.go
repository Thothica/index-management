package search

import (
	"fmt"

	"github.com/Thothica/thothica/internal/opensearch"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

var (
	Index         string
	Query         string
	Size          int
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
			res, err := c.SemanticSearch(Query, Index, Size)
			if err != nil {
				cobra.CheckErr(err)
			}
			fmt.Println(res)
		},
	}
)

func init() {
	SearchCmd.Flags().StringVarP(&Index, "index", "i", "", "Index on which semantic search will be performed.")
	SearchCmd.MarkFlagRequired("index")
	SearchCmd.Flags().StringVarP(&Query, "query", "q", "", "Query for semantic search.")
	SearchCmd.MarkFlagRequired("query")
	SearchCmd.Flags().IntVarP(&Size, "size", "s", 5, "Result size for returning response.")
}
