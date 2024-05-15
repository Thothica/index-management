package index

import (
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
	i           index
	createCmd   = &cobra.Command{
		Use:   "create",
		Short: "Creates a index for semantic search.",
		Long:  `create (thothica index create) is a tool which helps you to define mappings and create index in opensearch to perform semantic search.`,
		Run: func(cmd *cobra.Command, args []string) {
			InitialseIndex(&i)
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
