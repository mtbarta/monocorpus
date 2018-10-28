VERSION:=latest

MAKEFILE_PATH := $(abspath $(lastword $(MAKEFILE_LIST)))
CURRENT_DIR := $(notdir $(patsubst %/,%,$(dir $(MAKEFILE_PATH))))
PARENT_DIR="$(dirname "$(CURRENT_DIR)")"
GOPATH:=$(shell go env GOPATH)
#make sure GOBIN and GOPATH are set!

default: build

build: protobuf
	go get -d -v ./...

	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go install -a -installsuffix cgo -ldflags '-w -s' notes.go

container: protobuf
	docker build -t $(CURRENT_DIR):$(VERSION) -f Dockerfile $(GOPATH)/src

protobuf:
	cd ../../pkg/notes && protoc --go_out=plugins=grpc:. --micro_out=. notes.proto
