package commands

import (
	"fmt"

	"github.com/pockats/pockats-backend/version"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version information of Pockats",
	Long: `Packats follows [SemVer](http://semver.org/) versioning system, visit
SemVer official docs for more informations.`,
	Run: func(cmd *cobra.Command, args []string) {
		var versionStr string

		versionStr = fmt.Sprintf("v%s", version.Code)

		if version.CommitHash != "" {
			versionStr += fmt.Sprintf(" - %s", version.CommitHash)
		}
		if version.BuildTime != "" {
			versionStr += fmt.Sprintf(" - %s", version.BuildTime)
		}

		fmt.Printf("Pockats %s\n", versionStr)
	},
}
