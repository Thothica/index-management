package index

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

type index struct {
	Settings struct {
		IndexKNN        bool   `json:"index.knn"`
		DefaultPipeline string `json:"default_pipeline"`
	} `json:"settings"`
	Mappings struct {
		Properties map[string]interface{} `json:"properties"`
	} `json:"mappings"`
}

var (
	Pipeline    string
	VectorIndex string
	Filename    string
	data        []map[string]interface{}
    BadFormatError = errors.New("Wrong file format, only json file of format - \n \t[{\"data_key\":\"data_value\", ..}..{}] are accepted")
	i           index
	createCmd   = &cobra.Command{
		Use:   "create",
		Short: "Creates a index for semantic search.",
		Long:  `create (thothica index create) is a tool which helps you to define mappings and create index in opensearch to perform semantic search.`,
		Run: func(cmd *cobra.Command, args []string) {
			InitialseIndex(&i)

			filetype := strings.Split(Filename, ".")[1]
			if filetype != "json" {
				cobra.CheckErr(BadFormatError)
			}

			file, err := os.Open(Filename)
			if err != nil {
				cobra.CheckErr(err)
			}

			err = json.NewDecoder(file).Decode(&data)
			if err != nil {
				cobra.CheckErr(err)
			}

			for k, v := range data[0] {
				switch v.(type) {
				case string:
					i.Mappings.Properties[k] = map[string]string{
						"type": "text",
					}
				case float32:
					i.Mappings.Properties[k] = map[string]string{
						"type": "float",
					}
				case int:
					i.Mappings.Properties[k] = map[string]string{
						"type": "integer",
					}
				}
			}

			IndexJson, err := json.MarshalIndent(i, "", "   ")
			if err != nil {
				cobra.CheckErr(err)
			}

			fmt.Println(string(IndexJson))
		},
	}
)

func init() {
	createCmd.Flags().StringVarP(&Pipeline, "pipeline", "p", "", "Name of embedding pipeline for this index.")
	createCmd.MarkFlagRequired("pipeline")
	createCmd.Flags().StringVarP(&VectorIndex, "vector-index", "v", "", "Name of mapping/column which will hold embeddings as defined in the pipeline.")
	createCmd.MarkFlagRequired("vector-index")
	createCmd.Flags().StringVarP(&Filename, "file", "f", "", "Path of file containing data for semantic search.\n Json file should be of format - \n \t[{\"data_key\":\"data_value\", ..}..{}]")
	createCmd.MarkFlagRequired("vector-index")
}

func InitialseIndex(i *index) {
	i.Settings.IndexKNN = true
	i.Settings.DefaultPipeline = Pipeline
	i.Mappings.Properties = make(map[string]interface{})
}
