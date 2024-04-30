package profile

import (
	"github.com/spf13/cobra"
)

var (
	check_profiles []Profile

	useCmd = &cobra.Command{
		Use:   "use [name]",
		Short: "Uses the specified connection profile",
		Long: `(thothica profile use [name]) uses the specified profile's credentials
        to connect to opensearch cluster`,

		Run: func(cmd *cobra.Command, args []string) {
		},
	}
)
