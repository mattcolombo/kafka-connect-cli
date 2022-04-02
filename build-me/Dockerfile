# defining the build environment
FROM golang:alpine AS builder 
# creating the working directory, adding the module and sums file and installing the dependencies
WORKDIR /builder 
COPY ../go.mod /builder/
COPY ../go.sum /builder/
RUN go mod download

# transferring the main executable file and the rest of the packages and building the linux executable
#COPY ./*.go /builder/
#COPY ./cmd/ /builder/cmd/
RUN GOOS=linux GOARCH=amd64 go build -o kconnect-cli

# building the actual useable image
FROM alpine:latest
# Setting the working directory in the root and adding the script managing the sleep and the graceful exit if required
WORKDIR /background
COPY ./utils/stay_alive.sh /background/stay_alive.sh
# making sure the script is executable
RUN chmod +x /background/stay_alive.sh
# setting the command to run when the container is started
CMD ["/background/stay_alive.sh"]

# creating the working directory, adding the built executable from the previous step and adding the current workdir to the PATH
WORKDIR /usr/cli
COPY --from=builder /builder/kconnect-cli /usr/cli/kconnect-cli
ENV PATH="/usr/cli:${PATH}"