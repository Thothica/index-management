package profile

import (
	"github.com/spf13/cobra"
)

var (
	pingCmd = &cobra.Command{
		Use:   "ping",
		Short: "pings the opensearch cluster with current configuration",
		Long: `(thothica profile ping) pings the opensearch cluster
        Useful to check connection status.`,

		Run: func(cmd *cobra.Command, args []string) {
		},
	}
)
