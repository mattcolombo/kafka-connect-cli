# logger

Allows to list the loggers running in a Connect worker, and modify the log levels for specific connector plugins or loggers. Notice that this is specific to the worker, rather than to the cluster. Therefore the list or modify commands will be issues to all the Connect workers one after the other. For this reason it is important that in the CLI configuration file, all the hosts for a specific cluster are specified. 

Further information on the use of the Connect Rest API to control loggers at runtime (notice, this does not require restarting the workers) please consult this section of the [Confluent documentation](https://docs.confluent.io/platform/current/connect/logging.html#using-the-kconnect-api).

## list

`kafka-connect-cli connector list` provides a list of the currently configured loggers and their respective log levels for each of the Connect workers. This command uses the `GET /admin/loggers` endpoint for each of the hosts specified in the CLI configuration file.
