#!/bin/bash

NAME="$1"

printUsage () {
    echo "Usage: connector-full-restart.sh <connector_name>"
}

# making sure that keytool and openssl are installed
checkCommands () {
    if ! command -v $1 &> /dev/null
    then
        echo "$1 could not be found"
        exit 1
    fi
}

checkCommands kconnect-cli
checkCommands jq

# checking that the number of arguments is the correct one. Printing usage if not
if [[ "$#" -ne 1 ]]; then
    echo "Illegal number of parameters"
    printUsage
    exit 1
fi

# getting the id of the last task from the tasks list
LAST_TASK=$(kconnect-cli task list $NAME | jq -r .[-1].id.task)

# restarting the connector main process
kconnect-cli connector restart $NAME

# restarting all the connector tasks
for ((i=0;i<=LAST_TASK;i++)); do
        kconnect-cli task restart $NAME --id $i
done