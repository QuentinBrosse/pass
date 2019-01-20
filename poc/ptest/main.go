package main

import (
	"flag"
	"fmt"
	"io/ioutil"
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
	flagValuePassword := flag.String("pv", "", "the password value in clear")
	flagPath := flag.String("pp", "", "the path of the file containing the password")
	flag.Parse()

	varenvPassword := os.Getenv("PASSWORD")

	// Try with flag value
	if *flagValuePassword != "" {
		PrintPasswordAndExist("Flag value password", *flagValuePassword)
	}

	// Try with flag path
	if *flagPath != "" {
		raw, err := ioutil.ReadFile(*flagPath)
		if err != nil {
			panic(err)
		}
		PrintPasswordAndExist("Flag path password", string(raw))
	}

	// Try with varenv
	if varenvPassword != "" {
		PrintPasswordAndExist("Varenv password", varenvPassword)
	}

	fmt.Println("No password found\n")
	flag.Usage()
	os.Exit(1)
}
