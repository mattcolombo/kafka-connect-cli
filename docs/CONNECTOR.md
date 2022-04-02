# connector

Allows to list, gather information and manage connectors.

## list

`kconnect-cli connector list`: provides a list of the connectors currently present in the Connect cluster. Uses the `GET /connectors` Rest endpoint. Allows the following optional flags:

* `--show-info`: expands the configuration for each connector (using the `expand=info` query parameter). 
:warning: using this operation will expose the complete configuration of the connector, including credentials if present
* `--show-status` expands the status of each connector and its tasks (using the `expand=status` query parameter)

The above flags can be used at the same time, in which case both the status and configuration information will be shown.

## get

`kconnect-cli connector get`: requires flag `--name` flag (shorthand `-n`) for the connector name, and allows optional flags `--status` and `--config`. Provides information on the connector specified by the `--name` flag, together with the current configuration and status of the connector and tasks. Use the `--config` optional flag to show _only_ the configuration and the `--status` optional flag to show _only_ the status of the connector and related tasks. 

:warning: as also stated in the previous command, showing the configuration will expose in clear text any confidential information contained in the connector configuration (including credentials) unless they are secured in some other way (see [this page](https://docs.confluent.io/platform/current/connect/security.html#externalizing-secrets) for possible solutions to this issue; however securing Connect is out of the scope of this project so no more will be discussed here).

:warning: the `--config` and `--status` flags are mutually exclusive, therefore only one can be used at the same time.

## pause

`kconnect-cli connector pause`: requires flag `--name` flag (shorthand `-n`) for the connector name. Brings the connector in PAUSED state, where processing is suspended. Notice that by using this operation the tasks continue to exist, though the processing will be interrupted. Uses the `PUT /connectors/(string:name)/pause` Rest endpoint.

## resume

`kconnect-cli connector resume`: requires flag `--name` flag (shorthand `-n`) for the connector name. Resumes a previously paused task and allow message processing to continue. Notice that as stated above pause will not remove the task, only suspend it, so this will not restart a task, only resume the processing. If a task is in FAILED state, it will continue to be after a pause and resume. To actually restart a task (or a connector) use the `task restart` or `connector restart` actions. Uses the `PUT /connectors/(string:name)/resume` Rest endpoint.

## delete

`kconnect-cli connector delete`: requires flag `--name` flag (shorthand `-n`) for the connector name. Deletes the connector specified. Uses the `DELETE /connectors/(string:name)` Rest endpoint.