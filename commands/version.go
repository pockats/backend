package commands

import (
	"fmt"

	"github.com/pockats/pockats-backend/helpers"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Pockats",
	Long: `Packats follows [SemVer](http://semver.org/) versioning system, visit
SemVer official docs for more informations.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%s\n", helpers.PockatsVersion())
	},
}
