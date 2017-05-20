SHELL:=/bin/bash
ARCH:=amd64
ENGINE_DIR:=$(CURDIR)/../../engine
CLI_DIR:=$(CURDIR)/../../cli
GITCOMMIT?=$(shell cd $(ENGINE_DIR) && git rev-parse --short HEAD)
VERSION?=$(shell cat $(ENGINE_DIR)/VERSION)
DOCKER_EXPERIMENTAL:=0
CHOWN:=docker run --rm -v $(CURDIR):/v -w /v alpine chown

.PHONY: help clean deb ubuntu debian ubuntu-xenial ubuntu-trusty ubuntu-yakkety ubuntu-zesty debian-jessie debian-stretch debian-wheezy

help: ## show make targets
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {sub("\\\\n",sprintf("\n%22c"," "), $$2);printf " \033[36m%-20s\033[0m  %s\n", $$1, $$2}' $(MAKEFILE_LIST)

clean: ## remove build artifacts
	[ ! -d debbuild ] || $(CHOWN) -R $(shell id -u):$(shell id -g) debbuild
	$(RM) -r debbuild

deb: ubuntu debian ## build all deb packages

ubuntu: ubuntu-zesty ubuntu-yakkety ubuntu-xenial ubuntu-trusty ## build all ubuntu deb packages

debian: debian-stretch debian-wheezy debian-jessie ## build all debian deb packages

ubuntu-xenial: ## build ubuntu xenial deb packages
	docker build -t debbuild-$@/$(ARCH) -f $(CURDIR)/$@/Dockerfile.$(ARCH) .
	docker run --rm \
		-e VERSION=$(VERSION) \
		-e DOCKER_GITCOMMIT=$(GITCOMMIT) \
		-v $(CURDIR)/debbuild/$@:/build \
		-v $(ENGINE_DIR):/engine \
		-v $(CLI_DIR):/cli \
		debbuild-$@/$(ARCH)
	$(CHOWN) -R $(shell id -u):$(shell id -g) debbuild/$@

ubuntu-trusty: ## build ubuntu trusty deb packages
	docker build -t debbuild-$@/$(ARCH) -f $(CURDIR)/$@/Dockerfile.$(ARCH) .
	docker run --rm \
		-e VERSION=$(VERSION) \
		-e DOCKER_GITCOMMIT=$(GITCOMMIT) \
		-v $(CURDIR)/debbuild/$@:/build \
		-v $(ENGINE_DIR):/engine \
		-v $(CLI_DIR):/cli \
		debbuild-$@/$(ARCH)
	$(CHOWN) -R $(shell id -u):$(shell id -g) debbuild/$@

ubuntu-yakkety: ## build ubuntu yakkety deb packages
	docker build -t debbuild-$@/$(ARCH) -f $(CURDIR)/$@/Dockerfile.$(ARCH) .
	docker run --rm \
		-v $(CURDIR)/debbuild/$@:/build \
		-v $(ENGINE_DIR):/engine \
		-v $(CLI_DIR):/cli \
		debbuild-$@/$(ARCH)
	$(CHOWN) -R $(shell id -u):$(shell id -g) debbuild/$@

ubuntu-zesty: ## build ubuntu zesty deb packages
	docker build -t debbuild-$@/$(ARCH) -f $(CURDIR)/$@/Dockerfile.$(ARCH) .
	docker run --rm \
		-v $(CURDIR)/debbuild/$@:/build \
		-v $(ENGINE_DIR):/engine \
		-v $(CLI_DIR):/cli \
		debbuild-$@/$(ARCH)
	$(CHOWN) -R $(shell id -u):$(shell id -g) debbuild/$@

debian-jessie: ## build debian jessie deb packages
	docker build -t debbuild-$@/$(ARCH) -f $(CURDIR)/$@/Dockerfile.$(ARCH) .
	docker run --rm \
		-v $(CURDIR)/debbuild/$@:/build \
		-v $(ENGINE_DIR):/engine \
		-v $(CLI_DIR):/cli \
		debbuild-$@/$(ARCH)
	$(CHOWN) -R $(shell id -u):$(shell id -g) debbuild/$@

debian-stretch: ## build debian stretch deb packages
	docker build -t debbuild-$@/$(ARCH) -f $(CURDIR)/$@/Dockerfile.$(ARCH) .
	docker run --rm \
		-v $(CURDIR)/debbuild/$@:/build \
		-v $(ENGINE_DIR):/engine \
		-v $(CLI_DIR):/cli \
		debbuild-$@/$(ARCH)
	$(CHOWN) -R $(shell id -u):$(shell id -g) debbuild/$@

debian-wheezy: ## build debian wheezy deb packages
	docker build -t debbuild-$@/$(ARCH) -f $(CURDIR)/$@/Dockerfile.$(ARCH) .
	docker run --rm \
		-v $(CURDIR)/debbuild/$@:/build \
		-v $(ENGINE_DIR):/engine \
		-v $(CLI_DIR):/cli \
		debbuild-$@/$(ARCH)
	$(CHOWN) -R $(shell id -u):$(shell id -g) debbuild/$@
