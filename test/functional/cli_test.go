package functional

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCliShouldExitWithZeroWhenEnvVarIsSet(t *testing.T) {
	for _, envVar := range expCIEngineEnvVars {
		err := buildCommandWithEnv(envVar).Run()
		assert.Nil(t, err)
	}
}

func TestCliShouldExitWithNonZeroWhenEnvVarIsNotSet(t *testing.T) {
	err := buildCommandWithoutEnv().Run()
	assert.NotNil(t, err)
}
