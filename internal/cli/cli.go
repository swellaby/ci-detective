package cli

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/swellaby/ci-detective/cidetective"
)

var log = fmt.Println
var logf = fmt.Printf

// Runner describes a CLI runner
type Runner interface {
	Execute() error
}

var rootCmd = &cobra.Command{
	Use:     "ci-detective",
	Version: cidetective.Version,
}

// GetRunner returns the runner
func GetRunner() Runner {
	return rootCmd
}
