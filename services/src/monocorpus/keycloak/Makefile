VERSION?=latest

PARENT_DIR="$(dirname "$(CURDIR)")"

#make sure GOBIN and GOPATH are set!

default: build

build: ;
	#pass

container:
	docker build -t $(CURDIR):$(VERSION) .