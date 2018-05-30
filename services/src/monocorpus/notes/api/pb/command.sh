#!/bin/bash

#make sure to have protoc-gen-go on your path. this was in /home/matt/go/bin for me, not in /usr/local

protoc --go_out=plugins=grpc:. notes.proto

