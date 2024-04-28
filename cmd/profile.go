package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var profileCmd = &cobra.Command{
	Use:   "profile [command]",
	Short: "Configuration for opensearch cluster",
	Long: `profile (thothica profile) is used to manage connection configuratio
    for the underlying opensearch cluster. A profile needs to be selected to use
    this tool.`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("profile called")
	},
}

func init() {
	rootCmd.AddCommand(profileCmd)
}
