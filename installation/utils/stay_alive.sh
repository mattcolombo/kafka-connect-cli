#!/usr/bin/sh

## NOTE: when working on Windows, always make sure that this file is saved with LF line endings (rather than CRLF).
#        Depending on how Git is setup, the files may be retrieved as CRLF, but if so, then building the Dockerfile will cause the 
#        resulting image to not stay alive since this script will not work (it will throw a "file not found error").

## handle here anything that may need to be gracefully terminated when a TERM signal is received
die_func() {
    echo "Terminating"
    exit 1
}

# when TERM signal is given, the above function executes before the script exits
trap die_func TERM

## the below block allows to start the "sleep" process as a background process, and waits for it to finish. 
#  This way if a TERM signal is received, wait will interrupt the background command and execute the exit
#  routines instead than waiting for sleep to finish (which may take a while)
echo "Sleeping..."
# the script will complete and hence terminate automatically after one day
sleep 86400 &
wait