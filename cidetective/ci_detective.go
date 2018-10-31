package cidetective

import "os"

var lookupEnv = os.LookupEnv

// IsCI returns a boolean indicating whether or not
// the current environment is a CI server.
func IsCI() bool {
	for _, env := range ciEngineEnvVars {
		if _, exists := lookupEnv(env); exists {
			return true
		}
	}
	return false
}
