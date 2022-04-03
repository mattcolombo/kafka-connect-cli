# adding an argument as the version, to add to the executables once built
ARG CLIVERSION=0.0.1

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
RUN GOOS=linux GOARCH=amd64 go build -o /builder/output/kconnect-cli_$CLIVERSION
RUN GOOS=windows GOARCH=amd64 go build -o /builder/output/kconnect-cli_$CLIVERSION.exe

FROM scratch as artifact
COPY --from=builder /builder/output/kconnect-cli* /build-output/

# building the actual useable image
FROM alpine:latest
# refreshing ARG value for current image
ARG CLIVERSION
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
COPY --from=builder /builder/output/kconnect-cli_$CLIVERSION /usr/cli/kconnect-cli
ENV PATH="/usr/cli:${PATH}"