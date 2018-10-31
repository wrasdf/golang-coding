FROM golang:1.11-alpine3.8

RUN apk add --no-cache git bash make curl jq
RUN go get -u github.com/golang/dep/cmd/dep
