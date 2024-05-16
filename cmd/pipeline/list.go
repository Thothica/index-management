package pipeline

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	listCmd = &cobra.Command{
		Use:   "list",
		Short: "lists all pipelines.",
		Long:  "lists data pipeline present in cluster \n Most common use case will be to see currently present embedding pipelines",
		Run: func(cmd *cobra.Command, args []string) {
			res, err := c.GetPipeline()
			if err != nil {
				cobra.CheckErr(err)
			}
			fmt.Println(res)
		},
	}
)
