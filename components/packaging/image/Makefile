SHELL:=/bin/bash
ENGINE_DIR:=$(CURDIR)/../../engine
CLI_DIR:=$(CURDIR)/../../cli
VERSION?=0.0.0-dev
STATIC_VERSION=$(shell ../static/gen-static-ver $(ENGINE_DIR) $(VERSION))
DOCKER_HUB_ORG?=docker
ENGINE_IMAGE?=ce-engine

.PHONY: help
help: ## show make targets
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {sub("\\\\n",sprintf("\n%22c"," "), $$2);printf " \033[36m%-20s\033[0m  %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.PHONY: clean
clean: ## remove build artifacts
	docker rmi $(DOCKER_HUB_ORG)/$(ENGINE_IMAGE):$(STATIC_VERSION)

.PHONY: image
image: image-linux

.PHONY: image-linux
image-linux:
	docker build -t $(DOCKER_HUB_ORG)/$(ENGINE_IMAGE):$(STATIC_VERSION) \
	    --build-arg VERSION=$(STATIC_VERSION) \
	    --build-arg DOCKER_GITCOMMIT=$(cd $(ENGINE_DIR) && git rev-parse --short=7 HEAD) \
	    --file ./Dockerfile $(ENGINE_DIR)

.PHONY: release
release:
	docker push $(DOCKER_HUB_ORG)/$(ENGINE_IMAGE):$(STATIC_VERSION)
