include ../common.mk

PLUGINS_DIR=$(realpath $(CURDIR)/../plugins)
GO_BASE_IMAGE=golang
GO_IMAGE?=$(GO_BASE_IMAGE):$(GO_VERSION)-buster
GEN_RPM_VER=$(shell ./gen-rpm-ver $(realpath $(CURDIR)/../src/github.com/docker/cli) "$(VERSION)")
CLI_GITCOMMIT?=$(shell cd $(realpath $(CURDIR)/../src/github.com/docker/cli) && git rev-parse --short HEAD)
ENGINE_GITCOMMIT?=$(shell cd $(realpath $(CURDIR)/../src/github.com/docker/docker) && git rev-parse --short HEAD)

ifdef BUILD_IMAGE
	BUILD_IMAGE_FLAG=--build-arg $(BUILD_IMAGE)
endif
BUILD?=DOCKER_BUILDKIT=1 \
	docker build \
	$(BUILD_IMAGE_FLAG) \
	--build-arg GO_IMAGE=$(GO_IMAGE) \
	-t rpmbuild-$@/$(ARCH) \
	-f $@/Dockerfile \
	.

SPEC_FILES?=docker-ce.spec docker-ce-cli.spec docker-ce-rootless-extras.spec
SPECS?=$(addprefix SPECS/, $(SPEC_FILES))
RPMBUILD_FLAGS?=-ba\
	--define '_gitcommit_cli $(CLI_GITCOMMIT)' \
	--define '_gitcommit_engine $(ENGINE_GITCOMMIT)' \
	--define '_release $(word 2,$(GEN_RPM_VER))' \
	--define '_version $(word 1,$(GEN_RPM_VER))' \
	--define '_origversion $(word 4, $(GEN_RPM_VER))' \
	$(RPMBUILD_EXTRA_FLAGS) \
	$(SPECS)

# Additional flags may be necessary at some point
RUN_FLAGS=
RUN?=docker run --rm \
	-e PLATFORM \
	-v $(CURDIR)/rpmbuild/SOURCES:/root/rpmbuild/SOURCES \
	-v $(CURDIR)/rpmbuild/RPMS:/root/rpmbuild/RPMS \
	-v $(CURDIR)/rpmbuild/SRPMS:/root/rpmbuild/SRPMS \
	$(RUN_FLAGS) \
	rpmbuild-$@/$(ARCH) $(RPMBUILD_FLAGS)

FEDORA_RELEASES := fedora-32 fedora-31
CENTOS_RELEASES := centos-7 centos-8
RHEL_RELEASES := rhel-7
DISTROS := $(FEDORA_RELEASES) $(CENTOS_RELEASES) $(RHEL_RELEASES)

.PHONY: help
help: ## show make targets
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {sub("\\\\n",sprintf("\n%22c"," "), $$2);printf " \033[36m%-20s\033[0m  %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.PHONY: clean
clean: ## remove build artifacts
	[ ! -d rpmbuild ] || $(CHOWN) -R $(shell id -u):$(shell id -g) rpmbuild
	$(RM) -r rpmbuild/

.PHONY: rpm
rpm: fedora centos ## build all rpm packages

.PHONY: fedora
fedora: $(FEDORA_RELEASES) ## build all fedora rpm packages

.PHONY: centos-8
centos-8: RPMBUILD_EXTRA_FLAGS=--define '_without_btrfs 1'

.PHONY: centos
centos: $(CENTOS_RELEASES) ## build all centos rpm packages

.PHONY: rhel
rhel: $(RHEL_RELEASES) ## build all rhel rpm packages

.PHONY: $(DISTROS)
$(DISTROS): rpmbuild/SOURCES/engine.tgz rpmbuild/SOURCES/cli.tgz rpmbuild/SOURCES/docker.service rpmbuild/SOURCES/docker.socket rpmbuild/SOURCES/plugin-installers.tgz
	@echo "== Building packages for $@ =="
	$(CHOWN) -R root:root rpmbuild
	$(BUILD)
	$(RUN)
	$(CHOWN) -R $(shell id -u):$(shell id -g) rpmbuild

rpmbuild/SOURCES/engine.tgz:
	mkdir -p $(@D)
	docker run --rm -w /v \
		-v $(realpath $(CURDIR)/../src/github.com/docker/docker):/engine \
		-v $(CURDIR)/$(@D):/v \
		alpine \
		tar -C / -c -z -f /v/engine.tgz --exclude .git engine

rpmbuild/SOURCES/cli.tgz:
	mkdir -p $(@D)
	docker run --rm -w /v \
		-v $(realpath $(CURDIR)/../src/github.com/docker/cli):/cli \
		-v $(CURDIR)/$(@D):/v \
		alpine \
		tar -C / -c -z -f /v/cli.tgz --exclude .git cli

rpmbuild/SOURCES/docker.service: ../systemd/docker.service
	mkdir -p $(@D)
	cp $< $@

rpmbuild/SOURCES/docker.socket: ../systemd/docker.socket
	mkdir -p $(@D)
	cp $< $@

rpmbuild/SOURCES/plugin-installers.tgz: $(wildcard ../plugins/*)
	docker run --rm -w /v \
		-v $(PLUGINS_DIR):/plugins \
		-v $(CURDIR)/$(@D):/v \
		alpine \
		tar -C / -c -z -f /v/plugin-installers.tgz --exclude .git plugins
