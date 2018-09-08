VERSION:=3.6.2-jessie

MAKEFILE_PATH := $(abspath $(lastword $(MAKEFILE_LIST)))
CURRENT_DIR := $(notdir $(patsubst %/,%,$(dir $(MAKEFILE_PATH))))
PARENT_DIR="$(dirname "$(CURRENT_DIR)")"

default: container

container:
	docker build -t monocorpus/$(CURRENT_DIR):$(VERSION) -f Dockerfile .