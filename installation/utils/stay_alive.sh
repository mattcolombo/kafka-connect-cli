#!/bin/sh

## handle here anything that may need to be gracefully terminated when a TERM signal is received
die_func() {
    echo "Terminating"
    exit 1
}

# when TERM signal is given, the aboev function executes before the script exits
trap die_func TERM

## the below block allows to start the "sleep" process as a background process, and waits for it to finish. 
#  This way if a TERM signal is received, wait will interupt the background command and execute the exit 
#  routines instead than waiting for sleep to finish (which may take a while)
echo "Sleeping..."
# the script will complete and hence teminate automatically after one day
sleep 86400 &
wait