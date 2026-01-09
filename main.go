package main

import (
	"fmt"
	"log"
	"os"

	"github.com/armadi1809/vaulta/vault"
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
		err := vault.InitVault("./vault.json")
		if err != nil {
			log.Fatalf("Failed to initialize the vault.. %v", err)
		}
	case "add":
		// add entry workflow
	case "replace":
		// replace entry workflow
	case "delete":
		// delete entry workflow
	}

}
