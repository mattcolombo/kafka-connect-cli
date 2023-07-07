# kafka-connect-cli

A comprehensive command line tool to manage Kafka connectors and some general configurations in a Connect installation; main command `kconnect-cli`. The CLI leverages the Connect Rest API, creating a client and issuing requests to said API. It allows to work with variously configured Connect clusters (including using mTLS, basic and token/API-key based authentication). 

## Configuration requirements

In order to be able to connect with the correct Kafka Connect cluster, a configuration file is required. This configuration file (a sample of which can be found [here](/samples-templates/kconnect-cli-config-template.yaml)) can be either added to an environment variable called `CONNECTCFG` or located in the folder where the CLI is invoked from. Notice that in the latter case the name **must** be `kconnect-cli-config.yaml`.

A complete guide to the configuration is found [here](/docs/CONFIGURATION.md).

## Top level commands and resources

The main commands to manage a connect cluster resource are as follow:

* `cluster` allows to gather information on the connect cluster. Documentation on the usage of this resource can be found [here](/docs/CLUSTER.md);
* `connector` allows to gather information, create and manage connectors. Documentation on the usage of this resource can be found [here](/docs/CONNECTOR.md);
* `task` allows to gather information and manage connector tasks. Documentation on the usage of this resource can be found [here](/docs/TASK.md);
* `logger` allows to gather information and manage loggers and log levels. Documentation on the usage of this resource can be found [here](/docs/LOGGER.md);
* `version` returns the current version of the CLI. Using the flag `--json` (shorthand `-j`) will print the extended version in Json format;
* `get-config` shows the path that the active CLI configuration file is being taken from, and prints out the full config file;
* `help` provides help on the usage of the CLI. It can be used as a command or through the `--help` flag (shorthand `-h`).

## Quickstart

THe easiest and quickest way to try out the CLI is to run it in Docker. Note for this to work Docker needs to be able to connect to the instance of Kafka Connect we would like to work against. 

To do so, you will need to follow few simple steps:
1. prepare a configuration file with the details of the Kafka Connect instance you would like to manage using [this template](/samples-templates/kconnect-cli-config-template.yaml) as base
2. pull the image containing the CLI from Docker Hub using `docker pull mattcolombo/kafka-connect-cli:<tag>` (repository on Docker Hub can be found [here](https://hub.docker.com/r/mattcolombo/kafka-connect-cli))
3. run the image using the location where the configuration file created in step 1 is located using `docker run --rm -it --mount type=bind,source=<absolute-path-to-source-dir>,target=/usr/cli/config,readonly mattcolombo/kafka-connect-cli:<tag> bash` (see the full detail [here](/docs/INSTALLATION-GUIDE.md#running-the-cli-image-locally-in-docker))
4. create an environment variable `export CONNECTCFG=/usr/cli/config/<name-of-config.yaml>` 
5. enjoy

Further information on how to install or build the CLI can be found in the [installation guide](/docs/INSTALLATION-GUIDE.md); a set of samples/utility files and resources can be found in the [samples](/samples-templates/) folder.

## Feedback & Contributing

This CLI was created as a personal project aimed at learning more about various technologies (primarily Go, k8s, and Kafka Connect). It is in no way perfect and there may be bugs or unexpected behaviours. If you find any please feel very free to open an issue in the GitHub page of this project, or get in touch with me directly at `mattcolombo.dev@gmail.com` and I will be delighted to have a look. Bear in mind however that this is a side project I am doing on personal time, so I am not guaranteeing any timelines for such fixes (though I will strive to fix simple things timely).

If you would like to contribute, or if you would like to modify and adapt this CLI for your own specific requirements, please feel free to fork the project and if you feel so, open a pull request later on. It's very important to me that this project remains open source and free for as many people as possible to modify, play with and generally do as they wish, so any feedback or contribution is welcome.

## Planned improvements and future work

Few things I would like to do in the future (when time allows)

- the output of the `validate` command is very lengthy and most of it is of quite little use in most situation. While the full output should stay, I would like to find a way to shorten it and return only the errors as default;
- JSON output is the most useful (can be parsed by jq and used in different ways) but I would like to add a human output option;
- when restarting using `--include-tasks` the output is not clear. Decision needs to be taken on what the output should be and then implemented.

## Further reading

A great guide on how the Connect API works in general, and all the endpoints it offers can be found for example in the [Confluent official documentation](https://docs.confluent.io/platform/current/connect/references/restapi.html).