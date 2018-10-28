VERSION:=latest

MAKEFILE_PATH := $(abspath $(lastword $(MAKEFILE_LIST)))
CURRENT_DIR := $(notdir $(patsubst %/,%,$(dir $(MAKEFILE_PATH))))
PARENT_DIR="$(dirname "$(CURRENT_DIR)")"
GOPATH:=$(shell go env GOPATH)

#make sure GOBIN and GOPATH are set!

default: build

build: protobuf
	go get -d -v ./...

	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go install -a -installsuffix cgo -ldflags '-w -s' search.go

container: 
	docker build -t $(CURRENT_DIR):$(VERSION) -f Dockerfile $(GOPATH)/src

protobuf:
	# cd ../../ && \
	# 	protoc --go_out=plugins=grpc:pkg/search \
	# 	--micro_out=./pkg/search \
	# 	-I=./pkg/notes \
	# 	--proto_path=./pkg/search search.proto
	cd ../../ && \
		protoc --go_out=plugins=grpc,Mnotes.proto=github.com/mtbarta/monocorpus/pkg/notes:pkg/search \
		--micro_out=./pkg/search \
		-I=./pkg/notes \
		--proto_path=./pkg/search search.proto
