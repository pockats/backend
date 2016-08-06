// Package commands defines command-line interface for Pockats. The implementation
// of this packages is based on Cobra.
package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var configFile string

// PockatsCmd represents the base command when called without any subcommands
var PockatsCmd = &cobra.Command{
	Use:   "pockats-backend",
	Short: "get awesome reports about your Pocket's reading list",
	Long:  "Pockats is a tool for gathering reports about your Pocket's reading list",
}

// Setup adds all child commands to the root command and sets flags appropriately.
func Setup() {
	if err := PockatsCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

}

func init() {
	cobra.OnInitialize(initializeConfig)

	PockatsCmd.PersistentFlags().StringVar(&configFile, "config", "", "config file (default is $HOME/.pockats-backend.yaml)")
}

// initializeConfig loads the configurations from `.pockats-backend` file or Env
// variables with `POCKATS_` prefix.
func initializeConfig() {
	if configFile != "" {
		viper.SetConfigFile(configFile)
	}

	viper.SetConfigName(".pockats-backend")
	viper.AddConfigPath("$HOME")
	viper.AutomaticEnv()
	viper.SetEnvPrefix("pockats")

	viper.ReadInConfig()
}
