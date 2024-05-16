package pipeline

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	pipelineName string
	deleteCmd    = &cobra.Command{
		Use:   "delete",
		Short: "deletes the pipeline.",
		Long:  `This command will delete the give pipeline if it exsists.`,
		Run: func(cmd *cobra.Command, args []string) {
			res, err := c.DeletePipeline(pipelineName)
			if err != nil {
				cobra.CheckErr(err)
			}
			fmt.Println(res)
		},
	}
)

func init() {
	deleteCmd.Flags().StringVarP(&pipelineName, "pipeline-name", "p", "", "Name of embedding pipeline to delete.")
	deleteCmd.MarkFlagRequired("pipeline-name")
}
