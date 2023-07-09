package utilities

import (
	"errors"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// architectural decision: use one yaml config file per environment, and that needs to be set in the default location or by ENV variable
var defaultLocation = "./kconnect-cli-config.yaml"
var ConfigLoadPath = FindConfig()
var ConnectConfiguration = ImportConfig(ConfigLoadPath)

func FindConfig() string {
	path, isSet := os.LookupEnv("CONNECTCFG")
	if !isSet {
		path = defaultLocation
	}
	return path
}

func ImportConfig(path string) ConfigurationYaml {
	//fmt.Println("I am importing the configuration file from", path) // control statement print
	file, err := os.ReadFile(path)
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

	configuration := ConfigurationYaml{}
	err = yaml.Unmarshal(file, &configuration)
	if err != nil {
		fmt.Printf("Error parsing YAML file: %s\n", err)
		os.Exit(1)
	}

	return configuration
}
