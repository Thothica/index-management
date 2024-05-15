package opensearch

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/opensearch-project/opensearch-go/opensearchapi"
)

func (c Client) CreateIndex(Body, Index string) (string, error) {
	req := opensearchapi.IndicesCreateRequest{
		Index: Index,
		Body:  strings.NewReader(Body),
	}

	res, err := req.Do(context.Background(), c.Client)
	if err != nil {
		return "", err
	}

	var resp map[string]interface{}
	err = json.NewDecoder(res.Body).Decode(&resp)
	resopnse, err := json.MarshalIndent(resp, "", " ")
	if err != nil {
		return "", err
	}

	if res.Status() != "200 OK" {
		return "", fmt.Errorf("opensearch returned error\n %s", string(resopnse))
	}

	return string(resopnse), nil
}
