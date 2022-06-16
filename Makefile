# Commands
CURRENT_DIR = $(shell pwd)
BUILD_DIR=../build/
BIN_NAME=tf-compose
DOCKER_RUN?=docker run --rm -it -v "$(CURRENT_DIR)":/usr/src/myapp -w /usr/src/myapp 
DOCKER_RUN_GO?=$(DOCKER_RUN) --mount source=gopath,target=/go golang:1.18

## Help command
.DEFAULT_GOAL := help
.PHONY: help
help: ## Display commands list
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m\033[0m\n"} /^[$$()% 0-9a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

## Local stack:
bash: ## Launch a bash
	$(DOCKER_RUN_GO) bash

run: ## Build application for all platforms
	$(DOCKER_RUN_GO) bash -c "cd src && go run *.go"

compile: ## Build application for all platforms
	$(DOCKER_RUN_GO) bash ./scripts/compile.sh

install: ## Build application for all platforms
	$(DOCKER_RUN_GO) go init
