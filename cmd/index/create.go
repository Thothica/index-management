package index

import (
	"github.com/spf13/cobra"
)

var (
	createCmd = &cobra.Command{
		Use:   "create",
		Short: "Creates a index for semantic search.",
		Long:  `create (thothica index create) is a tool which helps you to define mappings and create index in opensearch to perform semantic search.`,
		Run: func(cmd *cobra.Command, args []string) {
		},
	}
)
