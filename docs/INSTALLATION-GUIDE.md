# Installation Guide

In this page we can find all the required information to compile and run the CLI. Various options exist and all of them are described to some detail.

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

If the CLI is being built for multiple systems (or for a system different from the one used, e.g. a Windows executable built on a Linux sytem), please specify `GOOS` and `GOARCH` as environment variables or directly within the build command. Also notice tha in case the CLI is being built for Linux systems, it is advisable to disable CGO (by using the `CGO_ENABLED=0` option) in order to force the final executable to be statically linked. Since the CLI uses the `os` package otherwise the resulting executable may be dynamically linked which could create issues executing it. Below is a complete command for a Linux build.

```(shell)
env GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o /builder/output/kconnect-cli_$CLIVERSION_linux
```
**NOTE:** building in this way will mean that the `version` command will not return any useful information aside from the fact that it was built manually since all fields will be set to `manual_build`. Actual values can be passed through LDFLAGS at build time. This is done in the Dockerfile so checking there is advised on the specific variables to set.

## Download from GitHub release

In order to download directly a version of the CLI without using a browser, it's possible to use `cURL` to download the binary for the required version and system.

### Install on Linux/MacOS

To install the Linux version, we can use the following commands to set the desired version to download
```(bash)
export VERSION="vM.m.p"
```
or alternatively get the latest version from GitHub
```(bash)
export VERSION=$(curl -s https://api.github.com/repos/mattcolombo/kafka-connect-cli/releases/latest | jq -r .tag_name)
```
Once the `VERSION` variable is set, for Linux systems we can simply run
```
curl -LO https://github.com/mattcolombo/kafka-connect-cli/releases/download/$VERSION/kconnect-cli_linux_amd64_$VERSION
```
while for MacOS we can use (for `amd64` architectures)
```
curl -LO https://github.com/mattcolombo/kafka-connect-cli/releases/download/$VERSION/kconnect-cli_darwin_amd64_$VERSION
```
or (for `arm64` architectures)
```
curl -LO https://github.com/mattcolombo/kafka-connect-cli/releases/download/$VERSION/kconnect-cli_darwin_arm64_$VERSION
```

At this point the binary is available so we can change the permissions as required (by default the binary will not be executable), add the download location ot the path or move it to a path location (usuall `/usr/local/bin` or `/usr/bin`).

### Install on Windows

To download the Windows version without a browser we can follow more or less the same process as seen for linux. In powershell, first set the desired version in an environment variable using
```(powershell)
$env:VERSION="vM.m.p"
```
and then download it using `cURL` as
```(powershell)
curl -LO https://github.com/mattcolombo/kafka-connect-cli/releases/download/$env:VERSION/kconnect-cli_win_amd64_$env:VERSION.exe
```


## Using Docker to compile the executable

An alternative to compiling the executable directly is to use Docker to do that for us. The Dockerfile provided in this repository is structured as a multistage build, and includes a builder stage that leverages the go-alpine image to compile the CLI and produce the executables as output. Then depending on how the docker build is called, these executables (for Windows and Linux with AMD64 architectures for the time being) are either downloaded locally for use or distribution, or packaged in an Ubuntu.

To download the compiled executables locally, use (from the root folder of the repository)

```
docker build --target artifact --output type=local,dest=</path/to/installation> .
```

The resulting artifacts will appear in a folder called `build-output` inside of the destination path indicated. Note that it is usually best to provide a full path, but relative paths will also work.

Compiled artifacts for additional architectures/OSs will be added in the future if the requirement arises.

## Building and running in Docker

One may also wish to run the CLI in Docker directly. This may create some challenges with network connectivity depending on where Docker and Kafka Connect run, but assuming connectivity is present, this is completely possible. To do this there are two options.

### Pulling the prebuilt image from Docker Hub

Whenever a new version of the CLI is released, a Docker image will be available in Docker Hub with the latest version of the CLI prepackaged in. To use this simply pull the image using

```
docker pull mattcolombo/kafka-connect-cli:<tag>
```

and use it as described belolw.

### Building the Docker image locally

In case one would like to build the image locally (maybe due to using a slightly newer version, or some customised code from a fork of the official repo) this can be done by simply building the Dockerfile frovided. From the root of the repository, simply run

