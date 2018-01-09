#!/bin/sh
export GOPATH=$(pwd)

# run: ./lconsole
echo "Building...(lconsole)"
go build -o lconsole main.go
