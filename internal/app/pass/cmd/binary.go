package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

// Create new binary command from his name.
func NewBinaryCommand(binary string) *cobra.Command {
	return &cobra.Command{
		Use:   binary,
		Short: "Run " + binary + " with its injected password",
		Run:   runBinary,
	}
}

func runBinary(cmd *cobra.Command, args []string) {
	// TODO: Implement me
	fmt.Println("Run:", cmd.Use)
}

// Find the binary name in provided binary.
func findBinaryNameInArgs(args ...string) string {
	for _, arg := range args {
		if strings.HasPrefix(arg, defaultBinaryPrefix) && len(arg) > len(defaultBinaryPrefix) {
			return arg
		}
	}

	return ""
}

// Create binary command if it exists in arguments.
func createBinaryCmdFromArgs(args ...string) *cobra.Command {
	binaryName := findBinaryNameInArgs(args...)
	if binaryName == "" {
		return nil
	}

	return NewBinaryCommand(binaryName)
}
