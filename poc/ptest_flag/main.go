package main

import (
	"flag"
	"fmt"
	"os"
)

var echo string

func init() {
	flag.StringVar(&echo, "echo", "", "echo string")
}

func PrintPasswordAndExist(password string) {
	fmt.Println("Flag password:", password)
	fmt.Println("Echo:", echo)
	os.Exit(0)
}

func main() {
	flagPassword := flag.String("pv", "", "the password in clear")
	flag.Parse()

	if *flagPassword != "" {
		PrintPasswordAndExist(*flagPassword)
	}

	fmt.Println("No password found\n")
	flag.Usage()
	os.Exit(1)
}
