# STEP 1 build executable binary
FROM golang:alpine as builder

# Install SSL ca certificates
RUN apk update && apk add git && apk add ca-certificates

# Create appuser
RUN adduser -D -g '' appuser

# copy all sources in
COPY . $GOPATH/src
WORKDIR $GOPATH/src/github.com/mtbarta/monocorpus/pkg

# get dependancies
# should be done in CI
RUN go get -d -v ./...

#build the binary
WORKDIR $GOPATH/src/github.com/mtbarta/monocorpus/cmd
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags '-w -s' -o /go/bin/notes notes/notes.go

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -ldflags '-w -s' -o /go/bin/healthcheck healthcheck/healthcheck.go

# STEP 2 build a small image

# start from scratch
FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
# Copy our static executable
COPY --from=builder /go/bin/notes /go/bin/main
COPY --from=builder /go/bin/healthcheck /go/bin/healthcheck
USER appuser

ENTRYPOINT ["/go/bin/main"]