```
docker build -t <docker-repo>/<image>:<tag> .
```

using the desired names for the Docker repo, image and tag, and let the builder do its thing. Once completed, your new image is ready to use locally (or to push to whatever image repository you would like).

### Running the CLI image locally in Docker

Once the image is available to Docker (either having pulled the one from Docker Hub, or having built a custom one) the repository can be used directly through `docker run` commands. There is however a catch. In order to provide the configuration files for the CLI (see the [configuration documentation](/docs/CONFIGURATION.md)) to the container we need to mount a local volume containing the correct config file(s).

To do so, first create the config file(s) as described (the template [here](/samples-templates/kconnect-cli-config-template.yaml) can also be used as a starting point) and save them in a local folder. Next use the `docker run` command with the options to mount that location in the container. This can be done using

```
docker run --rm -d --mount type=bind,source=<absolute-path-to-source-dir>,target=/usr/cli/config,readonly <docker-repo>/<image>:<tag>
```

This will mount the directory with the configuration file in the container at `/usr/cli/config` at which point the files can be used as described in the [configuration documentation](/docs/CONFIGURATION.md).

--**NOTE**-- since the [stay_alive](/installation/utils/stay_alive.sh) script is set as the startup command of the container, running the command above will create a container that will simply sleep for one day and then terminate (unless terminated before by the user). This way the container can be started, and then once can ssh inside the container and run the CLI commands when needed. This is especially useful if running the CLI container in an environment like kubernetes. If on the other hand one would like to simply execute a shell directly in the container, and destroy it when done, simply add the `-it` flag and the `bash` command to `docker run`. Also, the `-d` flag needs to be removed, else the container will start in detached mode and not return the shell. This way the startup script will be skipped, a shell will be obtained directly in the container and once the user closes the session the container will be destroyed. Complete command for this would be

```
docker run --rm -it --mount type=bind,source=<absolute-path-to-source-dir>,target=/usr/cli/config,readonly <docker-repo>/<image>:<tag> bash
```

## Running in k8s

One alternative way to run the Docker image is to use Kubernetes. This is quite easy to do and can be achieved simply by creating a secret with the configuration file(s) required and creating a deployment that will start the container from the image discussed above. Notice that when run in k8s the default startup behavior will apply (unless that has been changed). This means that the `stay_alive.sh` script (found [here](/installation/utils/stay_alive.sh)) will be run at startup, which will cause the pod to remain alive for one entire day unless killed before by the user. However since we are running in k8s, after the sleep time has expired and the container terminates, the k8s controller will take care of starting it once again. In this way the pod will always be available save for a very short time every 24 hours when the pod will terminate and get recreated.

-- **NOTE** -- Since the pod is getting recreated every day, any environment variables (specifically the `CONNECTCFG` one necessary for the CLI to run) will be lost. When working with one configuration file only, this can be mitigated by setting the correct environment variable in the deployment manifest.

In order to provide the correct configuration file(s) to the container, we will load it from a k8s secret. Note that a secret was chosen over a config map simply since this may contain some sensitive information (passwords and so on). However, a config map can be used easily in a very similar way and to the same effect. First of all create the secret for the configuration file (this can be done by editing as required the template provided [here](/samples-templates/aks/secret-config-template.yaml)). In case multiple configuration files are required they can all be stored in the same secret (as can be seen [here](/samples-templates/aks/secret-multi-config-template.yaml)).

Once the secret template is ready, simply create the secret using 

```
kubectl apply -f </path/to/secret/manifest.yaml>
```

Notice that the names given to the entries of the secret will be the names of the files mounted in the location selected.

Once this is done, the configuration file(s) are available for mounting in the container, so simply customise the deployment manifest template found [here](/samples-templates/aks/deployment-template.yaml) and once again run

```
kubectl apply -f </path/to/deployment/manifest.yaml>
```

After few seconds, the pods running the CLI will be ready and can be executed in to run any commands required. Of course this will depend on the connectivity from this pod and the Kafka Connect required. Notice that the requests provided in the manifest are simply an example, and can be tweaked as required. Limits are not provided since in many cases they represent an anti-pattern, but they can be easily added if required. 