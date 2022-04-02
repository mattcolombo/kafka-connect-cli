# task

Allows to list, gather information and manage connector tasks.

## list

`kconnect-cli task list`: requires flags `--name` flag (shorthand `-n`) for the connector name. Produces a list of the tasks for the connector specified, including each task's configuration. This command uses the `GET /connectors/(string:name)/tasks`.

## get

`kconnect-cli task restart`: requires flags `--name` flag (shorthand `-n`) for the connector name, and `--id` for the task ID. Gathers information on the status of task `id` for connector `name`. This command uses the `GET /connectors/(string:name)/tasks/(int:taskID)/status`.

## restart

`kconnect-cli task restart`: requires flags `--name` flag (shorthand `-n`) for the connector name, and `--id` for the task ID. Restarts task `id` for connector `name`. This command uses the `POST /connectors/(string:name)/tasks/(int:taskID)/restart`.
