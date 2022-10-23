# kafka-connect-cli Intallation Guide

In this page we can find all the required information to compile and run the CLI. Various options exist and all of them are descirbed to some detail.

## Building locally (Go must be installed)

To build the executable from the source code, simply run the following commands in the command line. Note that this requires Go to be installed in the system used to build the CLI.

In the main directory run 
```(shell)
go mod download
```
Then move to `cli` directory and run 
```(shell)
go build -o kconnect-cli<.extension>
```
to build for the current system. Notice the extension is generally only required if building for Windows. 

If the CLI is being built for multiple systems (or for a system different than the one used, e.g. a Windows executable built on a Linux sytem), please specify `GOOS` and `GOARCH` as environment variables or directly within the build command. Also notice tha in case the CLI is being built for Linux systems, it is advisable to disable CGO (by using the `CGO_ENABLED=0` option) in order to force the final executable to be statically linked. Since the CLI uses the `os` package otherwise the resulting executable may be dynamically linked which could create issues executing it. Below is a complete command for a Linux build.

```(shell)
env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o /builder/output/kconnect-cli_$CLIVERSION_linux
```

## Using Docker to compile the executable

An alternative to compiling the executable directly is to use Docker to do that for us. The Dockerfile provided in this repository is structured as a multistage build, and includes a builder stage that leverages the go-alpine image to compile the CLI and produce the executables as output. Then depending on how the docker build is called, these executables (for Windows and Linux with AMD64 architectures for the time being) are either downloaded locally for use or distribution, or packaged in an Ubuntu.

To dowload the compiled executables locally, use (from the root folder of the repository)

```
docker build --target artifact --output type=local,dest=</path/to/installation> .
```

The resulting artifacts will appear in a folder called `build-output` inside of the destination path indicated. Note that it is usually best to provide a full path, but relative paths will also work.

## Building and running in Docker

## Running in k8s



guide

```
docker build (--build-arg version=0.0.1) --target artifact --output type=local,dest=.\installation\ .
```

```
docker build (--build-arg version=0.0.1) -t local/kconnect-cli:0.0.1 .
```

```
docker run --rm -d local/kconnect-cli:0.0.1
```
