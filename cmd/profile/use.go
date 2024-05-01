package profile

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	check_profiles []Profile

	useCmd = &cobra.Command{
		Use:   "use [name]",
		Short: "Uses the specified connection profile",
		Long: `(thothica profile use [name]) uses the specified profile's credentials
        to connect to opensearch cluster`,
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			profileToSwitch := args[0]
			for _, profile := range c.Profiles {
				if profile.Name == profileToSwitch {
					c.Current = profileToSwitch
					fmt.Printf("Now using profile - %v\n", c.Current)
					viper.Set("current-profile", c.Current)
					if err := viper.WriteConfig(); err != nil {
						cobra.CheckErr(err)
					}
					return
				}
			}
			fmt.Printf("Profile - %v does not exist\n", profileToSwitch)
		},
	}
)
