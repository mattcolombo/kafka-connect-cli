# task

Allows to list, gather information and manage connector tasks.

## list

`kafka-connect-cli task list` produces a list of the tasks for the connector specified using the `--name` flag (shorthand `-n`), including each task's configuration. This command uses the `GET /connectors/(string:name)/tasks`.
