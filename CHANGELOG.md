# CHANGELOG

Throughout this document, :boom: will denote a breaking change

# v1.1.0

* :boom: `logger get` and `logger set` now take the logger/plugin class name as a positional argument and not as a flag any longer
* :boom: `--level` flag for `logger set` is now not mandatory anymore and defaults to ERROR if not specified
* [add extra changes] 
* Documentation updated

# v1.0.0

* Initial Release
* Configuration file (mandatory) allows to specify the details about the Kafka Connect instance to target
* Full management of connectors and connector tasks (including deployments and validation)
* Management of log level at runtime
* Full instructions for building and usage including: building the executables, running in Docker and running in k8s
