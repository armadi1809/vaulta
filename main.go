package main

import (
	"fmt"

	"github.com/alecthomas/kong"
	"github.com/armadi1809/vaulta/vault"
)

const vaultPath string = "./vault.json"

type Init struct {
}

type List struct {
}

type Add struct {
}

type Get struct {
}

type Delete struct {
	Entry string `arg:"" name:"entry" help:"Entry to delete from the vault." type:"string"`
}

func (i *Init) Run() error {
	return vault.InitVault(vaultPath)
}

func (l *List) Run() error {
	res, err := vault.ListEntries("./vault.json")
	if err != nil {
		return fmt.Errorf("Failed to list entries from the vault.. %v", err)
	}
	fmt.Println(res)
	return nil
}

func (a *Add) Run() error {
	err := vault.AddEntry(vaultPath)
	if err != nil {
		return fmt.Errorf("Failed to add entry to the vault.. %v", err)
	}
	return nil
}

func (g *Get) Run() error {
	res, err := vault.GetEntry(vaultPath)
	if err != nil {
		return fmt.Errorf("Failed to get entry from the vault.. %v", err)
	}
	fmt.Println(res)
	return nil
}

func (d *Delete) Run() error {
	err := vault.DeleteEtnry(vaultPath, d.Entry)
	if err != nil {
		return fmt.Errorf("an error ocurred while deleting entry... %v", err)
	}
	fmt.Println("Entry succesfully deleted")
	return nil
}

var cli struct {
	Init   Init   `cmd:"" help:"Initialize the vault."`
	List   List   `cmd:"" help:"List entries in the vault."`
	Get    Get    `cmd:"" help:"Get an entry in the vault."`
	Add    Add    `cmd:"" help:"Add an entry to the vault."`
	Delete Delete `cmd:"" help:"Delete an entry from the vault."`
}

func main() {
	ctx := kong.Parse(&cli)
	err := ctx.Run()
	ctx.FatalIfErrorf(err)
}
