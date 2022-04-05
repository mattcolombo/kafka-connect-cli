# kafka-connect-cli Intallation Guide

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

# Manually (presumes Go installed)

In the main directory run 
```(shell)
go mod download
```

Go to `cli` directory and run 
```(shell)
go build -o kconnect-cli<.extension>
```
to install for the current system. Else specify `GOOS` and `GOARCH` as environment variables to build for a different system or architecture.