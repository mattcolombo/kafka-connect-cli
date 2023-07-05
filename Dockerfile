# adding an argument as the version, to add to the executables once built
ARG GITVERSION=v1.1.0
ARG MAJVERSION=1
ARG MINVERSION=0

# defining the build environment
FROM golang:alpine AS builder 
# refreshing ARG value for current image
ARG GITVERSION
ARG MAJVERSION
ARG MINVERSION
ARG PACKAGE=github.com/mattcolombo/kafka-connect-cli/cmd

# installing git in the container so that I can find the hash
RUN apk update && apk add git

# creating the working directory, adding the module and sums file and installing the dependencies
WORKDIR /builder 
COPY . /builder
RUN go mod download

# building the linux and windows executable
WORKDIR /builder/cli/
#get the information about git hash , build timestamp and go version
RUN export COMMIT_HASH=$(git rev-parse --short HEAD)
RUN echo $COMMIT_HASH
RUN export BUILD_TIMESTAMP=$(date '+%Y-%m-%dT%H:%M:%S')
RUN export GO_VERSION=$(go version | awk {'print $3'})
#RUN GIT_COMMIT=$(git rev-list -1 HEAD)
ARG LDFLAGS="-X '$PACKAGE/version.MajorVersion=$MAJVERSION' \
    -X '$PACKAGE/version.MinorVersion=$MINVERSION' \
    -X '$PACKAGE/version.GitVersion=$GITVERSION' \
    -X '$PACKAGE/version.GitHash=$COMMIT_HASH' \
    -X '$PACKAGE/version.BuildDate=$BUILD_TIMESTAMP' \
    -X '$PACKAGE/version.GoVersion=$GO_VERSION'"
# control statement TODELETE
RUN echo $LDFLAGS 
RUN env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="$LDFLAGS" -o /builder/output/kconnect-cli_linux-amd64_$GITVERSION
RUN env GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="$LDFLAGS" -o /builder/output/kconnect-cli_win_amd64_$GITVERSION.exe

FROM scratch as artifact
COPY --from=builder /builder/output/kconnect-cli* /build-output/

# building the actual useable image
FROM ubuntu:22.04
# refreshing ARG value for current image
ARG GITVERSION
# Installing jq for ease of management of JSON responses and vim to allow for managing files if required
RUN apt update; apt install jq -y; apt install vim -y
# Setting the working directory in the root and adding the script managing the sleep and the graceful exit if required
WORKDIR /background
COPY ./installation/utils/stay_alive.sh /background/stay_alive.sh
# making sure the script is executable
RUN chmod +x /background/stay_alive.sh
# setting the command to run when the container is started
CMD ["/background/stay_alive.sh"]
# creating the working directory, adding the built executable from the previous step and adding the current workdir to the PATH
WORKDIR /usr/cli
COPY --from=builder /builder/output/kconnect-cli_linux-amd64_$GITVERSION /usr/cli/kconnect-cli
# adding the working directlry to the path, so that the CLI is accessible from any location
ENV PATH="/usr/cli:${PATH}"