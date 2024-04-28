package profile

import (
	"os"

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
	ProfileCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.thothica.yaml)")

	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".thothica")
	}

	if err := viper.ReadInConfig(); err != nil {
        cobra.CheckErr(err)
	}

	ProfileCmd.AddCommand(listCmd)
	ProfileCmd.AddCommand(createCmd)
}
