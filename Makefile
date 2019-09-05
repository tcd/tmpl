GO ?= go
SHELL := /bin/sh
GOBIN_DIR=${GOBIN}
PROJECT_DIR=$(shell pwd)
PROJECT_NAME=$(shell basename $(PROJECT_DIR))

mod:
	GO111MODULE=on go mod tidy

clean:
	$(GO) clean ./...
	rm -rf build

test:
	go test -v ./...

install:
	GO111MODULE=on go install

uninstall:
	rm -f $(GOBIN_DIR)/$(PROJECT_NAME)

.PHONY: clean test install uninstall
