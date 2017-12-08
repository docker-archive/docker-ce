SHELL:=/bin/bash
ARCH:=$(shell uname -m)
ENGINE_DIR:=$(CURDIR)/../../engine
CLI_DIR:=$(CURDIR)/../../cli
GITCOMMIT?=$(shell cd $(ENGINE_DIR) && git rev-parse --short HEAD)
VERSION?=0.0.0-dev
DOCKER_EXPERIMENTAL:=0
CHOWN:=docker run --rm -v $(CURDIR):/v -w /v alpine chown

.PHONY: help clean deb ubuntu debian ubuntu-xenial ubuntu-trusty ubuntu-zesty debian-buster debian-jessie debian-stretch debian-wheezy raspbian-jessie raspbian-stretch

help: ## show make targets
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {sub("\\\\n",sprintf("\n%22c"," "), $$2);printf " \033[36m%-20s\033[0m  %s\n", $$1, $$2}' $(MAKEFILE_LIST)

clean: ## remove build artifacts
	[ ! -d debbuild ] || $(CHOWN) -R $(shell id -u):$(shell id -g) debbuild
	$(RM) -r debbuild

deb: ubuntu debian raspbian ## build all deb packages

ubuntu: ubuntu-zesty ubuntu-xenial ubuntu-trusty ## build all ubuntu deb packages

debian: debian-stretch debian-wheezy debian-jessie ## build all debian deb packages

raspbian: raspbian-stretch debian-jessie ## build all raspbian deb packages

ubuntu-xenial: ## build ubuntu xenial deb packages
	docker build -t debbuild-$@/$(ARCH) -f $(CURDIR)/$@/Dockerfile.$(ARCH) .
	docker run --rm -i \
		-e VERSION=$(VERSION) \
		-e DOCKER_GITCOMMIT=$(GITCOMMIT) \
		-v $(CURDIR)/debbuild/$@:/build \
		-v $(ENGINE_DIR):/engine \
		-v $(CLI_DIR):/cli \
		-v $(CURDIR)/systemd:/root/build-deb/systemd \
		debbuild-$@/$(ARCH)
	$(CHOWN) -R $(shell id -u):$(shell id -g) debbuild/$@

ubuntu-trusty: ## build ubuntu trusty deb packages
	docker build -t debbuild-$@/$(ARCH) -f $(CURDIR)/$@/Dockerfile.$(ARCH) .
	docker run --rm -i \
		-e VERSION=$(VERSION) \
		-e DOCKER_GITCOMMIT=$(GITCOMMIT) \
		-v $(CURDIR)/debbuild/$@:/build \
		-v $(ENGINE_DIR):/engine \
		-v $(CLI_DIR):/cli \
		-v $(CURDIR)/systemd:/root/build-deb/systemd \
		debbuild-$@/$(ARCH)
	$(CHOWN) -R $(shell id -u):$(shell id -g) debbuild/$@

ubuntu-zesty: ## build ubuntu zesty deb packages
	docker build -t debbuild-$@/$(ARCH) -f $(CURDIR)/$@/Dockerfile.$(ARCH) .
	docker run --rm -i \
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
		-e VERSION=$(VERSION) \
		-e DOCKER_GITCOMMIT=$(GITCOMMIT) \
		-v $(CURDIR)/debbuild/$@:/build \
		-v $(ENGINE_DIR):/engine \
		-v $(CLI_DIR):/cli \
		-v $(CURDIR)/systemd:/root/build-deb/systemd \
		debbuild-$@/$(ARCH)
	$(CHOWN) -R $(shell id -u):$(shell id -g) debbuild/$@

