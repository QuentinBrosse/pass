package cmdparsing

import "testing"

func TestConstruct(t *testing.T) {
	expected := &CmdArgs{getDefaultConfDirPath()}
	result := Construct()

	if result.ConfDirPath != expected.ConfDirPath {
		t.Errorf("Default configuration directory path was incorrect, expected \"%s\", received \"%s\"", expected.ConfDirPath, result.ConfDirPath)
	}
}
