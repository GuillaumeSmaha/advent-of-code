.PHONY: build

NAME=/advent-of-code-2017

build:
	docker run --rm -v $(PWD):$(NAME) -w $(NAME) golang:latest /bin/bash -c "OS=$(shell uname | tr A-Z a-z) make _build"

_build:
	go test -v ./...
	GOOS=$(OS) go build
