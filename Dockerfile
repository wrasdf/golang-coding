# Build the manager binary
FROM golang:1.16.6 as builder
WORKDIR /go/src/github.com/wrasdf/golang-coding/

COPY go.* /go/src/github.com/wrasdf/golang-coding/
RUN go mod download -x

# Copy the go source
COPY . /go/src/github.com/wrasdf/golang-coding/

# Build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o ./bin/code main.go

# Release binary
FROM alpine:3.14 as release
RUN apk --update add --no-cache bash \
  && rm -rf /var/cache/apk/*

WORKDIR /
COPY --from=Builder /go/src/github.com/wrasdf/golang-coding/bin/code /

ENTRYPOINT ["/code"]