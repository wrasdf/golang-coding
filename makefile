IMAGE=coding:latest
PROJECT=golang-coding

sh:
	@docker build -t $(IMAGE) .
	@docker run -it  -w /go/src/$(PROJECT) \
	 -v $$(pwd):/go/src/$(PROJECT) \
	 -v $(HOME)/.kube:/root/.kube \
	 -v $(HOME)/.aws:/root/.aws \
	 $(IMAGE) /bin/bash

run:
	@go run main.go

build:
	@go build -v

clean:
	@go clean -i -n

test:
	@go test
