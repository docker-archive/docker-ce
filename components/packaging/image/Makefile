SHELL:=/bin/bash
ENGINE_DIR:=$(CURDIR)/../../engine
CLI_DIR:=$(CURDIR)/../../cli
VERSION?=0.0.0-dev
GO_BASE_IMAGE=golang
GO_VERSION:=1.10.4
GO_IMAGE=$(GO_BASE_IMAGE):$(GO_VERSION)
STATIC_VERSION=$(shell ../static/gen-static-ver $(ENGINE_DIR) $(VERSION))
DOCKER_HUB_ORG?=dockereng
ARCH=$(shell uname -m)
ENGINE_IMAGE?=engine-community
DEFAULT_PRODUCT_LICENSE?=Community Engine

.PHONY: help
help: ## show make targets
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {sub("\\\\n",sprintf("\n%22c"," "), $$2);printf " \033[36m%-20s\033[0m  %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.PHONY: clean
clean: ## remove build artifacts
	-$(RM) $(ENGINE_DIR)/Dockerfile.engine
	-docker rmi $(DOCKER_HUB_ORG)/$(ENGINE_IMAGE):$(STATIC_VERSION)
	-rm -f image-linux

.PHONY: image
image: image-linux


$(ENGINE_DIR)/Dockerfile.engine:
	cp Dockerfile.engine $(ENGINE_DIR)

# builds across multiple archs because the base images
# utilize manifests
image-linux: $(ENGINE_DIR)/Dockerfile.engine
	docker build -t $(DOCKER_HUB_ORG)/$(ENGINE_IMAGE):$(STATIC_VERSION).$(ARCH) \
		--build-arg GO_IMAGE="$(GO_IMAGE)" \
		--build-arg VERSION="$(STATIC_VERSION)" \
		--build-arg GITCOMMIT="$$(cd $(ENGINE_DIR) && git rev-parse --short=7 HEAD)" \
		--build-arg BUILDTIME="$(BUILDTIME)" \
		--build-arg PLATFORM="$(PLATFORM)" \
		--build-arg PRODUCT="$(PRODUCT)" \
		--build-arg DEFAULT_PRODUCT_LICENSE="$(DEFAULT_PRODUCT_LICENSE)" \
		--file $< $(ENGINE_DIR)
	echo $(DOCKER_HUB_ORG)/$(ENGINE_IMAGE):$(STATIC_VERSION).$(ARCH) > $@

engine-$(ARCH).tar: image-linux
	docker save -o $@ $$(cat $<)


.PHONY: release
release:
	docker push $(DOCKER_HUB_ORG)/$(ENGINE_IMAGE):$(STATIC_VERSION).$(ARCH)
