#!/bin/sh

env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -tags=jsoniter -buildmode=exe -trimpath -ldflags "-w -s -extldflags -static" -o main-linux main.go
env CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -tags=jsoniter -buildmode=exe -trimpath -ldflags "-w -s -extldflags -static" -o main-darwin main.go
