package cmd

import (
	"os/user"
	"path/filepath"

	"github.com/QuentinBrosse/pass/internal/app/pass/onboarding"
	"github.com/QuentinBrosse/pass/internal/app/pass/printer"

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
		if err := cmd.Usage(); err != nil {
			panic("cannot print usage: " + err.Error())
		}
	},
}

var PersistentFlags = new(PersistentFlagsVars)

// Holds all arguments passed through the command line
type PersistentFlagsVars struct {
	ConfDirPath string
}

// Build the command
func Build(args ...string) {
	cobra.OnInitialize(initialization)

	Pass.SetUsageTemplate(helpTemplate)

	flagSet := Pass.PersistentFlags()
	// TODO: Implement it
	flagSet.StringVar(&PersistentFlags.ConfDirPath, "conf-dir", getDefaultConfDirPath(), "Directory for all configuration/user files")

	Pass.AddCommand(
		add,
		edit,
		delete,
	)

	if binaryCmd := createBinaryCmdFromArgs(args...); binaryCmd != nil {
		Pass.AddCommand(binaryCmd)
	}
}

func initialization() {
	err := onboarding.Run()
	if err != nil {
		printer.PrintlnErrExit(err)
	}
}

// Execute the pass command.
func Execute() {
	if err := Pass.Execute(); err != nil {
		printer.PrintlnErrExit(err)
	}
}

// Returns the default directory path for all configuration/user files based on the current user (~/.pass).
func getDefaultConfDirPath() string {
	currentUser, err := user.Current()
	if err != nil {
		panic("cannot get the user home directory")
	}

	return filepath.Join(currentUser.HomeDir, defaultConfigDirName)
}
