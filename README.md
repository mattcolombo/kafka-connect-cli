# kafka-connect-cli

A comprehensive command line tool to manage a Kafka Connect installation, invoked by using `kconnect-cli`. The CLI leverages the Connect API, effectively creating a client under the cover and issuing requests to said API. It will allow to work with variously configured Connect clusters (including using mTLS, basic and token/API key based authentication). 

## Configuration requirements

The hosts where Connect is located at, along with some other options and information, need to be stored in a configuration file (in YAML format). Since the configuration file is loaded as soon as the CLI is called and before any flags and commands are parsed, the location of such configuration file need to be stored in an environment variable named `CONNECTCFG`. If such environment variable is not set, the CLI will look for a file called `./connect-config.yaml` in the current folder. **NOTE** the current folder refers to the folder where the user is located running the commands, not to the folder where the CLI executable is located (in case such location being added to the PATH, and used from elsewhere). If such file is not present ant the environment variable is not set, an error will be thrown.

A sample configuration file can be found [here](./samples/connect-config.yaml.tmpl). Simply make a copy, rename it as `connect-config.yaml`, fill in the required information, set the environment variable to the path to this file (the fully qualified path is best as it will allow to run the CLI from any location in the system) and you are good to go.

--**NOTE**-- In the `hostnames` field within the configuration file, it's very important to list all the hosts that make up the Connect cluster we wish to manage. While for the vast majority of API call we could use any of the hosts (and in fact here we will use only the first in the list), specifically setting and checking loggers and log levels need to be done in each of the hosts independently. For this reason the `logger` commands will loop over all the hosts. If a host is missing from the list log levels will not be set dynamically for that host.

More info [here](docs/CONFIGURATION.md).

## Top level commands and resources

The main commands to manage a connect cluster resource are as follow:

* `cluster` allows to gather information on the connect cluster. Documentation on the usage of this resource can be found [here](docs/CLUSTER.md);
* `connector` allows to gather information, create and manage connectors. Documentation on the usage of this resource can be found [here](docs/CONNECTOR.md);
* `task` allows to gather information and manage connector tasks. Documentation on the usage of this resource can be found [here](docs/TASK.md);
* `logger` allows to gather information and manage loggers and log levels. Documentation on the usage of this resource can be found [here](docs/LOGGER.md);
* `version` returns the current version of the CLI.
* `help` provides help on the usage of the CLI. It can be used as a command or more commonly through the `--help` flag (or alternatively the shorthand `-h`).

## Further reading

A great guide on how the Connect API works in general, and all the endpoints it offers can be found for example in the [Confluent official documentation](https://docs.confluent.io/platform/current/connect/references/restapi.html).