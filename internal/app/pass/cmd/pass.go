package cmd

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)

const defaultConfigDirName = ".pass"

// TODO: Make it customizable (in conf file)
const defaultBinaryPrefix = "@"

var Pass = &cobra.Command{
	Use: "pass",

	// TODO: Fill the short
	Short: "🗝 CLI password manager",

	// TODO: Fill the long
	Long: ``,

	Run: func(cmd *cobra.Command, args []string) {
		cmd.Usage()
	},
}

var PersistentFlags = new(PersistentFlagsVars)

// Holds all arguments passed through the command line
type PersistentFlagsVars struct {
	ConfDirPath string
}

// Build the command
func Build(args ...string) {
	Pass.SetUsageTemplate(helpTemplate)

	flagSet := Pass.PersistentFlags()
	// TODO: Implement it
	flagSet.StringVar(&PersistentFlags.ConfDirPath, "conf-dir", getDefaultConfDirPath(), "Directory for all configuration/user files")

	Pass.AddCommand(
		add,
		edit,
		delete,
	)

	if binaryCmd := createBinaryCommand(args...); binaryCmd != nil {
		Pass.AddCommand(binaryCmd)
	}
}

// Execute the pass command
func Execute() {
	if err := Pass.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// Create binary command if it exists in arguments
func createBinaryCommand(args ...string) *cobra.Command {
	binaryName := findBinaryNameInArgs(args...)
	if binaryName == "" {
		return nil
	}

	return NewBinaryCommand(binaryName)
}

// Find the binary name in provided binary
func findBinaryNameInArgs(args ...string) string {
	for _, arg := range args {
		if strings.HasPrefix(arg, defaultBinaryPrefix) && len(arg) > len(defaultBinaryPrefix) {
			return arg
		}
	}

	return ""
}

// Returns the default directory path for all configuration/user files based on the current user (~/.pass)
func getDefaultConfDirPath() string {
	currentUser, err := user.Current()
	if err != nil {
		panic("cannot get the user home directory")
	}

	return filepath.Join(currentUser.HomeDir, defaultConfigDirName)
}
