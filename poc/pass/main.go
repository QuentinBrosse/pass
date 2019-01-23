package main

import (
	"fmt"
	"os"
	"pass/poc"
)

const ThePassword = "🤫🤫🤫🤫"

func ExitWithError(err error) {
	fmt.Fprintln(os.Stderr, "error:", err)
	os.Exit(1)
}

func main() {
	if len(os.Args) < 2 {
		ExitWithError(fmt.Errorf("a command is required"))
	}

	// Parse arguments
	binary := os.Args[1]
	args := os.Args[2:]

	// Find plugin
	plugin, err := poc.NewPlugin(binary)
	if err != nil {
		ExitWithError(err)
	}

	// Prepare plugin context
	err = plugin.Prepare()
	if err != nil {
		ExitWithError(err)
	}

	// Set password to plugin
	plugin.SetPassword(ThePassword)

	// Create command
	cmd := poc.NewCommand(binary, args)

	// Inject password
	err = plugin.InjectPassword(cmd)
	if err != nil {
		ExitWithError(err)
	}

	// Run command
	err = cmd.Run()
	if err != nil {
		ExitWithError(err)
	}

	// Clean up plugin context
	err = plugin.CleanUp()
	if err != nil {
		ExitWithError(err)
	}
}
