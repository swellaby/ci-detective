package cidetective

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var expCIEngineEnvVars = [...]string{
	"CI",
	"CONTINUOUS_INTEGRATION",
	"BUILD_NUMBER",
	"APPVEYOR",
	"bamboo_planKey",
	"BITBUCKET_COMMIT",
	"CIRCLECI",
	"TRAVIS",
}

func TestEnginesCheckedInCorrectOrder(t *testing.T) {
	for i, actVar := range ciEngineEnvVars {
		assert.Equal(t, expCIEngineEnvVars[i], actVar)
	}
}
