#!/bin/bash
export GOROOT=$(pwd)/.go_runtime/go
export PATH=$GOROOT/bin:$PATH
export GOPATH=$(pwd)/.go_path
mkdir -p $GOPATH

echo "Go environment configured."
go version
