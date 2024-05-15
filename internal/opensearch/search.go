package opensearch

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/opensearch-project/opensearch-go/opensearchapi"
)

// NOTE: Change model_id and mapping column (Text_embedding) in searchBody as per your need.
// This function not meant to be consumed without editing the mapping column inside !!
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
                                        "model_id": "AbDZGo8BB3UUeZ_94CHA",
                                        "k": 5
                                }
                        }
                },
                "size": %v
        }`, query, size))

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
