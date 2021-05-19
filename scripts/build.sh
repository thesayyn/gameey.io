#!/bin/bash
export GOOS="linux"
export GOARCH="amd64"
go build let/cmd.go

curl --upload-file ./cmd https://transfer.sh/let