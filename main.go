package main

import (
	fmt "fmt"
	os "os"

	flags "github.com/jessevdk/go-flags"

	api "github.com/eolme/backmemes/api"
	cli "github.com/eolme/backmemes/cli"
	config "github.com/eolme/backmemes/config"
)

type Options struct {
	Mode string `long:"mode" description:"run mode (api, cli)"`
}

func main() {
	var opts Options
	parser := flags.NewParser(&opts, flags.Default)
	parser.Name = "backmemes"
	parser.Usage = "--mode=<mode>"

	_, err := flags.Parse(&opts)
	if err != nil || flags.WroteHelp(err) {
		return
	}

	if opts.Mode == "api" || opts.Mode == "server" {
		config.Connect()
		api.Server()
		return
	}

	if opts.Mode == "cli" || opts.Mode == "terminal" {
		config.Connect()
		cli.Terminal()
		return
	}

	fmt.Printf("Unsupported mode '%s'\r\n", opts.Mode)
	os.Exit(1)
}
