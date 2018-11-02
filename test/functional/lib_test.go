package functional

import (
	"github.com/stretchr/testify/assert"
	"github.com/swellaby/ci-detective/cidetective"
	"os"
	"testing"
)

func init() {
	for _, envVar := range expCIEngineEnvVars {
		os.Unsetenv(envVar)
	}
}

func TestLibShouldReturnTrueWhenEnvVarIsSet(t *testing.T) {
	for _, envVar := range expCIEngineEnvVars {
		os.Setenv(envVar, "")
		assert.True(t, cidetective.IsCI())
		os.Unsetenv(envVar)
	}
}

func TestLibShouldReturnFalseWhenEnvVarIsNotSet(t *testing.T) {
	assert.False(t, cidetective.IsCI())
}
