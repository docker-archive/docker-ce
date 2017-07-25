SHELL:=/bin/bash
ALPINE_IMG:=$(shell $(CURDIR)/../detect_alpine_image)
ARCH:=$(shell uname -m)
KERNEL:=$(shell docker version --format '{{.Server.Os}}')
ENGINE_DIR:=$(CURDIR)/../../engine
CLI_DIR:=$(CURDIR)/../../cli
COMMON_DIR:=$(CURDIR)/../common
PLUGIN_DIR:=$(CURDIR)/plugin
GITCOMMIT?=$(shell cd $(ENGINE_DIR) && git rev-parse --short HEAD)
VERSION?=$(shell cat $(ENGINE_DIR)/VERSION)
PLUGIN_VERSION?=1.0.0.$(KERNEL)-$(ARCH)-test
DOCKER_EXPERIMENTAL:=0
CHOWN:=docker run --rm -v $(CURDIR):/v -w /v $(ALPINE_IMG) chown

.PHONY: help clean deb ubuntu debian ubuntu-xenial ubuntu-trusty ubuntu-zesty debian-jessie debian-stretch debian-wheezy raspbian-jessie raspbian-stretch

help: ## show make targets
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {sub("\\\\n",sprintf("\n%22c"," "), $$2);printf " \033[36m%-20s\033[0m  %s\n", $$1, $$2}' $(MAKEFILE_LIST)

clean: ## remove build artifacts
	[ ! -d debbuild ] || $(CHOWN) -R $(shell id -u):$(shell id -g) debbuild
	$(RM) -r debbuild

deb: ubuntu debian raspbian ## build all deb packages

ubuntu: ubuntu-zesty ubuntu-xenial ubuntu-trusty ## build all ubuntu deb packages

debian: debian-stretch debian-wheezy debian-jessie ## build all debian deb packages

raspbian: raspbian-stretch debian-jessie ## build all raspbian deb packages

metrics_plugin: ## pull telemetry plugin tarball
	mkdir -p $(PLUGIN_DIR)
	docker create --name tmp docker/telemetry-tgz:$(PLUGIN_VERSION)
	docker export tmp | tar -x -C $(PLUGIN_DIR)
	docker rm -f tmp

ubuntu-xenial: ## build ubuntu xenial deb packages
	docker build -t debbuild-$@/$(ARCH) -f $(CURDIR)/$@/Dockerfile.$(ARCH) .
	docker run --rm -i \
		-e VERSION=$(VERSION) \
		-e PLUGIN_VERSION=$(PLUGIN_VERSION) \
		-e DOCKER_GITCOMMIT=$(GITCOMMIT) \
		-v $(CURDIR)/debbuild/$@:/build \
		-v $(ENGINE_DIR):/engine \
		-v $(PLUGIN_DIR):/plugin \
		-v $(CLI_DIR):/cli \
		-v $(CURDIR)/systemd:/root/build-deb/systemd \
		-v $(COMMON_DIR):/root/build-deb/common \
		debbuild-$@/$(ARCH)
	$(CHOWN) -R $(shell id -u):$(shell id -g) debbuild/$@

ubuntu-trusty: ## build ubuntu trusty deb packages
	docker build -t debbuild-$@/$(ARCH) -f $(CURDIR)/$@/Dockerfile.$(ARCH) .
	docker run --rm -i \
		-e VERSION=$(VERSION) \
		-e PLUGIN_VERSION=$(PLUGIN_VERSION) \
		-e DOCKER_GITCOMMIT=$(GITCOMMIT) \
		-v $(CURDIR)/debbuild/$@:/build \
		-v $(ENGINE_DIR):/engine \
		-v $(PLUGIN_DIR):/plugin \
		-v $(CLI_DIR):/cli \
		-v $(CURDIR)/systemd:/root/build-deb/systemd \
		-v $(COMMON_DIR):/root/build-deb/common \
		debbuild-$@/$(ARCH)
	$(CHOWN) -R $(shell id -u):$(shell id -g) debbuild/$@

ubuntu-zesty: ## build ubuntu zesty deb packages
	docker build -t debbuild-$@/$(ARCH) -f $(CURDIR)/$@/Dockerfile.$(ARCH) .
	docker run --rm -i \
		-e VERSION=$(VERSION) \
		-e DOCKER_GITCOMMIT=$(GITCOMMIT) \
		-e PLUGIN_VERSION=$(PLUGIN_VERSION) \
		-v $(CURDIR)/debbuild/$@:/build \
		-v $(ENGINE_DIR):/engine \
		-v $(PLUGIN_DIR):/plugin \
		-v $(CLI_DIR):/cli \
		-v $(CURDIR)/systemd:/root/build-deb/systemd \
		-v $(COMMON_DIR):/root/build-deb/common \
		debbuild-$@/$(ARCH)
	$(CHOWN) -R $(shell id -u):$(shell id -g) debbuild/$@

