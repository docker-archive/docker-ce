include ../containerd.mk

SHELL:=/bin/bash
ARCH:=$(shell uname -m)
CLI_DIR:=$(CURDIR)/../../cli
GITCOMMIT?=$(shell cd $(CLI_DIR) && git rev-parse --short HEAD)
VERSION?=0.0.0-dev
GO_BASE_IMAGE=golang
GO_VERSION:=1.10.3
GO_IMAGE=$(GO_BASE_IMAGE):$(GO_VERSION)
DEB_VERSION=$(shell ./gen-deb-ver $(CLI_DIR) "$(VERSION)")
CHOWN:=docker run --rm -v $(CURDIR):/v -w /v alpine chown
EPOCH?=2

COMMON_FILES=common
BUILD?=docker build \
	--build-arg GO_IMAGE=$(GO_IMAGE) \
	--build-arg COMMON_FILES=$(COMMON_FILES) \
	-t debbuild-$@/$(ARCH) \
	-f $(CURDIR)/$@/Dockerfile .
RUN=docker run --rm -i \
	-e EPOCH='$(EPOCH)' \
	-e DEB_VERSION=$(word 1, $(DEB_VERSION)) \
	-e VERSION=$(word 2, $(DEB_VERSION)) \
	-e DOCKER_GITCOMMIT=$(GITCOMMIT) \
	-v $(CURDIR)/debbuild/$@:/build \
	debbuild-$@/$(ARCH)

SOURCE_FILES=containerd-proxy.tgz cli.tgz containerd-shim-process.tar docker.service dockerd.json
SOURCES=$(addprefix sources/, $(SOURCE_FILES))

.PHONY: help
help: ## show make targets
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {sub("\\\\n",sprintf("\n%22c"," "), $$2);printf " \033[36m%-20s\033[0m  %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.PHONY: clean
clean: ## remove build artifacts
	[ ! -d debbuild ] || $(CHOWN) -R $(shell id -u):$(shell id -g) debbuild
	$(RM) -r debbuild
	[ ! -d sources ] || $(CHOWN) -R $(shell id -u):$(shell id -g) sources
	$(RM) -r sources
	[ ! -d artifacts ] || $(CHOWN) -R $(shell id -u):$(shell id -g) artifacts
	$(RM) -r artifacts

engine-$(ARCH).tar:
	$(MAKE) -C ../image image-linux
	docker save -o $@ $$(cat ../image/image-linux)

.PHONY: deb
deb: ubuntu debian raspbian ## build all deb packages

.PHONY: ubuntu
ubuntu: ubuntu-bionic ubuntu-xenial ubuntu-trusty ## build all ubuntu deb packages

.PHONY: debian
debian: debian-stretch debian-jessie ## build all debian deb packages

.PHONY: raspbian
raspbian: raspbian-stretch debian-jessie ## build all raspbian deb packages

.PHONY: ubuntu-bionic
ubuntu-bionic: ## build ubuntu bionic deb packages
ubuntu-bionic: $(SOURCES)
	$(BUILD)
	$(RUN)
	$(CHOWN) -R $(shell id -u):$(shell id -g) debbuild/$@

.PHONY: ubuntu-xenial
ubuntu-xenial: ## build ubuntu xenial deb packages
ubuntu-xenial: $(SOURCES)
	$(BUILD)
	$(RUN)
	$(CHOWN) -R $(shell id -u):$(shell id -g) debbuild/$@

.PHONY: ubuntu-trusty
ubuntu-trusty: ## build ubuntu trusty deb packages
ubuntu-trusty: $(SOURCES)
	$(BUILD)
	$(RUN)
	$(CHOWN) -R $(shell id -u):$(shell id -g) debbuild/$@

.PHONY: debian-buster
debian-buster: ## build debian buster deb packages
debian-buster: $(SOURCES)
	$(BUILD)
	$(RUN)
	$(CHOWN) -R $(shell id -u):$(shell id -g) debbuild/$@

.PHONY: debian-jessie
debian-jessie: ## build debian jessie deb packages
debian-jessie: $(SOURCES)
	$(BUILD)
	$(RUN)
	$(CHOWN) -R $(shell id -u):$(shell id -g) debbuild/$@

.PHONY: debian-stretch
debian-stretch: ## build debian stretch deb packages
debian-stretch: $(SOURCES)
	$(BUILD)
	$(RUN)
	$(CHOWN) -R $(shell id -u):$(shell id -g) debbuild/$@

.PHONY: raspbian-jessie
raspbian-jessie: ## build raspbian jessie deb packages
raspbian-jessie: $(SOURCES)
	$(BUILD)
	$(RUN)
	$(CHOWN) -R $(shell id -u):$(shell id -g) debbuild/$@

.PHONY: raspbian-stretch
raspbian-stretch: ## build raspbian stretch deb packages
raspbian-stretch: $(SOURCES)
	$(BUILD)
	$(RUN)
	$(CHOWN) -R $(shell id -u):$(shell id -g) debbuild/$@

sources/cli.tgz:
	mkdir -p $(@D)
	docker run --rm -i -w /v \
		-v $(CLI_DIR):/cli \
		-v $(CURDIR)/$(@D):/v \
		alpine \
		tar -C / -c -z -f /v/cli.tgz --exclude .git cli

sources/containerd-proxy.tgz:
	mkdir -p tmp/
	curl -fL -o tmp/containerd-proxy.tgz "https://github.com/crosbymichael/containerd-proxy/archive/$(CONTAINERD_PROXY_COMMIT)/ours.tar.gz"
	tar xzf tmp/containerd-proxy.tgz -C tmp/
	mv tmp/containerd-proxy-$(CONTAINERD_PROXY_COMMIT) tmp/containerd-proxy
	mkdir -p $(@D)
	$(CHOWN) -R $(shell id -u):$(shell id -g) $$(dirname $(@D))
	tar -zcf $@ -C tmp/ containerd-proxy
	rm -rf tmp/

sources/containerd-shim-process.tar:
	$(CTR) content fetch $(CONTAINERD_SHIM_PROCESS_IMAGE)
	$(CTR) image export artifacts/containerd-shim-process.tar $(CONTAINERD_SHIM_PROCESS_IMAGE)
	mkdir -p $(@D)
	cp artifacts/containerd-shim-process.tar $@
	$(CHOWN) -R $(shell id -u):$(shell id -g) $$(dirname $(@D))

sources/docker.service: ../systemd/docker.service
	mkdir -p $(@D)
	cp $< $@

sources/dockerd.json: ../common/dockerd.json
	mkdir -p $(@D)
	cp $< $@

# TODO: Figure out a sufficient offline solution
