package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var edit = &cobra.Command{
	Use:   "edit",
	Short: "Edit a new entry",
	Run:   runEdit,
}

func runEdit(cmd *cobra.Command, args []string) {
	// TODO: Implement me
	fmt.Println("Run:", cmd.Use)
}
