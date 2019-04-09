include ../containerd.mk

SHELL:=/bin/bash
ARCH:=$(shell uname -m)
CLI_DIR:=$(CURDIR)/../../cli
ENGINE_DIR:=$(CURDIR)/../../engine
GITCOMMIT?=$(shell cd $(CLI_DIR) && git rev-parse --short HEAD)
VERSION?=0.0.0-dev
STATIC_VERSION=$(shell ../static/gen-static-ver $(ENGINE_DIR) $(VERSION))
GO_BASE_IMAGE=golang
GO_VERSION:=1.11.5
GO_IMAGE=$(GO_BASE_IMAGE):$(GO_VERSION)
DEB_VERSION=$(shell ./gen-deb-ver $(CLI_DIR) "$(VERSION)")
CHOWN:=docker run --rm -v $(CURDIR):/v -w /v alpine chown
EPOCH?=5

ifdef BUILD_IMAGE
	BUILD_IMAGE_FLAG=--build-arg $(BUILD_IMAGE)
endif

COMMON_FILES=common
BUILD?=docker build \
	$(BUILD_IMAGE_FLAG) \
	--build-arg GO_IMAGE=$(GO_IMAGE) \
	--build-arg COMMON_FILES=$(COMMON_FILES) \
	-t debbuild-$@/$(ARCH) \
	-f $(CURDIR)/$@/Dockerfile .
# Additional flags may be necessary at some point
RUN_FLAGS=
RUN=docker run --rm -i \
	-e EPOCH='$(EPOCH)' \
	-e DEB_VERSION=$(word 1, $(DEB_VERSION)) \
	-e VERSION=$(word 2, $(DEB_VERSION)) \
	-e DOCKER_GITCOMMIT=$(GITCOMMIT) \
	-e PLATFORM \
	-v $(CURDIR)/debbuild/$@:/build \
	$(RUN_FLAGS) \
	debbuild-$@/$(ARCH)

SOURCE_FILES=engine-image cli.tgz engine.tgz docker.service docker.socket distribution_based_engine.json plugin-installers.tgz
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
	$(RM) engine-image

engine-$(ARCH).tar:
	$(MAKE) -C ../image image-linux
	docker save -o $@ $$(cat ../image/image-linux)

.PHONY: deb
deb: ubuntu debian raspbian ## build all deb packages

.PHONY: ubuntu
ubuntu: ubuntu-bionic ubuntu-xenial ## build all ubuntu deb packages

.PHONY: debian
debian: debian-stretch ## build all debian deb packages

.PHONY: raspbian
raspbian: raspbian-stretch ## build all raspbian deb packages

.PHONY: ubuntu-%
ubuntu-%: ## build ubuntu deb packages
ubuntu-%: $(SOURCES)
	$(BUILD)
	$(RUN)
	$(CHOWN) -R $(shell id -u):$(shell id -g) debbuild/$@

.PHONY: debian-buster
debian-buster: ## build debian buster deb packages
debian-buster: $(SOURCES)
	$(BUILD)
	$(RUN)
	$(CHOWN) -R $(shell id -u):$(shell id -g) debbuild/$@

.PHONY: debian-stretch
debian-stretch: ## build debian stretch deb packages
debian-stretch: $(SOURCES)
	$(BUILD)
	$(RUN)
	$(CHOWN) -R $(shell id -u):$(shell id -g) debbuild/$@

.PHONY: raspbian-stretch
raspbian-stretch: ## build raspbian stretch deb packages
raspbian-stretch: $(SOURCES)
	$(BUILD)
	$(RUN)
	$(CHOWN) -R $(shell id -u):$(shell id -g) debbuild/$@

sources/engine.tgz:
	mkdir -p $(@D)
	docker run --rm -i -w /v \
		-v $(ENGINE_DIR):/engine \
		-v $(CURDIR)/$(@D):/v \
		alpine \
		tar -C / -c -z -f /v/engine.tgz --exclude .git engine

sources/cli.tgz:
	mkdir -p $(@D)
	docker run --rm -i -w /v \
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

# TODO: Figure out how to decouple this
# TODO: These might just end up being static files that are hardcoded
# TODO: FROM HERE <=====================
sources/distribution_based_engine.json: sources/engine-image
	mkdir -p $(@D)
	echo '{"platform":"Docker Engine - Community","engine_image":"engine-community","containerd_min_version":"1.2.0-beta.1","runtime":"host_install"}' > $@

sources/engine-image:
	mkdir -p $(@D)
	echo "docker.io/dockereng/engine-community:$(STATIC_VERSION)" > $@
# TODO: TO HERE <=====================

sources/plugin-installers.tgz: $(wildcard ../plugins/*)
	docker run --rm -i -w /v \
		-v $(shell readlink -e ../plugins):/plugins \
		-v $(CURDIR)/sources:/v \
		alpine \
		tar -C / -c -z -f /v/plugin-installers.tgz --exclude .git plugins
