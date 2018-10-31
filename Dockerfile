FROM golang:1.11-alpine3.8

WORKDIR /go/src/github.com/wrasdf/golang-coding
RUN apk add --no-cache gcc musl-dev libffi-dev linux-headers openssl-dev git bash make curl jq
RUN go get -u github.com/golang/dep/cmd/dep github.com/onsi/ginkgo/ginkgo

COPY Gopkg.* /go/src/github.com/wrasdf/golang-coding/
RUN dep ensure --vendor-only
