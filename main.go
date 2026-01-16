package main

import (
	"fmt"
	"os"

	"github.com/alecthomas/kong"
	"github.com/armadi1809/vaulta/config"
	"github.com/armadi1809/vaulta/ui"
	"github.com/armadi1809/vaulta/vault"
)

type Init struct {
}

type List struct {
}

type Add struct {
}

type Get struct {
	Entry string `arg:"" name:"entry" help:"Entry to get from the vault." type:"string"`
}

type Delete struct {
	Entry string `arg:"" name:"entry" help:"Entry to delete from the vault." type:"string"`
}

func (i *Init) Run(vault *vault.Vault) error {
	return vault.InitVault()
}

func (l *List) Run(vault *vault.Vault) error {
	res, err := vault.ListEntries()
	if err != nil {
		fmt.Println(ui.RenderError(fmt.Sprintf("Failed to list entries: %v", err)))
		os.Exit(1)
	}
	fmt.Println(res)
	return nil
}

func (a *Add) Run() error {
	err := vault.AddEntry()
	if err != nil {
		fmt.Println(ui.RenderError(fmt.Sprintf("Failed to add entry: %v", err)))
		os.Exit(1)
	}
	return nil
}

func (g *Get) Run(vault *vault.Vault) error {
	res, err := vault.GetEntry(g.Entry)
	if err != nil {
		fmt.Println(ui.RenderError(fmt.Sprintf("Failed to get entry: %v", err)))
		os.Exit(1)
	}
	fmt.Println(res)
	return nil
}

func (d *Delete) Run(vault *vault.Vault) error {
	err := vault.DeleteEntry(d.Entry)
	if err != nil {
		fmt.Println(ui.RenderError(fmt.Sprintf("Failed to delete entry: %v", err)))
		os.Exit(1)
	}
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
	ctx := kong.Parse(&cli,
		kong.Name("vaulta"),
		kong.Description(ui.SubtitleStyle.Render("🔐 A secure password vault for the command line")),
		kong.UsageOnError(),
	)
	vaultPath, err := config.DefaultVaultPath()
	if err != nil {
		fmt.Println(ui.RenderError(fmt.Sprintf("Failed to get vault path: %v", err)))
		os.Exit(1)
	}
	v, err := vault.New(vaultPath)
	if err != nil {
		fmt.Println(ui.RenderError(fmt.Sprintf("Failed to initialize vault: %v", err)))
		os.Exit(1)
	}
	err = ctx.Run(v)
	ctx.FatalIfErrorf(err)
}
