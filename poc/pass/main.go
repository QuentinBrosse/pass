package main

import (
	"fmt"
	"os"
	"os/exec"
	"pass/poc"
	"path"
	"path/filepath"
)

const ThePassword = "🤫🤫🤫🤫"

const PluginDirName = "plugins"

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
	ex, err := os.Executable()
	if err != nil {
		ExitWithError(err)
	}
	exPath := filepath.Dir(ex)
	filePath := path.Join(exPath, PluginDirName, binary+".yml")
	file, err := os.Open(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			ExitWithError(fmt.Errorf("binary not supported"))
		}
		ExitWithError(fmt.Errorf("fail to open plugin file %s: %s", file.Name(), err))
	}

	plugin, err := poc.NewPluginFromConfig(binary, file)
	if err != nil {
		ExitWithError(err)
	}

	if err != nil {
		ExitWithError(err)
	}

	_ = file.Close()

	// Prepare command
	cmd := exec.Command(binary, args...)
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	err = plugin.Prepare()
	if err != nil {
		ExitWithError(err)
	}

	err = plugin.InjectPassword(cmd, ThePassword)
	if err != nil {
		ExitWithError(err)
	}

	fmt.Println(cmd.Args)

	err = plugin.CleanUp()
	if err != nil {
		ExitWithError(err)
	}

	// Run !
	err = cmd.Run()
	if err != nil {
		ExitWithError(err)
	}
}
