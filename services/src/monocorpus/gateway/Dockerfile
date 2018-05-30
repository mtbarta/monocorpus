# STEP 1 build executable binary
FROM golang:alpine as builder

# Install SSL ca certificates
RUN apk update && apk add git && apk add ca-certificates

# Create appuser
RUN adduser -D -g '' appuser

# copy all sources in
COPY . $GOPATH/src/monocorpus/gateway
#super hack to avoid knowing the CI token
RUN mv $GOPATH/src/monocorpus/gateway/notes $GOPATH/src/monocorpus/notes
WORKDIR $GOPATH/src/monocorpus/gateway

# get dependancies
# should be done in CI
RUN go get -d -v ./...

#build the binary
WORKDIR $GOPATH/src/monocorpus/gateway/cmd
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags '-w -s' -o /go/bin/main main.go

WORKDIR $GOPATH/src/monocorpus/gateway/cmd/healthcheck
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags '-w -s' -o /go/bin/healthcheck healthcheck.go

# STEP 2 build a small image

# start from scratch
FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
# Copy our static executable
COPY --from=builder /go/bin/main /go/bin/main
COPY --from=builder /go/bin/healthcheck /go/bin/healthcheck
USER appuser

ENTRYPOINT ["/go/bin/main"]