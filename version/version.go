// Package version contains the application's version information.
package version

import (
	"errors"
	"fmt"
	"os"

	"github.com/blang/semver"
)

var (
	// Code is the application version code in SemVer method
	Code = "0.0.0"
	// CommitHash is the revision hash of the build's git repository
	CommitHash string
	// BuildTime is the build's time
	BuildTime string
	// SemVer is version code in semantic versioning format
	SemVer *semver.Version
)

func init() {
	var err error
	SemVer, err = semver.New(Code)
	if err != nil {
		fmt.Println(errors.New("Version format is invalid"))
		os.Exit(-1)
	}
}
