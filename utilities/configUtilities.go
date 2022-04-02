package utilities

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

// TODO find out how this can be done using yaml instead and implement it
// architectural decision: use one yaml config file per environment, and that needs to be set in the default location or by ENV variable
// add in the documentation aliases for Linux to load different environments

var ConnectConfiguration Configuration = ImportConfig()
var defaultLocation string = "./connect-config.json"

func ImportConfig() Configuration {
	path, isSet := os.LookupEnv("CONNECTCFG")

	if !isSet {
		path = defaultLocation
	}

	fmt.Println("I am importing the configuration file from", path) // control statement print - TOREMOVE
	file, err := os.Open(path)

	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("Config file not found! Please add the path to the configuration file as an environment variable named CONNECTCFG")
			os.Exit(1)
		} else {
			fmt.Println("Error while opening the configuration file")
			fmt.Println(err)
			os.Exit(1)
		}
	}

	defer file.Close()

	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err = decoder.Decode(&configuration)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return configuration
}
