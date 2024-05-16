package index

import (
	"github.com/Thothica/thothica/internal/opensearch"
	"github.com/spf13/cobra"
)

var (
	client   = opensearch.NewClient()
	IndexCmd = &cobra.Command{
		Use:   "index [command]",
		Short: "Manages opensearch indices",
		Long: `index (thothica index [command]) is used to manage opensearch data index.
        Use this to create or use data indices present on cluster.
        `,
	}
)

func init() {
	IndexCmd.AddCommand(listCmd)
	IndexCmd.AddCommand(createCmd)
	IndexCmd.AddCommand(deleteCmd)
}
