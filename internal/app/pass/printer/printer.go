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
