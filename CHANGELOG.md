# CHANGELOG

Throughout this document, :boom: will denote a breaking change

# v1.1.0

* :boom: `logger get` and `logger set` now take the logger/plugin class name as a positional argument and no longer as a flag
* `--level` flag for `logger set` is no longer mandatory and defaults to ERROR when not provided
* :boom: all `task` commands now take the connector name as a positional argument and no longer as a flag
* `--id` flag in `task restart` and `task get` is no longer mandatory and defaults to `0` when not provided
* :boom: all `connector` commands (where applicable) now take the connector name as a positional argument and no longer as a flag
* :boom: `--config-path` flag in `connector create` and `connector update` commands has been changed to `--config-file` and a more familiar `-f` shorthand has been added
* updated the utility script `connector-full-restart.sh` is updated to reflect the changes mentioned above
* Documentation updated

# v1.0.0

* Initial Release
* Configuration file (mandatory) allows to specify the details about the Kafka Connect instance to target
* Full management of connectors and connector tasks (including deployments and validation)
* Management of log level at runtime
* Full instructions for building and usage including: building the executables, running in Docker and running in k8s
