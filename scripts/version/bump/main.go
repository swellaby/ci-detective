package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"runtime"

	"github.com/Masterminds/semver"
	"github.com/swellaby/ci-detective/cidetective"
)

const versionFileContentsPrefix = `package cidetective

// Version defines the current version of the package
const Version = "`

const versionFileContentsSuffix = `"
`

func writeUpdatedVersionFile(bumpedVersion string) {
	_, currentFilePath, _, _ := runtime.Caller(0)
	versionFilepath := filepath.Join(path.Dir(currentFilePath), "../../../cidetective/version.go")
	fileContents := versionFileContentsPrefix + bumpedVersion + versionFileContentsSuffix
	err := ioutil.WriteFile(versionFilepath, []byte(fileContents), os.ModePerm)
	if err != nil {
		fmt.Printf("Failed to updated version file. Error: %s", err)
		os.Exit(1)
	}
}

func main() {
	origVersion := cidetective.Version
	version, err := semver.NewVersion(origVersion)
	if err != nil {
		fmt.Printf("Failed to parse version. Error: %s", err)
		os.Exit(1)
	}
	bumpedVersion := version.IncPatch()
	writeUpdatedVersionFile(bumpedVersion.String())
	fmt.Printf(bumpedVersion.String())
}
