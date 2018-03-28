SHELL:=/bin/bash
ARCH:=$(shell uname -m)
ENGINE_DIR:=$(CURDIR)/../../engine
CLI_DIR:=$(CURDIR)/../../cli
GITCOMMIT?=$(shell cd $(ENGINE_DIR) && git rev-parse --short HEAD)
VERSION?=0.0.0-dev
DEB_VERSION=$(shell ./gen-deb-ver $(ENGINE_DIR) $(VERSION))
DOCKER_EXPERIMENTAL:=0
CHOWN:=docker run --rm -v $(CURDIR):/v -w /v alpine chown

.PHONY: help
help: ## show make targets
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {sub("\\\\n",sprintf("\n%22c"," "), $$2);printf " \033[36m%-20s\033[0m  %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.PHONY: clean
clean: ## remove build artifacts
	[ ! -d debbuild ] || $(CHOWN) -R $(shell id -u):$(shell id -g) debbuild
	$(RM) -r debbuild

.PHONY: deb
deb: ubuntu debian raspbian ## build all deb packages

.PHONY: ubuntu
ubuntu: ubuntu-bionic ubuntu-artful ubuntu-xenial ubuntu-trusty ## build all ubuntu deb packages

.PHONY: debian
debian: debian-stretch debian-wheezy debian-jessie ## build all debian deb packages

.PHONY: raspbian
raspbian: raspbian-stretch debian-jessie ## build all raspbian deb packages

.PHONY: ubuntu-xenial
ubuntu-xenial: ## build ubuntu xenial deb packages
	docker build -t debbuild-$@/$(ARCH) -f $(CURDIR)/$@/Dockerfile.$(ARCH) .
	docker run --rm -i \
		-e DEB_VERSION=$(DEB_VERSION) \
		-e VERSION=$(VERSION) \
		-e DOCKER_GITCOMMIT=$(GITCOMMIT) \
		-v $(CURDIR)/debbuild/$@:/build \
		-v $(ENGINE_DIR):/engine \
		-v $(CLI_DIR):/cli \
		-v $(CURDIR)/systemd:/root/build-deb/systemd \
		debbuild-$@/$(ARCH)
	$(CHOWN) -R $(shell id -u):$(shell id -g) debbuild/$@

.PHONY: ubuntu-trusty
ubuntu-trusty: ## build ubuntu trusty deb packages
	docker build -t debbuild-$@/$(ARCH) -f $(CURDIR)/$@/Dockerfile.$(ARCH) .
	docker run --rm -i \
		-e DEB_VERSION=$(DEB_VERSION) \
		-e VERSION=$(VERSION) \
		-e DOCKER_GITCOMMIT=$(GITCOMMIT) \
		-v $(CURDIR)/debbuild/$@:/build \
		-v $(ENGINE_DIR):/engine \
		-v $(CLI_DIR):/cli \
		-v $(CURDIR)/systemd:/root/build-deb/systemd \
		debbuild-$@/$(ARCH)
	$(CHOWN) -R $(shell id -u):$(shell id -g) debbuild/$@

.PHONY: ubuntu-artful
ubuntu-artful: ## build ubuntu artful deb packages
	docker build -t debbuild-$@/$(ARCH) -f $(CURDIR)/$@/Dockerfile.$(ARCH) .
	docker run --rm -i \
		-e DEB_VERSION=$(DEB_VERSION) \
		-e VERSION=$(VERSION) \
		-e DOCKER_GITCOMMIT=$(GITCOMMIT) \
		-v $(CURDIR)/debbuild/$@:/build \
		-v $(ENGINE_DIR):/engine \
		-v $(CLI_DIR):/cli \
		-v $(CURDIR)/systemd:/root/build-deb/systemd \
		debbuild-$@/$(ARCH)
	$(CHOWN) -R $(shell id -u):$(shell id -g) debbuild/$@

.PHONY: ubuntu-bionic
ubuntu-bionic: ## build ubuntu bionic deb packages
	$(BUILD)
	$(RUN)
	$(CHOWN) -R $(shell id -u):$(shell id -g) debbuild/$@

.PHONY: debian-buster
debian-buster: ## build debian buster deb packages
	docker build -t debbuild-$@/$(ARCH) -f $(CURDIR)/$@/Dockerfile.$(ARCH) .
	docker run --rm -i \
		-e DEB_VERSION=$(DEB_VERSION) \
		-e VERSION=$(VERSION) \
		-e DOCKER_GITCOMMIT=$(GITCOMMIT) \
		-v $(CURDIR)/debbuild/$@:/build \
		-v $(ENGINE_DIR):/engine \
		-v $(CLI_DIR):/cli \
		-v $(CURDIR)/systemd:/root/build-deb/systemd \
		debbuild-$@/$(ARCH)
	$(CHOWN) -R $(shell id -u):$(shell id -g) debbuild/$@

.PHONY: debian-jessie
debian-jessie: ## build debian jessie deb packages
	docker build -t debbuild-$@/$(ARCH) -f $(CURDIR)/$@/Dockerfile.$(ARCH) .
	docker run --rm -i \
		-e DEB_VERSION=$(DEB_VERSION) \
		-e VERSION=$(VERSION) \
		-e DOCKER_GITCOMMIT=$(GITCOMMIT) \
		-v $(CURDIR)/debbuild/$@:/build \
		-v $(ENGINE_DIR):/engine \
		-v $(CLI_DIR):/cli \
		-v $(CURDIR)/systemd:/root/build-deb/systemd \
		debbuild-$@/$(ARCH)
	$(CHOWN) -R $(shell id -u):$(shell id -g) debbuild/$@

.PHONY: debian-stretch
debian-stretch: ## build debian stretch deb packages
	docker build -t debbuild-$@/$(ARCH) -f $(CURDIR)/$@/Dockerfile.$(ARCH) .
	docker run --rm -i \
		-e DEB_VERSION=$(DEB_VERSION) \
		-e VERSION=$(VERSION) \
		-e DOCKER_GITCOMMIT=$(GITCOMMIT) \
		-v $(CURDIR)/debbuild/$@:/build \
		-v $(ENGINE_DIR):/engine \
		-v $(CLI_DIR):/cli \
		-v $(CURDIR)/systemd:/root/build-deb/systemd \
		debbuild-$@/$(ARCH)
	$(CHOWN) -R $(shell id -u):$(shell id -g) debbuild/$@

.PHONY: debian-wheezy
debian-wheezy: ## build debian wheezy deb packages
	docker build -t debbuild-$@/$(ARCH) -f $(CURDIR)/$@/Dockerfile.$(ARCH) .
	docker run --rm -i \
		-e DEB_VERSION=$(DEB_VERSION) \
		-e VERSION=$(VERSION) \
		-e DOCKER_GITCOMMIT=$(GITCOMMIT) \
		-v $(CURDIR)/debbuild/$@:/build \
		-v $(ENGINE_DIR):/engine \
		-v $(CLI_DIR):/cli \
		-v $(CURDIR)/systemd:/root/build-deb/systemd \
		debbuild-$@/$(ARCH)
	$(CHOWN) -R $(shell id -u):$(shell id -g) debbuild/$@

.PHONY: raspbian-jessie
raspbian-jessie: ## build raspbian jessie deb packages
	docker build -t debbuild-$@/$(ARCH) -f $(CURDIR)/$@/Dockerfile.$(ARCH) .
	docker run --rm -i \
		-e DEB_VERSION=$(DEB_VERSION) \
		-e VERSION=$(VERSION) \
		-e DOCKER_GITCOMMIT=$(GITCOMMIT) \
		-v $(CURDIR)/debbuild/$@:/build \
		-v $(ENGINE_DIR):/engine \
		-v $(CLI_DIR):/cli \
		-v $(CURDIR)/systemd:/root/build-deb/systemd \
		debbuild-$@/$(ARCH)
	$(CHOWN) -R $(shell id -u):$(shell id -g) debbuild/$@

.PHONY: raspbian-stretch
raspbian-stretch: ## build raspbian stretch deb packages
	docker build -t debbuild-$@/$(ARCH) -f $(CURDIR)/$@/Dockerfile.$(ARCH) .
	docker run --rm -i \
		-e DEB_VERSION=$(DEB_VERSION) \
		-e VERSION=$(VERSION) \
		-e DOCKER_GITCOMMIT=$(GITCOMMIT) \
		-v $(CURDIR)/debbuild/$@:/build \
		-v $(ENGINE_DIR):/engine \
		-v $(CLI_DIR):/cli \
		-v $(CURDIR)/systemd:/root/build-deb/systemd \
		debbuild-$@/$(ARCH)
	$(CHOWN) -R $(shell id -u):$(shell id -g) debbuild/$@
