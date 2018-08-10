SHELL:=/bin/bash
ENGINE_DIR:=$(CURDIR)/../engine
CLI_DIR:=$(CURDIR)/../cli
VERSION?=0.0.0-dev
DOCKER_GITCOMMIT:=abcdefg
ARCH=$(shell uname -m)
STATIC_VERSION=$(shell static/gen-static-ver $(ENGINE_DIR) $(VERSION))
GO_VERSION:=1.10.3

# Taken from: https://www.cmcrossroads.com/article/printing-value-makefile-variable
print-%  : ; @echo $($*)

.PHONY: help
help: ## show make targets
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {sub("\\\\n",sprintf("\n%22c"," "), $$2);printf " \033[36m%-20s\033[0m  %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.PHONY: clean-engine
clean-engine:
	rm -rf $(ENGINE_DIR)

.PHONY: clean-image
clean-image:
	$(MAKE) ENGINE_DIR=$(ENGINE_DIR) -C image clean


.PHONY: clean
clean: clean-image ## remove build artifacts
	$(MAKE) -C rpm clean
	$(MAKE) -C deb clean
	$(MAKE) -C static clean

.PHONY: rpm
rpm: DOCKER_BUILD_PKGS:=$(shell find rpm -type d | grep ".*-.*" | sed 's/^rpm\///')
rpm: ## build rpm packages
	for p in $(DOCKER_BUILD_PKGS); do \
		$(MAKE) -C $@ VERSION=$(VERSION) ENGINE_DIR=$(ENGINE_DIR) CLI_DIR=$(CLI_DIR) GO_VERSION=$(GO_VERSION) $${p}; \
	done

.PHONY: deb
deb: DOCKER_BUILD_PKGS:=$(shell find deb -type d | grep ".*-.*" | sed 's/^deb\///')
deb: ## build deb packages
	for p in $(DOCKER_BUILD_PKGS); do \
		$(MAKE) -C $@ VERSION=$(VERSION) ENGINE_DIR=$(ENGINE_DIR) CLI_DIR=$(CLI_DIR) GO_VERSION=$(GO_VERSION) $${p}; \
	done

.PHONY: static
static: DOCKER_BUILD_PKGS:=static-linux cross-mac cross-win cross-arm
static: ## build static-compiled packages
	for p in $(DOCKER_BUILD_PKGS); do \
		$(MAKE) -C $@ VERSION=$(VERSION) ENGINE_DIR=$(ENGINE_DIR) CLI_DIR=$(CLI_DIR) GO_VERSION=$(GO_VERSION) $${p}; \
	done

# TODO - figure out multi-arch
.PHONY: image
image: DOCKER_BUILD_PKGS:=image-linux
image: ## build static-compiled packages
	for p in $(DOCKER_BUILD_PKGS); do \
		$(MAKE) -C $@ VERSION=$(VERSION) ENGINE_DIR=$(ENGINE_DIR) CLI_DIR=$(CLI_DIR) GO_VERSION=$(GO_VERSION) $${p}; \
	done

engine-$(ARCH).tar:
	$(MAKE) -C image $@

.PHONY: release
release:
	$(MAKE) -C image $@
