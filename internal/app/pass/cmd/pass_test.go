package cmd

import (
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
)

func Test_DefaultConfDirPath(t *testing.T) {
	expected := getDefaultConfDirPath()
	Build()
	actual := PersistentFlags.ConfDirPath

	assert.Equal(t, expected, actual)
}

func Test_createBinaryCommand(t *testing.T) {
	for _, c := range []struct {
		args     []string
		expected *cobra.Command
	}{
		{
			args:     []string{},
			expected: nil,
		},
		{
			args:     []string{"add", "--name", "test"},
			expected: nil,
		},
		{
			args:     []string{"@bin", "-v"},
			expected: NewBinaryCommand("@bin"),
		},
	} {
		actual := createBinaryCmdFromArgs(c.args...)
		if c.expected == nil {
			assert.Nil(t, nil)
			continue
		}
		assert.Equal(t, c.expected.Use, actual.Use)
		assert.Equal(t, c.expected.Short, actual.Short)
	}
}

func Test_findBinaryNameInArgs(t *testing.T) {
	for _, c := range []struct {
		args     []string
		expected string
	}{
		{
			args:     []string{},
			expected: "",
		},
		{
			args:     []string{"add", "--name", "test"},
			expected: "",
		},
		{
			args:     []string{"@bin", "-v"},
			expected: "@bin",
		},
		{
			args:     []string{"bin@bin", "-v"},
			expected: "",
		},
		{
			args:     []string{"@", "-v"},
			expected: "",
		},
	} {
		actual := findBinaryNameInArgs(c.args...)
		assert.Equal(t, c.expected, actual)
	}
}
