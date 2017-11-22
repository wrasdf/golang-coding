Build_IMAGE=go-glide:1.9

console:
	@docker build -t $(Build_IMAGE) .
	@docker run -it -v $$(pwd):/go/src/golang-coding -w /go/src/golang-coding $(Build_IMAGE) /bin/bash

build:
	@go build -v

clean:
	@go clean -i -n
