package profile

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

type Profile struct {
	Name     string `mapstructure:"name"`
	Endpoint string `mapstructure:"endpoint"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
}

type Config struct {
	Profiles []Profile `mapstructure:"profiles"`
	Current  string    `mapstructure:"current-profile"`
}

var (
    currentProfile Profile
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
	ProfileCmd.AddCommand(useCmd)
	ProfileCmd.AddCommand(pingCmd)
}
