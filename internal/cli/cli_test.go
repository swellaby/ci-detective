package cli

import (
	"github.com/stretchr/testify/assert"
	"github.com/swellaby/ci-detective/cidetective"
	"testing"
)

func TestRootCommandConfiguresUseCorrectly(t *testing.T) {
	assert.Equal(t, "ci-detective", rootCmd.Use)
}

func TestRootCommandConfiguresVersionCorrectly(t *testing.T) {
	assert.Equal(t, cidetective.Version, rootCmd.Version)
}

func TestGetRunnerReturnsRootCommand(t *testing.T) {
	assert.Equal(t, rootCmd, GetRunner())
}
