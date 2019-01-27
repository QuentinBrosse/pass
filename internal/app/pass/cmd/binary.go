package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Create new binary command from his name
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
