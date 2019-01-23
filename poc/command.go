package poc

import (
	"os"
	"os/exec"
)

func NewCommand(binary string, args []string) *exec.Cmd {
	cmd := exec.Command(binary, args...)
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	return cmd
}
