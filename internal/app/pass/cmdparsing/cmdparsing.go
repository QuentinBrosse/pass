package cmdparsing

import (
	"flag"
	"log"
	"os/user"
	"path/filepath"
)

// Holds all arguments passed through the command line
type CmdArgs struct {
	ConfDirPath string // Default directory for all configuration/user files
}

// Returns the default directory path for all configuration/user files based on the current user (~/.pass)
func getDefaultConfDirPath() string {
	currentUser, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	return filepath.Join(currentUser.HomeDir, ".pass")
}

// Setups cmd flags parsing and returns the linked CmdArgs struct where results will be stored
func Construct() *CmdArgs {
	cmdArgs := new(CmdArgs)
	defaultConfDirPath := getDefaultConfDirPath()

	flag.StringVar(&cmdArgs.ConfDirPath, "conf-dir", defaultConfDirPath, "Directory for all configuration/user files")

	return cmdArgs
}
