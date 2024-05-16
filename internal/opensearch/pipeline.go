package opensearch

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/Thothica/thothica/cmd/profile"
	"github.com/opensearch-project/opensearch-go/opensearchapi"
)

func (c Client) GetPipeline() (string, error) {
	req := opensearchapi.IngestGetPipelineRequest{}

	res, err := req.Do(context.Background(), c.Client)
	if err != nil {
		return "", err
	}

	var resp map[string]interface{}
	err = json.NewDecoder(res.Body).Decode(&resp)
	response, err := json.MarshalIndent(resp, "", " ")
	if err != nil {
		return "", err
	}

	if res.Status() != "200 OK" {
		return "", fmt.Errorf(string(response))
	}

	return string(response), nil
}

func (c Client) CreatePipeline(Pipeline, InputField, OutputField string) (string, error) {
	Body := strings.NewReader(fmt.Sprintf(`{
		"description": "embedding pipeline",
		"processors": [
			{
			"text_embedding": {
				"model_id": "%s",
				"field_map": {
					"%s": "%s"
					}
				}
			}
		]
	}`, profile.GetModelID(), InputField, OutputField))

	req := opensearchapi.IngestPutPipelineRequest{
		PipelineID: Pipeline,
		Body:       Body,
	}

	res, err := req.Do(context.Background(), c.Client)
	if err != nil {
		return "", err
	}

	var resp map[string]interface{}
	err = json.NewDecoder(res.Body).Decode(&resp)
	response, err := json.MarshalIndent(resp, "", " ")
	if err != nil {
		return "", err
	}

	if res.Status() != "200 OK" {
		return "", fmt.Errorf(string(response))
	}

	return string(response), nil
}

func (c Client) DeletePipeline(PipelineName string) (string, error) {
	req := opensearchapi.IngestDeletePipelineRequest{
		PipelineID: PipelineName,
	}

	res, err := req.Do(context.Background(), c.Client)
	if err != nil {
		return "", err
	}

	var resp map[string]interface{}
	err = json.NewDecoder(res.Body).Decode(&resp)
	response, err := json.MarshalIndent(resp, "", " ")
	if err != nil {
		return "", err
	}

	if res.Status() != "200 OK" {
		return "", fmt.Errorf(string(response))
	}

	return string(response), nil
}
