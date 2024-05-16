package cmd

import (
	"github.com/Thothica/thothica/cmd/index"
	"github.com/Thothica/thothica/cmd/pipeline"
	"github.com/Thothica/thothica/cmd/profile"
	"github.com/Thothica/thothica/cmd/search"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "thothica [command]",
	Short: "An internal cli too to manage semantic search",
	Long: `Welcome to Thothica cli !!
    You can manage semantic search cluster using this tool.

    This tool allows you to do many things like
    - Manage data indices
    - Manage data pipelines
    - Perform semantic search
    `,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		cobra.CheckErr(err)
	}
}

func init() {
	rootCmd.AddCommand(profile.ProfileCmd)
	rootCmd.AddCommand(index.IndexCmd)
	rootCmd.AddCommand(search.SearchCmd)
	rootCmd.AddCommand(pipeline.PipelineCmd)
}
