# adding an argument as the version, to add to the executables once built
ARG GITVERSION=v2.0.0
ARG MAJORVERSION=2
ARG MINORVERSION=0

# defining the build environment
FROM golang:1.20.5-alpine AS builder 
# refreshing ARG value for current image
ARG GITVERSION
ARG MAJORVERSION
ARG MINORVERSION
ARG PACKAGE=github.com/mattcolombo/kafka-connect-cli/cmd
# installing git in the container so that I can find the hash
RUN apk update && apk add git
# creating the working directory, adding the module and sums file and installing the dependencies
WORKDIR /builder 
COPY . /builder
RUN go mod download
# building the linux and windows executable
WORKDIR /builder/cli/
# creating a file to store the ldflags for go builder
RUN touch ./flags
#get the information about git hash , build timestamp and go version and add them to the flags file
RUN echo -n  "-X '$PACKAGE/version.MajorVersion=$MAJORVERSION'" >> flags
RUN echo -n " -X '$PACKAGE/version.MinorVersion=$MINORVERSION'" >> flags
RUN echo -n " -X '$PACKAGE/version.GitVersion=$GITVERSION'" >> flags
RUN echo -n " -X '$PACKAGE/version.GitHash=$(git rev-parse --short HEAD)'" >> flags
RUN echo -n " -X '$PACKAGE/version.BuildDate=$(date '+%Y-%m-%dT%H:%M:%S')'" >> flags
RUN echo -n " -X '$PACKAGE/version.GoVersion=$(go version | awk {'print $3'})'" >> flags
# run the build command with the flags created above
RUN env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="$(cat flags)" -o /builder/output/kconnect-cli_linux_amd64_$GITVERSION
RUN env GOOS=darwin GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="$(cat flags)" -o /builder/output/kconnect-cli_darwin_amd64_$GITVERSION
RUN env GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="$(cat flags)" -o /builder/output/kconnect-cli_win_amd64_$GITVERSION.exe

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
COPY --from=builder /builder/output/kconnect-cli_linux_amd64_$GITVERSION /usr/cli/kconnect-cli
# adding the working directlry to the path, so that the CLI is accessible from any location
ENV PATH="/usr/cli:${PATH}"