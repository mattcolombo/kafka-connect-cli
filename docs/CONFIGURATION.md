# Configuration Guide

For the CLI to work, some information about the Kafka Connect cluster to connect to must be provided. This is done through a configuration file in YAML format; the structure for the configuration file is explained in the (#configuration-file-structure) section. Notice that the configuration file is the first thing that is loaded as the CLI is invoked, so it cannot be provided as part of the command itself, but must be provided differently. There are two options for this, which are deswcribed below in the (#locating-the-configuration-file) section.

## Locating the configuration file

## Working with multiple configuration files (Linux only)

## Configuration file structure






## Configuration requirements

The hosts where Connect is located at, along with some other options and information, need to be stored in a configuration file (in YAML format). Since the configuration file is loaded as soon as the CLI is called and before any flags and commands are parsed, the location of such configuration file need to be stored in an environment variable named `CONNECTCFG`. If such environment variable is not set, the CLI will look for a file called `./connect-config.yaml` in the current folder. **NOTE** the current folder refers to the folder where the user is located running the commands, not to the folder where the CLI executable is located (in case such location being added to the PATH, and used from elsewhere). If such file is not present ant the environment variable is not set, an error will be thrown.

A sample configuration file can be found [here](samples/connect-config.yaml.tmpl). Simply make a copy, rename it as `connect-config.yaml`, fill in the required information, set the environment variable to the path to this file (the fully qualified path is best as it will allow to run the CLI from any location in the system) and you are good to go.

--**NOTE**-- in the `hostnames` field within the configuration file, it's very important to list all the hosts that make up the Connect cluster we wish to manage. While for the vast majority of API call we could use any of the hosts (and in fact here we will use only the first in the list), specifically setting and checking loggers and log levels need to be done in each of the hosts independently. For this reason the `logger` commands will loop over all the hosts. If a host is missing from the list log levels will not be set dynamically for that host.

More info [here](docs/CONFIGURATION.md).