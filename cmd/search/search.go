package search

import "github.com/spf13/cobra"

var (
	SearchCmd = &cobra.Command{
		Use:   "search",
		Short: "perform semantic search on index.",
		Long:  "perfoms semantic search on the provided index.",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}
)
