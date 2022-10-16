# adding an argument as the version, to add to the executables once built
ARG CLIVERSION=0.2.0

# defining the build environment
FROM golang:alpine AS builder 
# refreshing ARG value for current image
ARG CLIVERSION
# creating the working directory, adding the module and sums file and installing the dependencies
WORKDIR /builder 
COPY ./go.mod /builder
COPY ./go.sum /builder
RUN go mod download

# transferring the main executable file and the rest of the packages
COPY ./cli/*.go /builder/cli/
COPY ./utilities/ /builder/utilities/
COPY ./cmd/ /builder/cmd/
# building the linux and windows executable
WORKDIR /builder/cli/
RUN echo $CLIVERSION
RUN env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o /builder/output/kconnect-cli_linux-amd64_$CLIVERSION
RUN env GOOS=windows GOARCH=amd64 CGO_ENABLED=0 go build -o /builder/output/kconnect-cli_win_amd64_$CLIVERSION.exe

FROM scratch as artifact
COPY --from=builder /builder/output/kconnect-cli* /build-output/

# building the actual useable image
FROM ubuntu:22.04
# refreshing ARG value for current image
ARG CLIVERSION
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
RUN echo $CLIVERSION
COPY --from=builder /builder/output/kconnect-cli_$CLIVERSION_linux /usr/cli/kconnect-cli
# adding the template configuration file for convenience (it can be modified on-site, or a new image can be built from this one with the correct ones in it)
COPY ./samples/kconnect-cli-config.yaml.tmpl /usr/cli/kconnect-cli-config.yaml
ENV PATH="/usr/cli:${PATH}"