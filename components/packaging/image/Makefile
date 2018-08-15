SHELL:=/bin/bash
ENGINE_DIR:=$(CURDIR)/../../engine
CLI_DIR:=$(CURDIR)/../../cli
VERSION?=0.0.0-dev
STATIC_VERSION=$(shell ../static/gen-static-ver $(ENGINE_DIR) $(VERSION))
DOCKER_HUB_ORG?=dockereng
ARCH=$(shell uname -m)
ENGINE_IMAGE?=engine-community

.PHONY: help
help: ## show make targets
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {sub("\\\\n",sprintf("\n%22c"," "), $$2);printf " \033[36m%-20s\033[0m  %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.PHONY: clean
clean: ## remove build artifacts
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
	    --build-arg VERSION=$(STATIC_VERSION) \
	    --build-arg DOCKER_GITCOMMIT=$$(cd $(ENGINE_DIR) && git rev-parse --short=7 HEAD) \
	    --file $< $(ENGINE_DIR)
	echo $(DOCKER_HUB_ORG)/$(ENGINE_IMAGE):$(STATIC_VERSION).$(ARCH) > $@

engine-$(ARCH).tar: image-linux
	docker save -o $@ $$(cat $<)


.PHONY: release
release:
	docker push $(DOCKER_HUB_ORG)/$(ENGINE_IMAGE):$(STATIC_VERSION).$(ARCH)
