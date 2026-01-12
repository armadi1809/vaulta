package main

import (
	"fmt"
	"log"
	"os"

	"github.com/armadi1809/vaulta/vault"
)

const usage string = `Usage: vaulta <command> <parameters>
Available Commands: 
- init: Initialize a vault if none exists on your machine
- add: Add an entry to the vault
- list: List entries available in the vault
- get <entry>: Get information for the specified entry from the vault
- delete <entry>: Delete information for the specified entry form the vault`

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
		err := vault.AddEntry("./vault.json")
		if err != nil {
			log.Fatalf("Failed to add entry to the vault.. %v", err)
		}
	case "get":
		// get entry workflow
		res, err := vault.GetEntry("./vault.json")
		if err != nil {
			log.Fatalf("Failed to get entry from the vault.. %v", err)
		}
		fmt.Println(res)
	case "list":
		res, err := vault.ListEntries("./vault.json")
		if err != nil {
			log.Fatalf("Failed to list entries from the vault.. %v", err)
		}
		fmt.Println(res)
	case "replace":
		// replace entry workflow
	case "delete":
		if len(args) != 3 {
			fmt.Println(usage)
			os.Exit(-1)
		}
		err := vault.DeleteEtnry("./vault.json", args[2])
		if err != nil {
			fmt.Printf("an error ocurred while deleting entry... %v", err)
			os.Exit(-1)
		}
		fmt.Println("Entry succesfully deleted")
	}

}
