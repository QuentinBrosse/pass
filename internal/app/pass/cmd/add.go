package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var add = &cobra.Command{
	Use:   "add",
	Short: "Add a new entry",
	Run:   runAdd,
}

// RunAdd runs the add command.
func runAdd(cmd *cobra.Command, args []string) {
	// TODO: Implement me
	fmt.Println("Run:", cmd.Use)
}
