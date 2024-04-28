package profile

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	profiles []Profile

	listCmd = &cobra.Command{
		Use:   "list",
		Short: "Lists all exsisting opensearch configuration",
		Long: `(thothica profile list) lists all the exsisting connection configurations
    used to connect this cli tool to opensearch cluster.
    `,

		Run: func(cmd *cobra.Command, args []string) {
			if err := viper.UnmarshalKey("profiles", &profiles); err != nil {
				cobra.CheckErr(err)
			}

			for _, profile := range profiles {
				fmt.Print("\n")
				fmt.Printf("Profile Name %v :-\n", profile.Name)
				fmt.Printf("    endpoint    - %v\n", profile.Endpoint)
				fmt.Printf("    user        - %v\n", profile.User)
				fmt.Printf("    password    - %v\n\n", profile.Password)
			}
		},
	}
)
