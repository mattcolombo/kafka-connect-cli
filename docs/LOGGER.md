# logger

Allows to list the loggers running in a Connect worker, and modify the log levels for specific connector plugins or loggers. Notice that this is specific to the worker, rather than to the cluster. Therefore the list or modify commands will be issues to all the Connect workers one after the other. For this reason it is important that in the CLI configuration file, all the hosts for a specific cluster are specified. 

Further information on the use of the Connect API to control loggers at runtime (notice, this does not require restarting the workers) please consult this section of the [Confluent documentation](https://docs.confluent.io/platform/current/connect/logging.html#using-the-kconnect-api).

## list

`kconnect-cli logger list`: provides a list of the currently configured loggers and their respective log levels in each of the Connect workers present. This command uses the `GET /admin/loggers` endpoint for each of the hosts specified in the CLI configuration file.

## get

`kconnect-cli logger get`: requires the logger/plugin class name as first positional argument. Provides the log level for the logger specified in each of the Connect workers present. This command uses the `GET /admin/loggers/(string:plugin-class)` endpoint for each of the hosts specified in the CLI configuration file.

## set

`kconnect-cli logger list`: requires the logger/plugin class name as first positional argument. Flag `--level` for the log level to set is available, if absent the level will default to `ERROR`. Allows to set the desired log level for the specified logger in each of the Connect workers present. Allowed log levels are the usual Java/log4j levels which are (in increasing verbosity order) `OFF`, `FATAL`, `ERROR`, `WARN`, `INFO`, `DEBUG`, `TRACE`.This command uses the `PUT /admin/loggers/(string:plugin-class)` endpoint for each of the hosts specified in the CLI configuration file.

**NOTE:** The log level set in this way is persisted only until the Connect worker gets restarted. Once the Connector worker restarts the log levels for all loggers revert back to the ones set in the Connect worker configuration.
