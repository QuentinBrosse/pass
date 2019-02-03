package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var delete = &cobra.Command{
	Use:   "delete",
	Short: "Delete a new entry",
	Run:   runDelete,
}

// RunDelete runs the delete command.
func runDelete(cmd *cobra.Command, args []string) {
	// TODO: Implement me
	fmt.Println("Run:", cmd.Use)
}