debian-buster: ## build debian buster deb packages
	docker build -t debbuild-$@/$(ARCH) -f $(CURDIR)/$@/Dockerfile.$(ARCH) .
	docker run --rm -i \
		-e VERSION=$(VERSION) \
		-e DOCKER_GITCOMMIT=$(GITCOMMIT) \
		-v $(CURDIR)/debbuild/$@:/build \
		-v $(ENGINE_DIR):/engine \
		-v $(CLI_DIR):/cli \
		-v $(CURDIR)/systemd:/root/build-deb/systemd \
		debbuild-$@/$(ARCH)
	$(CHOWN) -R $(shell id -u):$(shell id -g) debbuild/$@

debian-jessie: ## build debian jessie deb packages
	docker build -t debbuild-$@/$(ARCH) -f $(CURDIR)/$@/Dockerfile.$(ARCH) .
	docker run --rm -i \
		-e VERSION=$(VERSION) \
		-e DOCKER_GITCOMMIT=$(GITCOMMIT) \
		-v $(CURDIR)/debbuild/$@:/build \
		-v $(ENGINE_DIR):/engine \
		-v $(CLI_DIR):/cli \
		-v $(CURDIR)/systemd:/root/build-deb/systemd \
		debbuild-$@/$(ARCH)
	$(CHOWN) -R $(shell id -u):$(shell id -g) debbuild/$@

debian-stretch: ## build debian stretch deb packages
	docker build -t debbuild-$@/$(ARCH) -f $(CURDIR)/$@/Dockerfile.$(ARCH) .
	docker run --rm -i \
		-e VERSION=$(VERSION) \
		-e DOCKER_GITCOMMIT=$(GITCOMMIT) \
		-v $(CURDIR)/debbuild/$@:/build \
		-v $(ENGINE_DIR):/engine \
		-v $(CLI_DIR):/cli \
		-v $(CURDIR)/systemd:/root/build-deb/systemd \
		debbuild-$@/$(ARCH)
	$(CHOWN) -R $(shell id -u):$(shell id -g) debbuild/$@

debian-wheezy: ## build debian wheezy deb packages
	docker build -t debbuild-$@/$(ARCH) -f $(CURDIR)/$@/Dockerfile.$(ARCH) .
	docker run --rm -i \
		-e VERSION=$(VERSION) \
		-e DOCKER_GITCOMMIT=$(GITCOMMIT) \
		-v $(CURDIR)/debbuild/$@:/build \
		-v $(ENGINE_DIR):/engine \
		-v $(CLI_DIR):/cli \
		-v $(CURDIR)/systemd:/root/build-deb/systemd \
		debbuild-$@/$(ARCH)
	$(CHOWN) -R $(shell id -u):$(shell id -g) debbuild/$@

raspbian-jessie: ## build raspbian jessie deb packages
	docker build -t debbuild-$@/$(ARCH) -f $(CURDIR)/$@/Dockerfile.$(ARCH) .
	docker run --rm -i \
		-e VERSION=$(VERSION) \
		-e DOCKER_GITCOMMIT=$(GITCOMMIT) \
		-v $(CURDIR)/debbuild/$@:/build \
		-v $(ENGINE_DIR):/engine \
		-v $(CLI_DIR):/cli \
		-v $(CURDIR)/systemd:/root/build-deb/systemd \
		debbuild-$@/$(ARCH)
	$(CHOWN) -R $(shell id -u):$(shell id -g) debbuild/$@

raspbian-stretch: ## build raspbian stretch deb packages
	docker build -t debbuild-$@/$(ARCH) -f $(CURDIR)/$@/Dockerfile.$(ARCH) .
	docker run --rm -i \
		-e VERSION=$(VERSION) \
		-e DOCKER_GITCOMMIT=$(GITCOMMIT) \
		-v $(CURDIR)/debbuild/$@:/build \
		-v $(ENGINE_DIR):/engine \
		-v $(CLI_DIR):/cli \
		-v $(CURDIR)/systemd:/root/build-deb/systemd \
		debbuild-$@/$(ARCH)
	$(CHOWN) -R $(shell id -u):$(shell id -g) debbuild/$@
