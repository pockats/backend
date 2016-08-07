package commands

import (
	"fmt"
	"sort"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "List all configuration keys and values",
	Long:  "List all Viper configuration keys and values in alphabetical order",
	Run: func(cmd *cobra.Command, args []string) {
		settings := viper.AllSettings()
		keys := viper.AllKeys()
		sort.Strings(keys)

		fmt.Printf("Configurations list (records are sorted in alphabetical order)\n\n")

		for _, k := range keys {
			fmt.Printf("%s:\"%+v\"\n", k, settings[k])
		}
	},
}
