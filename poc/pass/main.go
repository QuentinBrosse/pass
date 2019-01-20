package main

import (
	"fmt"
	"os"
	"os/exec"
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
	plugin, err := poc.NewPluginFromConfig(binary)
	if err != nil {
		ExitWithError(err)
	}

	if err != nil {
		ExitWithError(err)
	}

	// Create command
	cmd := exec.Command(binary, args...)
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	// Set password to plugin
	plugin.SetPassword(ThePassword)

	// Prepare plugin context
	err = plugin.Prepare()
	if err != nil {
		ExitWithError(err)
	}

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
