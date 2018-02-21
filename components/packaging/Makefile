SHELL:=/bin/bash
ENGINE_DIR:=$(CURDIR)/../engine
CLI_DIR:=$(CURDIR)/../cli
VERSION?=0.0.0-dev
DOCKER_GITCOMMIT:=abcdefg

.PHONY: help
help: ## show make targets
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {sub("\\\\n",sprintf("\n%22c"," "), $$2);printf " \033[36m%-20s\033[0m  %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.PHONY: clean
clean: ## remove build artifacts
	$(MAKE) -C rpm clean
	$(MAKE) -C deb clean
	$(MAKE) -C static clean

.PHONY: rpm
rpm: DOCKER_BUILD_PKGS:=fedora-27 fedora-26 centos-7
rpm: ## build rpm packages
	for p in $(DOCKER_BUILD_PKGS); do \
		$(MAKE) -C $@ VERSION=$(VERSION) ENGINE_DIR=$(ENGINE_DIR) CLI_DIR=$(CLI_DIR) $${p}; \
	done

.PHONY: deb
deb: DOCKER_BUILD_PKGS:=ubuntu-xenial ubuntu-trusty debian-buster debian-stretch debian-wheezy debian-jessie raspbian-stretch raspbian-jessie
deb: ## build deb packages
	for p in $(DOCKER_BUILD_PKGS); do \
		$(MAKE) -C $@ VERSION=$(VERSION) ENGINE_DIR=$(ENGINE_DIR) CLI_DIR=$(CLI_DIR) $${p}; \
	done

.PHONY: static
static: DOCKER_BUILD_PKGS:=static-linux cross-mac cross-win cross-arm
static: ## build static-compiled packages
	for p in $(DOCKER_BUILD_PKGS); do \
		$(MAKE) -C $@ VERSION=$(VERSION) ENGINE_DIR=$(ENGINE_DIR) CLI_DIR=$(CLI_DIR) $${p}; \
	done
