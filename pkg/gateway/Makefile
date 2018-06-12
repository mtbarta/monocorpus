VERSION?=latest

MAKEFILE_PATH := $(abspath $(lastword $(MAKEFILE_LIST)))
CURRENT_DIR := $(notdir $(patsubst %/,%,$(dir $(MAKEFILE_PATH))))
PARENT_DIR="$(dirname "$(CURRENT_DIR)")"

#make sure GOBIN and GOPATH are set!

default: build

build:
	go get -d -v ./...

	cd ../../cmd; \
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go install -a -installsuffix cgo -ldflags '-w -s' gateway.go

container:
	docker build -t $(CURRENT_DIR):$(VERSION) -f Dockerfile $(GOPATH)/src