package main

import (
	"fmt"
	"os"

	"github.com/QuentinBrosse/pass/internal/app/pass/cmd"
)

func globalRecover() {
	if r := recover(); r != nil {
		if _, err := fmt.Fprintln(os.Stderr, "error:", r); err != nil {
			fmt.Println("error:", r)
		}
		os.Exit(1)
	}
}

func main() {
	defer globalRecover()

	cmd.Build(os.Args...)
	cmd.Execute()

	fmt.Println("\nPersistentFlags:", cmd.PersistentFlags)
}
