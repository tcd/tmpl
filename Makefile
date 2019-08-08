.DEFAULT_GOAL := help

SHELL := /bin/bash
GOBIN_DIR=${GOBIN}
PROJECT_DIR=$(shell pwd)
PROJECT_NAME=$(shell basename $(PROJECT_DIR))

go-uninstall:
	@rm -f $(GOBIN_DIR)/$(PROJECT_NAME)

all:
	@echo "Nothing to do for all"

build:
	@echo "Nothing to do for build"

clean:
	go clean ./...

test:
	go test -v ./...

install:
	@echo "Nothing to do for install"

uninstall: go-uninstall

help:
	@echo
	@echo "  clean – clean all files built by 'go build'"
	@echo "  test  – run 'go test' for the entire project"
	@echo

.PHONY: all build clean test help install uninstall cmd
