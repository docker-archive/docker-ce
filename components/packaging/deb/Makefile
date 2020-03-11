include ../common.mk

CLI_DIR=$(realpath $(CURDIR)/../../cli)
ENGINE_DIR=$(realpath $(CURDIR)/../../engine)
PLUGINS_DIR=$(realpath $(CURDIR)/../plugins)
GITCOMMIT?=$(shell cd $(CLI_DIR) && git rev-parse --short HEAD)
GO_BASE_IMAGE=golang
GO_IMAGE?=$(GO_BASE_IMAGE):$(GO_VERSION)-buster
CHOWN:=docker run --rm -v $(CURDIR):/v -w /v alpine chown
DEB_VERSION=$(shell ./gen-deb-ver $(CLI_DIR) "$(VERSION)")
EPOCH?=5

ifdef BUILD_IMAGE
	BUILD_IMAGE_FLAG=--build-arg $(BUILD_IMAGE)
endif

COMMON_FILES=common
BUILD?=DOCKER_BUILDKIT=1 \
	docker build \
	$(BUILD_IMAGE_FLAG) \
	--build-arg GO_IMAGE=$(GO_IMAGE) \
	--build-arg COMMON_FILES=$(COMMON_FILES) \
	-t debbuild-$@/$(ARCH) \
	-f $@/Dockerfile \
	.

# Additional flags may be necessary at some point
RUN_FLAGS=
RUN?=docker run --rm \
	-e PLATFORM \
	-e EPOCH='$(EPOCH)' \
	-e DEB_VERSION=$(word 1, $(DEB_VERSION)) \
	-e VERSION=$(word 2, $(DEB_VERSION)) \
	-e DOCKER_GITCOMMIT=$(GITCOMMIT) \
	-v $(CURDIR)/debbuild/$@:/build \
	$(RUN_FLAGS) \
	debbuild-$@/$(ARCH)

DEBIAN_VERSIONS := debian-stretch debian-buster
UBUNTU_VERSIONS := ubuntu-xenial ubuntu-bionic ubuntu-cosmic ubuntu-disco ubuntu-eoan
RASPBIAN_VERSIONS := raspbian-stretch raspbian-buster
DISTROS := $(DEBIAN_VERSIONS) $(UBUNTU_VERSIONS) $(RASPBIAN_VERSIONS)

.PHONY: help
help: ## show make targets
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {sub("\\\\n",sprintf("\n%22c"," "), $$2);printf " \033[36m%-20s\033[0m  %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.PHONY: clean
clean: ## remove build artifacts
	[ ! -d debbuild ] || $(CHOWN) -R $(shell id -u):$(shell id -g) debbuild
	$(RM) -r debbuild
	[ ! -d sources ] || $(CHOWN) -R $(shell id -u):$(shell id -g) sources
	$(RM) -r sources

.PHONY: deb
deb: ubuntu debian ## build all deb packages except for raspbian

.PHONY: ubuntu
ubuntu: $(UBUNTU_VERSIONS) ## build all ubuntu deb packages

.PHONY: debian
debian: $(DEBIAN_VERSIONS) ## build all debian deb packages

.PHONY: raspbian
raspbian: $(RASPBIAN_VERSIONS) ## build all raspbian deb packages

.PHONY: $(DISTROS)
$(DISTROS): sources/cli.tgz sources/engine.tgz sources/docker.service sources/docker.socket sources/plugin-installers.tgz
	@echo "== Building packages for $@ =="
	$(BUILD)
	$(RUN)
	$(CHOWN) -R $(shell id -u):$(shell id -g) debbuild/$@

sources/engine.tgz:
	mkdir -p $(@D)
	docker run --rm -w /v \
		-v $(ENGINE_DIR):/engine \
		-v $(CURDIR)/$(@D):/v \
		alpine \
		tar -C / -c -z -f /v/engine.tgz --exclude .git engine

sources/cli.tgz:
	mkdir -p $(@D)
	docker run --rm -w /v \
		-v $(CLI_DIR):/cli \
		-v $(CURDIR)/$(@D):/v \
		alpine \
		tar -C / -c -z -f /v/cli.tgz --exclude .git cli

sources/docker.service: ../systemd/docker.service
	mkdir -p $(@D)
	cp $< $@

sources/docker.socket: ../systemd/docker.socket
	mkdir -p $(@D)
	cp $< $@

sources/plugin-installers.tgz: $(wildcard ../plugins/*)
	docker run --rm -w /v \
		-v $(PLUGINS_DIR):/plugins \
		-v $(CURDIR)/$(@D):/v \
		alpine \
		tar -C / -c -z -f /v/plugin-installers.tgz --exclude .git plugins
