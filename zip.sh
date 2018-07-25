#!/bin/bash
export GOPATH=${PWD}
export PATH=$PATH:$(go env GOPATH)/bin

go get github.com/aws/aws-lambda-go/lambda
go get github.com/golang/protobuf/proto
go get github.com/golang/protobuf/ptypes/timestamp


rm main main.zip
GOOS=linux GOARCH=amd64 go build -o main main.go
zip main.zip main