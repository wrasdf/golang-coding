FROM golang:1.13.10-alpine3.11

ENV GO111MODULE=on
WORKDIR /go/src/app/
RUN apk add --no-cache gcc musl-dev ca-certificates git bash make curl jq
COPY go.mod go.sum ./
RUN go mod download

COPY . ./
