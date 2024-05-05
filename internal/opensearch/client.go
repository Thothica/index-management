package opensearch

import (
	"crypto/tls"
	"net/http"

	"github.com/Thothica/thothica/cmd/profile"
	"github.com/opensearch-project/opensearch-go"
	"github.com/spf13/cobra"
)

type Client struct {
	*opensearch.Client
}

func NewClient() Client {
	currentProfile := profile.GetCurrentProfile()

	opensearchConfig := &opensearch.Config{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
		Addresses: []string{currentProfile.Endpoint},
		Username:  currentProfile.User,
		Password:  currentProfile.Password,
	}

	client, err := opensearch.NewClient(*opensearchConfig)
	if err != nil {
		cobra.CheckErr(err)
	}

	return Client{client}
}
