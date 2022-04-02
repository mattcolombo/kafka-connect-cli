# cluster

Allows to gather information on a Connect cluster, including if desired the list of connector plugins currently installed.

## get

`kconnect-cli cluster get` provides the information on the currently installed cluster. This information includes the Connect version installed and the ID of the Kafka cluster that Connect is bound to. This command uses the `GET /` endpoint. 

Including the flag `--show-plugins` will show the information as above, and in addition show a list of installed Connector plugins installed (including the version reported by the plugin) using the `GET /connector-plugins` endpoint.
