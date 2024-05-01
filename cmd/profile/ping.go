package profile

import (
	"crypto/tls"
	"fmt"
	"net/http"

	"github.com/opensearch-project/opensearch-go"
	"github.com/spf13/cobra"
)

var (
	pingCmd = &cobra.Command{
		Use:   "ping",
		Short: "pings the opensearch cluster with current configuration",
		Long: `(thothica profile ping) pings the opensearch cluster
        Useful to check connection status.`,

		Run: func(cmd *cobra.Command, args []string) {
			currentProfile := GetCurrentProfile()
			client, err := opensearch.NewClient(opensearch.Config{
				Transport: &http.Transport{
					TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
				},
				Addresses: []string{currentProfile.Endpoint},
				Username:  currentProfile.User,
				Password:  currentProfile.Password,
			})
			if err != nil {
				cobra.CheckErr(err)
			}
			res, err := client.Ping()
			if err != nil {
				cobra.CheckErr(err)
			}
			fmt.Println(res)
		},
	}
)
