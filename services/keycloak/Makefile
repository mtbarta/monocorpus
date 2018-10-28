VERSION?=latest

MAKEFILE_PATH := $(abspath $(lastword $(MAKEFILE_LIST)))
CURRENT_DIR := $(notdir $(patsubst %/,%,$(dir $(MAKEFILE_PATH))))
PARENT_DIR="$(dirname "$(CURRENT_DIR)")"

default: build

build: ;
	#pass

container:
	docker build -t jboss/keycloak-postgres:4.0.0.Beta1 -f ./base/Dockerfile ./base
	docker build -t $(CURRENT_DIR):$(VERSION) .
