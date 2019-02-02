package printer

import (
	"fmt"
	"os"
)

func PrintlnErr(a ...interface{}) {
	_, err := fmt.Fprintln(os.Stderr, a...)
	if err != nil {
		panic(err)
	}
}

func PrintfErr(format string, a ...interface{}) {
	_, err := fmt.Fprintf(os.Stderr, format, a...)
	if err != nil {
		panic(err)
	}
}

func PrintlnErrExit(a ...interface{}) {
	PrintlnErr(a...)
	os.Exit(1)
}

func PrintfErrExit(format string, a ...interface{}) {
	PrintfErr(format, a...)
	os.Exit(1)
}
