package profile

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	listCmd = &cobra.Command{
		Use:   "list",
		Short: "Lists all exsisting opensearch configuration",
		Long: `(thothica profile list) lists all the exsisting connection configurations
    used to connect this cli tool to opensearch cluster.
    `,

		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("\nProfile in use: %v\n", c.Current)
			for _, profile := range c.Profiles {
				fmt.Print("\n")
				fmt.Printf("Profile Name    - %v\n", profile.Name)
				fmt.Printf("    endpoint    - %v\n", profile.Endpoint)
				fmt.Printf("    user        - %v\n", profile.User)
				fmt.Printf("    password    - %v\n", profile.Password)
				fmt.Printf("    model-id    - %v\n\n", profile.DefaultModel)
			}
		},
	}
)
