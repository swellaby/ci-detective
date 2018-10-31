package cidetective

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsCIReturnsFalseWhenNoEnvVarsFound(t *testing.T) {
	origFunc := lookupEnv
	defer func() { lookupEnv = origFunc }()
	lookupEnv = func(string) (string, bool) {
		return "", false
	}
	assert.False(t, IsCI())
}

func TestIsCIReturnsTrueWhenEnvVarFound(t *testing.T) {
	origFunc := lookupEnv
	defer func() { lookupEnv = origFunc }()
	lookupEnv = func(key string) (string, bool) {
		if key == "CIRCLECI" {
			return "", true
		}
		return "", false
	}
	assert.True(t, IsCI())
}
