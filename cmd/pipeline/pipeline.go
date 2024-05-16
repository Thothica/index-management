package pipeline

import (
	"github.com/Thothica/thothica/internal/opensearch"
	"github.com/spf13/cobra"
)

var (
	c           = opensearch.NewClient()
	PipelineCmd = &cobra.Command{
		Use:   "pipeline",
		Short: "Create and view pipelines.",
		Long:  "Allows you to create and view pipelines in opensearch.",
	}
)

func init() {
	PipelineCmd.AddCommand(createCmd)
	PipelineCmd.AddCommand(listCmd)
	PipelineCmd.AddCommand(deleteCmd)
}
