# Configuration Guide

For the CLI to work, some information about the Kafka Connect cluster to connect to must be provided. This is done through a configuration file in YAML format; the structure for the configuration file is explained in the ["Configuration file scructure"](#configuration-file-structure) section. Notice that the configuration file is the first thing that is loaded as the CLI is invoked, before any flags or arguments are parsed, so it cannot be provided as part of the command itself but must instead be provided elsewhere. There are two options for this, which are deswcribed below in the ["Locating the configuration file"](#locating-the-configuration-file) section.

## Locating the configuration file

There are two options to load the correct configuration file with the instructions for contacting the Kafka Connect installation. 

The first option is to set an environment variable called `CONNECTCFG` with the path to the configuration file. Using the complete path is normally better to avoid confusion (particularly if the CLI executable location is added to the system path and can therefore be reached from anywhere) but any path will in principle work. Also not that unlike the second option, since we are providing the complete path to the file, any name for the configuration file will work fine.

If an environment variable as per above is not set, then the CLI will default to looking for a file called `kconnect-cli-config.yaml` in the local directory where the command is executed from. If a file cannot be found in either case, the execution will throw an error. **NOTE** here current folder refers to the folder where the user is located when running the commands, not to the folder where the CLI executable is located (in case such location is added to the PATH, and used from elsewhere).

Environment variables can be created using the following commands.

For Windows systems:
```(powershell)
$env:CONNECTCFG = '<Drive>:\path\to\config-file.yaml'
```

For Linux systems:
```(shell)
export CONNECTCFG=/path/to/config-file.yaml
```

## Working with multiple configuration files

In the case we need to work with multiple Connect clusters, and therefore it is required to be able to switch between different configuration files, the best option is to store all the configuration files in a common directory by giving them clear names to specify which cluster they refer to and then use the environment variable `CONNECTCFG` to switch between them. This could also be done by setting beforehand multiple environment variables with the full path to each file, and then simply assigning the correct one to `CONNECTCFG` as required.

Alternatively, one could use different directories for each cluster and add in each one a `kconnect-cli-config.yaml` file with the correct information in it; switching would then happen by moving directories. This last option however requires all the configuration files to have the same name so it would make less clear which one is being used at this specific time.

### Using aliases for faster switching

When running the CLI on a Linux system, a quicker way to switch between configuration files is to create aliases to run the `export` command. Aliases are created using
```(shell)
alias kcli-config1='export CONNECTCFG=/path/to/config-file-1.yaml'
alias kcli-config2='export CONNECTCFG=/path/to/config-file-2.yaml'
```
and so on for however many config files one may need to use. Switching then can simply be done by running the alias (e.g. `kcli-config1` will set the `config-file-1.yaml` as current in the environment variable, and so on). 

--**NOTE**-- exports and aliases are only going to be persisted for the length of the session. Once a new session is started, they will not be available any longer. To make those permanent, simply add them to the `.bashrc` file or similar depending on the system in use. This way they will be loaded in as soon as a session is started and will therefore be always present.

## Configuration file structure

The configuration file contains all the information required to connect to the desired installation of Kafka Connect. This includes the hosts for Connect, the protocol to use for the connection and the authentication information (if required) for the connection. A sample configuration file can be found [here](/samples/kconnect-cli-config.yaml.tmpl). Simply make a copy, rename it as `kconnect-cli-config.yaml`, fill in the required information, set the environment variable to the path to this file (the fully qualified path is best as it will allow to run the CLI from any location in the system) and you are good to go.

* `hostnames` contains a list of all the hostnames for the workers in the Kafka Connect cluster. While most calls are going to be directed to the first one in the list, all of them are required (see note below)
* `protocol` can be HTTP or HTTPS; in most cases it will be HTTPS if mTLS is enabled, and in general should be HTTPS if any of the authentication methods are used, but it has been left as an independent configuration for specific cases where it may be required
* `tls` is required in the case where mTLS is enabled, and the paths to the client certificate, key and truststore needs to be specified
* `basicauth` allows to specify a username and password
* `tokenauth` allows to specify the token to be used while using `Bearer` in the `Authorization` header
* `apikeyauth` works the same as `tokenauth` but allows to specify the header to use (normally something like `x-api-key` or similar)

--**NOTE**-- in the `hostnames` field within the configuration file, it's very important to list all the hosts that make up the Connect cluster we wish to manage. While for the vast majority of API call we could use any of the hosts (and in fact here we will use only the first in the list), specifically setting and checking loggers and log levels need to be done in each of the hosts independently. For this reason the `logger` commands will loop over all the hosts. If a host is missing from the list log levels will not be set dynamically for that host.