debian-jessie: ## build debian jessie deb packages
	docker build -t debbuild-$@/$(ARCH) -f $(CURDIR)/$@/Dockerfile.$(ARCH) .
	docker run --rm -i \
		-e VERSION=$(VERSION) \
		-e DOCKER_GITCOMMIT=$(GITCOMMIT) \
		-e PLUGIN_VERSION=$(PLUGIN_VERSION) \
		-v $(CURDIR)/debbuild/$@:/build \
		-v $(ENGINE_DIR):/engine \
		-v $(PLUGIN_DIR):/plugin \
		-v $(CLI_DIR):/cli \
		-v $(CURDIR)/systemd:/root/build-deb/systemd \
		-v $(COMMON_DIR):/root/build-deb/common \
		debbuild-$@/$(ARCH)
	$(CHOWN) -R $(shell id -u):$(shell id -g) debbuild/$@

debian-stretch: ## build debian stretch deb packages
	docker build -t debbuild-$@/$(ARCH) -f $(CURDIR)/$@/Dockerfile.$(ARCH) .
	docker run --rm -i \
		-e VERSION=$(VERSION) \
		-e DOCKER_GITCOMMIT=$(GITCOMMIT) \
		-e PLUGIN_VERSION=$(PLUGIN_VERSION) \
		-v $(CURDIR)/debbuild/$@:/build \
		-v $(ENGINE_DIR):/engine \
		-v $(PLUGIN_DIR):/plugin \
		-v $(CLI_DIR):/cli \
		-v $(CURDIR)/systemd:/root/build-deb/systemd \
		-v $(COMMON_DIR):/root/build-deb/common \
		debbuild-$@/$(ARCH)
	$(CHOWN) -R $(shell id -u):$(shell id -g) debbuild/$@

debian-wheezy: ## build debian wheezy deb packages
	docker build -t debbuild-$@/$(ARCH) -f $(CURDIR)/$@/Dockerfile.$(ARCH) .
	docker run --rm -i \
		-e VERSION=$(VERSION) \
		-e DOCKER_GITCOMMIT=$(GITCOMMIT) \
		-e PLUGIN_VERSION=$(PLUGIN_VERSION) \
		-v $(CURDIR)/debbuild/$@:/build \
		-v $(ENGINE_DIR):/engine \
		-v $(PLUGIN_DIR):/plugin \
		-v $(CLI_DIR):/cli \
		-v $(CURDIR)/systemd:/root/build-deb/systemd \
		-v $(COMMON_DIR):/root/build-deb/common \
		debbuild-$@/$(ARCH)
	$(CHOWN) -R $(shell id -u):$(shell id -g) debbuild/$@

raspbian-jessie: ## build raspbian jessie deb packages
	docker build -t debbuild-$@/$(ARCH) -f $(CURDIR)/$@/Dockerfile.$(ARCH) .
	docker run --rm -i \
		-e VERSION=$(VERSION) \
		-e DOCKER_GITCOMMIT=$(GITCOMMIT) \
		-e PLUGIN_VERSION=$(PLUGIN_VERSION) \
		-v $(CURDIR)/debbuild/$@:/build \
		-v $(ENGINE_DIR):/engine \
		-v $(PLUGIN_DIR):/plugin \
		-v $(CLI_DIR):/cli \
		-v $(CURDIR)/systemd:/root/build-deb/systemd \
		-v $(COMMON_DIR):/root/build-deb/common \
		debbuild-$@/$(ARCH)
	$(CHOWN) -R $(shell id -u):$(shell id -g) debbuild/$@

raspbian-stretch: ## build raspbian stretch deb packages
	docker build -t debbuild-$@/$(ARCH) -f $(CURDIR)/$@/Dockerfile.$(ARCH) .
	docker run --rm -i \
		-e VERSION=$(VERSION) \
		-e DOCKER_GITCOMMIT=$(GITCOMMIT) \
		-e PLUGIN_VERSION=$(PLUGIN_VERSION) \
		-v $(CURDIR)/debbuild/$@:/build \
		-v $(ENGINE_DIR):/engine \
		-v $(PLUGIN_DIR):/plugin \
		-v $(CLI_DIR):/cli \
		-v $(CURDIR)/systemd:/root/build-deb/systemd \
		-v $(COMMON_DIR):/root/build-deb/common \
		debbuild-$@/$(ARCH)
	$(CHOWN) -R $(shell id -u):$(shell id -g) debbuild/$@
