package pipeline

import "github.com/spf13/cobra"

var (
	pipelineCmd = &cobra.Command{
		Use:   "pipeline",
		Short: "Create and view pipelines.",
		Long:  "Allows you to create and view pipelines in opensearch.",
	}
)
