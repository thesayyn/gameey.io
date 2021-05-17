#!/bin/bash
export PATH="$PATH:$(go env GOPATH)/bin";
protoc -I=gameeylet --go_out=gameeylet protocol/gameeylet.proto