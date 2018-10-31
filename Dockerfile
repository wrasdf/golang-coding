FROM golang:1.11-alpine3.8

RUN apk add --no-cache git bash make curl jq
RUN go get -u github.com/golang/dep/cmd/dep

WORKDIR /go/src/github.com/wrasdf/golang-coding
COPY Gopkg.* /go/src/github.com/wrasdf/golang-coding/
RUN dep ensure -v -vendor-only
