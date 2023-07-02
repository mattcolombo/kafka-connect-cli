# connector

Allows to list, gather information and manage connectors.

All the commands that require a file path as input, require a connector configuration file in JSON format. A sample connector configuration file as required in all these commands can be found [here](/samples-templates/connector-config-template.json).

## list

`kconnect-cli connector list`: provides the list of all connectors currently deployed in the Connect cluster. Uses the `GET /connectors` endpoint. Allows the following optional flags:

* `--show-info`: expands the configuration for each connector (using the `expand=info` query parameter). 
* `--show-status` expands the status of each connector and its tasks (using the `expand=status` query parameter)

The above flags can be used at the same time, in which case both the status and configuration information will be shown.

:danger: --**NOTE**-- using this operation will expose the complete configuration of the connector, including credentials if present

## get

`kconnect-cli connector get`: requires the connector name as first positional argument. Allows the optional flags `--status-only` (shorthand `-s`) and `--config-only` (shorthand `-c`). Provides information on the connector specified in the argument, together with the current configuration and status of the connector and tasks. Use the `--config-only` optional flag to show _only_ the configuration and the `--status-only` optional flag to show _only_ the status of the connector and related tasks. Uses the `GET /connectors/(string:connector_name)` (optionally depending on the flags, with also the `/status` or `/config` paths added at the end).

:danger: --**NOTE**-- as already stated in the previous section, showing the configuration will expose in clear text any confidential information contained in the connector configuration (including credentials) unless they are secured in some other way (see [this page](https://docs.confluent.io/platform/current/connect/security.html#externalizing-secrets) for possible solutions to this issue; however securing Connect is out of the scope of this project and will not be discussed further here).

--**NOTE**-- the `--config-only` and `--status-only` flags are mutually exclusive, therefore only one can be used at the same time.

## create

`kconnect-cli connector create`: requires flag `--config-file` (shorthand `-f`) flag for the path to the connector configuration file in JSON format; allows optional boolean flag `--validate`. Creates a new connector using the configuration file selected. If the `--validate` flag is added, the connector is _NOT_ created, but instead the configuration is validated against the connector plugin class. Uses the `POST /connectors` endpoint for the creation, and `PUT /connector-plugins/(string:plugin_type)/config/validate` endpoint for the validation.

--**NOTE**-- the validation of the connector configuration will throw at least one error due to the name missing in the configuration. This is because this endpoint (and exclusively this one) actually requires the connector name to be part of the configuration part of the JSON, rather than separate as requested for every other command. This is not an issue, as long as the user is aware that the `missing name` error is always going to be thrown. If this is the only error, the connector can be uploaded without failures.

## update

`kconnect-cli connector update`: requires the connector name as first positional argument. Requires the `--config-file` (shorthand `-f`) flag for the path to the connector configuration file in JSON format. Updates the configuration of the connector specified using the file provided. Before doing so, checks that the name specified passed as argument matches the connector name specified in the configuration file. If the names don't match, throws an error. Uses the `PUT /connectors/(string:connector_name)/config` endpoint.

## delete

`kconnect-cli connector delete`: requires the connector name as first positional argument. Deletes the connector specified. Uses the `DELETE /connectors/(string:connector_name)` endpoint.

## pause

`kconnect-cli connector pause`: requires the connector name as first positional argument. Brings the connector in PAUSED state, where processing is suspended. Uses the `PUT /connectors/(string:connector_name)/pause` endpoint.

--**NOTE**-- by using this operation the task(s) will continue to exist, though the processing will be interrupted. 

## resume

`kconnect-cli connector resume`: requires the connector name as first positional argument. Resumes a previously paused task and allow message processing to continue. Uses the `PUT /connectors/(string:connector_name)/resume` endpoint.

--**NOTE**-- as stated above for the pause operation, the resume operation will not remove the task but only un-suspend it; it will therefore also _NOT_ restart a task (as in the thread tied to the task), only resume the processing for it. If a task is in FAILED state, it will remain in the same state even after a pause and resume. To actually restart a task (or a connector) use the `task restart` or `connector restart` actions.

## restart

`kconnect-cli connector restart`: requires the connector name as first positional argument. Allows boolean flags `--include-tasks` and `--failed-only`. In the vanilla version, restarts the connector specified in the argument. Note that this only restarts the connector process itself, it does _NOT_ restart any of the tasks owned by said connector. Tasks will need to be restarted using the `kconnect-cli task restart` command. However, adding the `--include-tasks` will restart the connector and all the related tasks. Using the `--failed-only` flag will only restart the tasks (and the connector itself) in the case that the state is FAILED. Uses the `PUT /connectors/(string:connector_name)/restart` endpoint; adds the `includeTasks=true` and `onlyFailed=true` if the relative flags are selected.

--**NOTE**-- the query parameters `includeTasks=true` and `onlyFailed=true` were only added to Kafka Connect as part of [KIP-745](https://cwiki.apache.org/confluence/pages/viewpage.action?pageId=181308623). They were added quite recently, and therefore will not work on older versions of Connect. If such is the case, adding the flags will not cause the command to fail, but will also not make any difference. On older versions of Connect the flags are redundant and the behaviour of this command will be the same with or without flags.

To work around this point, a utility shell script is provided in the [utility-scripts](/utility-scripts/) folder. The [connector-full-restart.sh](/utility-scripts/connector-full-restart.sh) script allows to provide the name of the connector to fully restart, extracts the number of tasks and restarts the main connector process and recursively all the tasks. This will not make a difference between tasks that are running and tasks that are not running, but will at least provide a quick way to fully restart a connector without having to run the restart-task commands of this CLI many times. This is equivalent to `kconnect-cli connector restart --include-tasks`.