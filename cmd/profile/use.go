package profile

import (
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

		Run: func(cmd *cobra.Command, args []string) {

			if len(args) != 1 {
				cobra.CheckErr("Usage thothica profile use [name]")
			}

			if err := viper.UnmarshalKey("profiles", &check_profiles); err != nil {
				cobra.CheckErr(err)
			}

			for _, profile := range check_profiles {
				if profile.Name == args[0] {
					currentProfile = profile
					return
				}
			}

			cobra.CheckErr("profile name " + args[0] + " does not exsist")

		},
	}
)
