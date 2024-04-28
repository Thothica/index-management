package profile

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Profile struct {
	Name     string `yaml:"name"`
	Endpoint string `yaml:"endpoint"`
	Admin    string `yaml:"admin"`
	Password string `yaml:"password"`
}

var (
	cfgFile    string
	ProfileCmd = &cobra.Command{
		Use:   "profile [command]",
		Short: "Configuration for opensearch cluster",
		Long: `profile (thothica profile) is used to manage connection configuration
    for the underlying opensearch cluster. A profile needs to be selected to use
    for this cli.`,
	}
)

func init() {
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")
	viper.SetConfigName(".thothica")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			cobra.CheckErr("No config file found, use thothica profile create to create a profile")
		}
		cobra.CheckErr(err)
	}

	ProfileCmd.AddCommand(listCmd)
	ProfileCmd.AddCommand(createCmd)
}
