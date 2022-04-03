#! /bin/bash

export GOARCH="amd64"
export GOOS="linux"
go build

export GOOS="windows"
go build

unset GOOS
unset GOARCH