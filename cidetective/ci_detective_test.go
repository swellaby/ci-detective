package cidetective

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsCI(t *testing.T) {
	assert.False(t, IsCI())
}
