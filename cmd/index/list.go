package index

import (
	"context"
	"fmt"

	"github.com/opensearch-project/opensearch-go/opensearchapi"
	"github.com/spf13/cobra"
)

var (
	listCmd = &cobra.Command{
		Use:   "list",
		Short: "lists present indices.",
		Long:  `list (thothica index list) will display all present indices.`,
		Run: func(cmd *cobra.Command, args []string) {
			req := opensearchapi.CatIndicesRequest{}
			res, err := req.Do(context.Background(), client)
			if err != nil {
				cobra.CheckErr(err)
			}
			fmt.Println(res)
		},
	}
)
