.PHONY: build

NAME=/advent-of-code-2018

build:
	docker run --rm -v $(PWD):$(NAME) -w $(NAME) golang:latest /bin/bash -c "OS=$(shell uname | tr A-Z a-z) make _build"

_build:
	GOOS=$(OS) go build

test:
	docker run --rm -v $(PWD):$(NAME) -w $(NAME) golang:latest /bin/bash -c "OS=$(shell uname | tr A-Z a-z) make _test"
	
_test:
	go test -v ./...
	GOOS=$(OS) go build

