package utilities

import (
	"encoding/json"
	"fmt"
	"os"
)

var ConnectConfiguration Configuration = ImportConfig()

func ImportConfig() Configuration {
	fmt.Println("I am importing the configuration file") // control statement print - TOREMOVE
	file, err := os.Open(os.Getenv("CONNECTCFG"))        // previously used hardcoded ./connect-config.json
	if err != nil {
		fmt.Println("Please add the configuration file as an environment variable named CONNECTCFG")
		fmt.Println(err)
		os.Exit(1)
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
