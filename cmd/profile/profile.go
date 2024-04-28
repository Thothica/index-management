package profile

import (
	"github.com/spf13/cobra"
)

var ProfileCmd = &cobra.Command{
	Use:   "profile [command]",
	Short: "Configuration for opensearch cluster",
	Long: `profile (thothica profile) is used to manage connection configuration
    for the underlying opensearch cluster. A profile needs to be selected to use
    for this cli.`,
}

func init() {
	ProfileCmd.AddCommand(listCmd)
    ProfileCmd.AddCommand(createCmd)
}
