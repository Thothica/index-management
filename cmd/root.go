package cmd

import (
	"github.com/Thothica/thothica/cmd/profile"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "thothica [command]",
	Short: "An internal cli too to manage semantic search",
	Long: `Welcome to Thothica cli !!
    You can manage semantic search cluster using this tool.

    This tool allows you to do many things like
    - Manage data indices
    - Perform semantic search
    - Ingest data
    `,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		cobra.CheckErr(err)
	}
}

func init() {
	rootCmd.AddCommand(profile.ProfileCmd)
}
