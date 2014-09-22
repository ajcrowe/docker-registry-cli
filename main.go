package main

import (
	"encoding/json"
	"github.com/codegangsta/cli"
	"io/ioutil"
	"log"
	"os"
)

var (
	VERSION           = "0.0.1"
	INDEX_API_VERSION = "v1"
	INDEX_URL         string
	config            *Config
)

type Config struct {
	RegistryURI string `json:"registry_uri"`
}

func main() {

	INDEX_URL = getIndexURI()

	app := cli.NewApp()
	// configure app
	app.Name = "drcli"
	app.Usage = "Manage you docker registry images & tags"
	app.Version = VERSION
	app.Author = "Alex Crowe"
	// register commands with app
	app.Commands = Commands
	// run app and pass args
	app.Run(os.Args)
}

// getIndexURI reads .drcli file and parses it into the Config struct
// if this file doesn't exist it will attempt to read the REGISTRY_URI
// environment variable.
func getIndexURI() string {
	text, err := ioutil.ReadFile(os.Getenv("HOME") + "/.drcli")
	if err != nil {
		env := os.Getenv("REGISTRY_URI")
		if env == "" {
			log.Println("Could not set registry uri!")
			os.Exit(1)
		}
		return env
	}
	json.Unmarshal(text, &config)
	return config.RegistryURI
}
