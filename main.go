package main

import (
	"fmt"
	"os"
)

const usage string = `Usage: vaulta <command> <parameters>
Available Commands: 
- init: Initialize a vault if none exists on your machine`

func main() {
	args := os.Args
	if len(args) <= 1 {
		fmt.Println(usage)
		os.Exit(-1)
	}
	switch args[1] {
	case "init":
		// init vault workflow
	case "add":
		// add entry workflow
	case "replace":
		// replace entry workflow
	case "delete":
		// delete entry workflow
	}

}
