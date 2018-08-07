SHELL:=/bin/bash
ARCH:=$(shell uname -m)
CLI_DIR:=$(CURDIR)/../../cli
GITCOMMIT?=$(shell cd $(CLI_DIR) && git rev-parse --short HEAD)
VERSION?=0.0.0-dev
GO_VERSION:=1.10.3
DEB_VERSION=$(shell ./gen-deb-ver $(CLI_DIR) "$(VERSION)")
CHOWN:=docker run --rm -v $(CURDIR):/v -w /v alpine chown
EPOCH?=2

BUILD=docker build --build-arg GO_VERSION=$(GO_VERSION) -t debbuild-$@/$(ARCH) -f $(CURDIR)/$@/Dockerfile.$(ARCH) .
RUN=docker run --rm -i \
	-e EPOCH='$(EPOCH)' \
	-e DEB_VERSION=$(word 1, $(DEB_VERSION)) \
	-e VERSION=$(word 2, $(DEB_VERSION)) \
	-e DOCKER_GITCOMMIT=$(GITCOMMIT) \
	-v $(CURDIR)/debbuild/$@:/build \
	-v $(CLI_DIR):/cli \
	-v $(CURDIR)/systemd:/root/build-deb/systemd \
	debbuild-$@/$(ARCH)

.PHONY: help
help: ## show make targets
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {sub("\\\\n",sprintf("\n%22c"," "), $$2);printf " \033[36m%-20s\033[0m  %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.PHONY: clean
clean: ## remove build artifacts
	[ ! -d debbuild ] || $(CHOWN) -R $(shell id -u):$(shell id -g) debbuild
	$(RM) -r debbuild

engine-$(ARCH).tar:
	$(MAKE) -C ../image image-linux
	docker save -o $@ $(shell cat ../image/image-linux)

.PHONY: deb
deb: ubuntu debian raspbian ## build all deb packages

.PHONY: ubuntu
ubuntu: ubuntu-bionic ubuntu-xenial ubuntu-trusty ## build all ubuntu deb packages

.PHONY: debian
debian: debian-stretch debian-jessie ## build all debian deb packages

.PHONY: raspbian
raspbian: raspbian-stretch debian-jessie ## build all raspbian deb packages

.PHONY: ubuntu-xenial
ubuntu-xenial: engine-$(ARCH).tar ## build ubuntu xenial deb packages
	$(BUILD)
	$(RUN)
	$(CHOWN) -R $(shell id -u):$(shell id -g) debbuild/$@

.PHONY: ubuntu-trusty
ubuntu-trusty: engine-$(ARCH).tar ## build ubuntu trusty deb packages
	$(BUILD)
	$(RUN)
	$(CHOWN) -R $(shell id -u):$(shell id -g) debbuild/$@

.PHONY: ubuntu-bionic
ubuntu-bionic: engine-$(ARCH).tar ## build ubuntu bionic deb packages
	$(BUILD)
	$(RUN)
	$(CHOWN) -R $(shell id -u):$(shell id -g) debbuild/$@

.PHONY: debian-buster
debian-buster: engine-$(ARCH).tar ## build debian buster deb packages
	$(BUILD)
	$(RUN)
	$(CHOWN) -R $(shell id -u):$(shell id -g) debbuild/$@

.PHONY: debian-jessie
debian-jessie: engine-$(ARCH).tar ## build debian jessie deb packages
	$(BUILD)
	$(RUN)
	$(CHOWN) -R $(shell id -u):$(shell id -g) debbuild/$@

.PHONY: debian-stretch
debian-stretch: engine-$(ARCH).tar ## build debian stretch deb packages
	$(BUILD)
	$(RUN)
	$(CHOWN) -R $(shell id -u):$(shell id -g) debbuild/$@

.PHONY: raspbian-jessie
raspbian-jessie: engine-$(ARCH).tar ## build raspbian jessie deb packages
	$(BUILD)
	$(RUN)
	$(CHOWN) -R $(shell id -u):$(shell id -g) debbuild/$@

.PHONY: raspbian-stretch
raspbian-stretch: engine-$(ARCH).tar ## build raspbian stretch deb packages
	$(BUILD)
	$(RUN)
	$(CHOWN) -R $(shell id -u):$(shell id -g) debbuild/$@
