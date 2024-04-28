package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "thothica [command]",
	Short: "An internal cli too to manage semantic search",
	Long: `
    Welcome to Thothica cli !!
    You can manage semantic search cluster using this tool.

    This tool allows you to do many things like
    - Manage data indices
    - Perform semantic search
    - Ingest data
    `,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
        log.Fatal(err)
    }
}

