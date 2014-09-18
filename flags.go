package main

import (
	"github.com/codegangsta/cli"
)

func QuietFlag() cli.BoolFlag {
	return cli.BoolFlag{
		Name:  "quiet, q",
		Usage: "only display results",
	}
}

func RegistryFlag() cli.StringFlag {
	return cli.StringFlag{
		Name:   "registry-uri, r",
		Usage:  "specify your registry uri",
		EnvVar: "REGISTRY_URI",
	}
}
