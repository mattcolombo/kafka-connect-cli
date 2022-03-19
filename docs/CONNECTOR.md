# connector

Allows to list, gather information and manage connectors.

## list

`kafka-connect-cli connector list` provides a list of the connectors currently present in the Connect cluster. Uses the `GET /connectors` Rest endpoint. Allows the following optional flags:

* `--show-info`: expands the configuration for each connector (using the `expand=info` query parameter). :warning: using this operation will expose the complete configuration of the connector, including credentials if present
* `--show-status` expands the status of each connector and its tasks (using the `expand=status` query parameter)

The above flags can be used at the same time, in which case both the status and configuration information will be shown.
