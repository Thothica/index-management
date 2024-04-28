package profile

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all exsisting opensearch configuration",
	Long: `(thothica profile list) lists all the exsisting connection configurations
    used to connect this cli tool to opensearch cluster.
    `,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list profile called")
	},
}
