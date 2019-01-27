package cmd

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBinaryCommand(t *testing.T) {
	actual := NewBinaryCommand("@bin")
	assert.Equal(t, "@bin", actual.Use)
	assert.Equal(t, "Run @bin with his injected password", actual.Short)
}
