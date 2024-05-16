package opensearch

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/Thothica/thothica/cmd/profile"
	"github.com/opensearch-project/opensearch-go/opensearchapi"
)

func (c Client) SemanticSearch(query, index string, size int) (string, error) {
	searchBody := strings.NewReader(fmt.Sprintf(`{
                "_source": {
                        "excludes": [
                                "Text_embedding"
                        ]
                },
                "query": {
                        "neural": {
                                "Text_embedding": {
                                        "query_text": "%v",
                                        "model_id": "%v",
                                        "k": 5
                                }
                        }
                },
                "size": %v
        }`, query, profile.GetModelID(), size))

	semanticSearchRequest := opensearchapi.SearchRequest{
		Index: []string{index},
		Body:  searchBody,
	}

	var searchResponse map[string]interface{}

	res, err := semanticSearchRequest.Do(context.Background(), c.Client)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&searchResponse)
	if err != nil {
		return "", err
	}

	data := searchResponse["hits"].(map[string]interface{})["hits"].([]interface{})

	searchResults, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return "", err
	}

	return string(searchResults), nil
}
