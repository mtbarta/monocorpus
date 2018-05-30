VERSION?=latest

PARENT_DIR="$(dirname "$(CURDIR)")"

#make sure GOBIN and GOPATH are set!

default: build

build:
	npm run build
	
container:
	docker build -t $(CURDIR):$(VERSION) .