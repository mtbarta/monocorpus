VERSION?=latest

PARENT_DIR="$(dirname "$(CURDIR)")"

#make sure GOBIN and GOPATH are set!

default: build

build:
	go get -d -v ./...

	cd cmd/healthcheck; \
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go install -a -installsuffix cgo -ldflags '-w -s' healthcheck.go

	cd cmd; \
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go install -a -installsuffix cgo -ldflags '-w -s' gateway.go

container:
	docker build -t $(CURDIR):$(VERSION) .