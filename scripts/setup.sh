#!/bin/bash

set -x

#1. sudo su
#2. add below to /etc/profile:
#   export GOROOT=/usr/local/go
#   export GOBIN=$GOROOT/bin
#   export PATH=$PATH:$GOBIN
#3. run below to install plugins:
#   go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
#   go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
