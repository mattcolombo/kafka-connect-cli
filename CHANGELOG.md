# CHANGELOG

Throughout this document, :boom: will denote a breaking change

# v2.0.0

This version was marked as 2.0 since it contains some design changes and breaking changes.

* :boom: `logger get` and `logger set` now take the logger/plugin class name as a positional argument and no longer as a flag
* `--level` flag for `logger set` is no longer mandatory and defaults to INFO when not provided
* :boom: All `task` commands now take the connector name and task ID as positional arguments and no longer as flags
* :boom: All `connector` commands (where applicable) now take the connector name as a positional argument and no longer as a flag
* :boom: `--config-path` flag in `connector create` and `connector update` commands has been changed to `--config-file` and a more familiar `-f` shorthand has been added
* Updated the utility script `connector-full-restart.sh` is updated to reflect the changes mentioned above
* Documentation updated
* Added `get-config` command for showing the config path and full file
* Removed `--show-config` flag from `version` (replaced by `get-config` command)
* Modified `version` command to provide more detail. Default is a short version with git tag version, build date and last commit. Optional JSON output with further detail added
* Dockerfile builder will now get the correct version and details for the version command directly at build time, instead than having to hardcode them in the CLI itself
* When building the CLI manually using `go build` alone, all the details on the version will be flagged as `manual_build`
* Added a utility script (bash only) that does the full restart of a connector by name; this is useful for older versions of connect that are not having the option for a full restart using the native Rest API commands
* Added build target for MacOSX as well (darwin architecture)

# v1.0.0

Initial Release

* Configuration file (mandatory) allows to specify the details about the Kafka Connect instance to target
* Full management of connectors and connector tasks (including deployments and validation)
* Management of log level at runtime
* Full instructions for building and usage including: building the executables, running in Docker and running in k8s
