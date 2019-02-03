package printer

import (
	"fmt"
	"os"
)

// PrintlnErr runs fmt.Println and handle the error.
func PrintlnErr(a ...interface{}) {
	_, err := fmt.Fprintln(os.Stderr, a...)
	if err != nil {
		panic(err)
	}
}

// PrintfErr runs fmt.PrintfErr and handle the error.
func PrintfErr(format string, a ...interface{}) {
	_, err := fmt.Fprintf(os.Stderr, format, a...)
	if err != nil {
		panic(err)
	}
}

// PrintlnErrExit runs PrintlnErr and exit the program with code 1.
func PrintlnErrExit(a ...interface{}) {
	PrintlnErr(a...)
	os.Exit(1)
}

// PrintfErrExit runs PrintfErr and exit the program with code 1.
func PrintfErrExit(format string, a ...interface{}) {
	PrintfErr(format, a...)
	os.Exit(1)
}
