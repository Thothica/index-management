package profile

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	Username string
	Password string
	Endpoint string
    Name    string

	createCmd = &cobra.Command{
		Use:   "create [flags]",
		Short: "Creates new connection configuration",
		Long: `(thothica profile create [flags]) creates new connection configuration which will be
    used to connect this cli to a opensearch cluster.
    `,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("create profile called")
		},
	}
)

func init() {
	createCmd.Flags().StringVarP(&Endpoint, "endpoint", "e", "", "Full endpoint of opensearch cluster")
	createCmd.MarkFlagRequired("endpoint")
	createCmd.Flags().StringVarP(&Username, "username", "u", "", "Username for opensearch cluster")
	createCmd.MarkFlagRequired("username")
	createCmd.Flags().StringVarP(&Password, "password", "p", "", "Password for opensearch cluster")
	createCmd.MarkFlagRequired("password")
	createCmd.Flags().StringVarP(&Name, "name", "n", "", "Name for connection profile")
	createCmd.MarkFlagRequired("name")
}
