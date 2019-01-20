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

func PrintPasswordAndExist(name, password string) {
	fmt.Println(name+":", password)
	fmt.Println("Echo:", echo)
	os.Exit(0)
}

func main() {

	// Try with value flag
	flagPassword := flag.String("pv", "", "the password in clear")
	flag.Parse()

	if *flagPassword != "" {
		PrintPasswordAndExist("Flag value password", *flagPassword)
	}

	// Try with varenv
	varenvPassword := os.Getenv("PASSWORD")
	if varenvPassword != "" {
		PrintPasswordAndExist("Varenv password", varenvPassword)
	}

	fmt.Println("No password found\n")
	flag.Usage()
	os.Exit(1)
}
