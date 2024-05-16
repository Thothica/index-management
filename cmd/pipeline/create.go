package pipeline

import "github.com/spf13/cobra"

var (
	InputField  string
	OutputField string
	createCmd   = &cobra.Command{
		Use:   "create",
		Short: "Creates embedding pipeline.",
		Long:  "Creates data pipeline with text_embedding preprocessor for data ingestion and searching.",
		Run:   func(cmd *cobra.Command, args []string) {},
	}
)

func init() {
	createCmd.Flags().StringVarP(&InputField, "input-field", "i", "", "Name of field containing text that you want to want to make semantic search on.")
	createCmd.MarkFlagRequired("input-field")
	createCmd.Flags().StringVarP(&OutputField, "output-field", "o", "", "Name of output field which will contain vectors.")
	createCmd.MarkFlagRequired("output-field")
}
