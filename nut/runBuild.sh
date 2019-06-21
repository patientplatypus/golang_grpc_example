#!/bin/bash

export GO111MODULE=on

rm -rf go.mod go.sum

go mod init secretsquirrel_nest

protoc --proto_path="$(pwd)/../nest/_proto/" --go_out=plugins=grpc:"$(pwd)/protomain" main.proto

go get google.golang.org/grpc
go get github.com/golang/protobuf/protoc-gen-go
go get github.com/improbable-eng/grpc-web/go/grpcweb

go build main.go

chmod 777 ./main && ./main

