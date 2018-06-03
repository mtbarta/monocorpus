VERSION?=latest

MAKEFILE_PATH := $(abspath $(lastword $(MAKEFILE_LIST)))
CURRENT_DIR := $(notdir $(patsubst %/,%,$(dir $(MAKEFILE_PATH))))
PARENT_DIR="$(dirname "$(CURRENT_DIR)")"
#make sure GOBIN and GOPATH are set!

default: build

build:
	npm run build
	
container:
	docker build -t $(CURRENT_DIR):$(VERSION) .