IMAGE=operator:latest
PROJECT=kube-go-operator

sh:
	@docker build -t $(IMAGE) .
	@docker run -it  -w /go/src/$(PROJECT) \
	 -v $$(pwd):/go/src/$(PROJECT) \
	 -v $(HOME)/.kube:/root/.kube \
	 -v $(HOME)/.aws:/root/.aws \
	 $(IMAGE) /bin/bash

build:
	@go build -v

clean:
	@go clean -i -n
