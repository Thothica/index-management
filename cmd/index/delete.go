package index

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	indexName string
	deleteCmd = &cobra.Command{
		Use:   "delete",
		Short: "deletes the index.",
		Long:  `This command will delete the give pipeline if it exsists.`,
		Run: func(cmd *cobra.Command, args []string) {
			res, err := c.DeleteIndex(indexName)
			if err != nil {
				cobra.CheckErr(err)
			}
			fmt.Println(res)
		},
	}
)

func init() {
	deleteCmd.Flags().StringVarP(&indexName, "index-name", "i", "", "Name of embedding index to delete.")
	deleteCmd.MarkFlagRequired("index-name")
}
