package main

import (
	"fmt"
	"os"
	"os/exec"
)

const ThePassword = "🤫🤫🤫🤫"

var (
	CmdRequired        = fmt.Errorf("a command is required")
	BinaryNotSupported = fmt.Errorf("binary not supported")
)

const (
	MethodType_Unknow = iota

	// The password is passed in clear in value of the flag.
	// eg. -p {password}
	MethodType_FlagValue

	// The password is in a file and the flag value is the file path.
	// eg. -p /etc/var/{password}
	MethodType_FlagPath
)

type Plugin struct {
	BinaryName string
	MethodType uint
	FlagName   string
}

var Plugins = map[string]*Plugin{
	"ptest": {
		BinaryName: "ptest",
		MethodType: MethodType_FlagValue,
		FlagName:   "-pv",
	},
}

func ExitWithError(err error) {
	fmt.Fprintln(os.Stderr, "error:", err)
	os.Exit(1)
}

func GetPlugin(binary string) *Plugin {
	plugin, ok := Plugins[binary]
	if ok {
		return plugin
	}
	return nil
}

func InjectPassword(cmd *exec.Cmd, plugin *Plugin) error {
	switch plugin.MethodType {
	case MethodType_FlagValue:
		begin := []string{cmd.Args[0], plugin.FlagName + "=" + ThePassword}
		cmd.Args = append(begin, cmd.Args[1:]...)
		return nil
	case MethodType_FlagPath:
		return fmt.Errorf("method type flag not implemented wet")
	default:
		return fmt.Errorf("wrong method type flag for %s plugin", plugin.BinaryName)
	}
}

func main() {
	if len(os.Args) < 2 {
		ExitWithError(CmdRequired)
	}

	// Parse arguments
	binary := os.Args[1]
	args := os.Args[2:]

	// Find plugin
	plugin := GetPlugin(binary)
	if plugin == nil {
		ExitWithError(BinaryNotSupported)
	}

	// Prepare command
	cmd := exec.Command(binary, args...)
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	err := InjectPassword(cmd, plugin)
	if err != nil {
		ExitWithError(BinaryNotSupported)
	}

	// Run !
	err = cmd.Run()
	if err != nil {
		ExitWithError(err)
	}
}
