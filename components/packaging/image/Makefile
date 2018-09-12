SHELL:=/bin/bash
ENGINE_DIR:=$(CURDIR)/../../engine
CLI_DIR:=$(CURDIR)/../../cli
VERSION?=0.0.0-dev
GO_BASE_IMAGE=golang
GO_VERSION:=1.10.4
ENGINE_GO_IMAGE=$(GO_BASE_IMAGE):$(GO_VERSION)
STATIC_VERSION=$(shell ../static/gen-static-ver $(ENGINE_DIR) $(VERSION))
DOCKER_HUB_ORG?=dockereng
ARCH=$(shell uname -m)
ENGINE_IMAGE?=engine-community
CHOWN:=docker run --rm -v $(CURDIR):/v -w /v alpine chown
DEFAULT_PRODUCT_LICENSE?=Community Engine

.PHONY: help
help: ## show make targets
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {sub("\\\\n",sprintf("\n%22c"," "), $$2);printf " \033[36m%-20s\033[0m  %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.PHONY: clean
clean: ## remove build artifacts
	-$(RM) $(ENGINE_DIR)/Dockerfile.engine
	-docker rmi $(DOCKER_HUB_ORG)/$(ENGINE_IMAGE):$(STATIC_VERSION)
	-rm -f image-linux
	-$(RM) -r artifacts
	-$(RM) *.tar

.PHONY: image
image: image-linux


$(ENGINE_DIR)/Dockerfile.engine:
	cp Dockerfile.engine $(ENGINE_DIR)

# builds across multiple archs because the base images
# utilize manifests
image-linux: $(ENGINE_DIR)/Dockerfile.engine
	docker build -t $(DOCKER_HUB_ORG)/$(ENGINE_IMAGE):$(STATIC_VERSION).$(ARCH) \
		--build-arg GO_IMAGE="$(ENGINE_GO_IMAGE)" \
		--build-arg VERSION="$(STATIC_VERSION)" \
		--build-arg GITCOMMIT="$$(cd $(ENGINE_DIR) && git rev-parse --short=7 HEAD)" \
		--build-arg BUILDTIME="$(BUILDTIME)" \
		--build-arg PLATFORM="$(PLATFORM)" \
		--build-arg PRODUCT="$(PRODUCT)" \
		--build-arg DEFAULT_PRODUCT_LICENSE="$(DEFAULT_PRODUCT_LICENSE)" \
		--file $< $(ENGINE_DIR)
	echo $(DOCKER_HUB_ORG)/$(ENGINE_IMAGE):$(STATIC_VERSION).$(ARCH) > $@

DOCKER2OCI=artifacts/docker2oci
$(DOCKER2OCI):
	-$(CHOWN) -R $(shell id -u):$(shell id -g) $(@D)
	docker run --name docker2oci $(ENGINE_GO_IMAGE) sh -c 'go get github.com/coolljt0725/docker2oci'
	mkdir -p $(@D)
	docker cp docker2oci:/go/bin/docker2oci "$@"
	docker rm -f docker2oci
	$(CHOWN) -R $(shell id -u):$(shell id -g) $(@D)

engine-$(ARCH).tar: engine-$(ARCH)-docker-compat.tar $(DOCKER2OCI)
	mkdir -p artifacts
	./$(DOCKER2OCI) -i $< artifacts/engine-image
	mkdir -p $(@D)
	tar c -C artifacts/engine-image . > $@

engine-$(ARCH)-docker-compat.tar: image-linux
	docker save -o $@ $$(cat $<)

.PHONY: release
release:
	docker push $(DOCKER_HUB_ORG)/$(ENGINE_IMAGE):$(STATIC_VERSION).$(ARCH)
