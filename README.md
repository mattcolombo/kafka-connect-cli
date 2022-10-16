# kafka-connect-cli

A comprehensive command line tool to manage Kafka connectors and some general configurations in a Connect installation; main command `kconnect-cli`. The CLI leverages the Connect Rest API, creating a client and issuing requests to said API. It allows to work with variously configured Connect clusters (including using mTLS, basic and token/API-key based authentication). 

## Configuration requirements

In order to be able to connect with the correct Kafka Connect cluster, a configuration file is required. This configuration file (a sample of which can be found [here](/samples/kconnect-cli-config.yaml.tmpl)) can be either added to an environment variable called `CONNECTCFG` or located in the folder where the CLI is invoked from. Notice that in the latter case the name **must** be `kconnect-cli-config.yaml`.

A complete guide to the configuration is found [here](/docs/CONFIGURATION.md).

## Top level commands and resources

The main commands to manage a connect cluster resource are as follow:

* `cluster` allows to gather information on the connect cluster. Documentation on the usage of this resource can be found [here](/docs/CLUSTER.md);
* `connector` allows to gather information, create and manage connectors. Documentation on the usage of this resource can be found [here](/docs/CONNECTOR.md);
* `task` allows to gather information and manage connector tasks. Documentation on the usage of this resource can be found [here](/docs/TASK.md);
* `logger` allows to gather information and manage loggers and log levels. Documentation on the usage of this resource can be found [here](/docs/LOGGER.md);
* `version` returns the current version of the CLI. If the `--show-cli-config` flag is added, will also provide information on the location of the currently used configuration file and the main URL and protocol of the Connect cluster in use;
* `help` provides help on the usage of the CLI. It can be used as a command or through the `--help` flag (shorthand `-h`).

## Quickstart

TODO

* Build it yourself
* Using docker image
* Using k8s and the provided image

Further information on installation or build of the CLI can be found in the [installation guide](/docs/INSTALLATION-GUIDE.md); a set of samples/utility files and resources can be found in the [samples](/samples/) folder.

## Further reading

A great guide on how the Connect API works in general, and all the endpoints it offers can be found for example in the [Confluent official documentation](https://docs.confluent.io/platform/current/connect/references/restapi.html).