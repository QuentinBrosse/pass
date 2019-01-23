package main

import (
	"flag"
	"fmt"

	"github.com/QuentinBrosse/pass/internal/app/pass/cmdparsing"
)

var cmdArgs *cmdparsing.CmdArgs

func init() {
	cmdArgs = cmdparsing.Construct()
}

func main() {
	flag.Parse()

	fmt.Println("ConfDirPath:", cmdArgs.ConfDirPath)
}
