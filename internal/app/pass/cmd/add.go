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

func runAdd(cmd *cobra.Command, args []string) {
	// TODO: implement me
	fmt.Println("Run:", cmd.Use)
}
