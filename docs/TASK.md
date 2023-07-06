# task

Allows to list, gather information and manage connector tasks.

## list

`kconnect-cli task list`: requires the connector name as first positional argument. Produces a list of the tasks for the connector specified, including each task's configuration. This command uses the `GET /connectors/(string:connector_name)/tasks`.

## get

`kconnect-cli task get`: requires the connector name as first positional argument, and flag `--id` for the task ID. Gathers information on the status of task `taskID` for connector `connector_name`. This command uses the `GET /connectors/(string:connector_name)/tasks/(int:taskID)/status`.

## restart

`kconnect-cli task restart`: requires the connector name as first positional argument, and flag `--id` for the task ID. Restarts task `taskID` for connector `connector_name`. This command uses the `POST /connectors/(string:connector_name)/tasks/(int:taskID)/restart`.
